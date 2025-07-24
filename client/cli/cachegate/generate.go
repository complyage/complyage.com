// ||------------------------------------------------------------------------------------------------||
// || Import
// ||------------------------------------------------------------------------------------------------||
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ||------------------------------------------------------------------------------------------------||
// || Types
// ||------------------------------------------------------------------------------------------------||
type Zone struct {
	ID        int
	State     string
	Country   string
	Effective string
}

// Translations maps language code to key->value map
type Translations map[string]map[string]string

// ||------------------------------------------------------------------------------------------------||
// || Config Paths
// ||------------------------------------------------------------------------------------------------||
const (
	tplPath   = "assets/age-gate.html"
	i18nDir   = "i18n"
	outputDir = "cache/age-gate"
)

// ||------------------------------------------------------------------------------------------------||
// || replaceMarkers
// ||------------------------------------------------------------------------------------------------||
// replaces [%%KEY%%] in tpl with values from vars; missing keys stay as [!!KEY!!]
func replaceMarkers(tpl string, vars map[string]string) string {
	re := regexp.MustCompile(`\[\%\%([A-Z0-9_]+)\%\%\]`)
	return re.ReplaceAllStringFunc(tpl, func(marker string) string {
		key := re.FindStringSubmatch(marker)[1]
		if v, ok := vars[key]; ok {
			return v
		}
		return "[!!" + key + "!!]"
	})
}

// ||------------------------------------------------------------------------------------------------||
// || main
// ||------------------------------------------------------------------------------------------------||
func main() {
	// load .env if present
	_ = godotenv.Load()

	// read DB credentials from env
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	if host == "" || port == "" || user == "" || pass == "" || name == "" {
		fmt.Fprintln(os.Stderr, "missing one of DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, or DB_NAME")
		os.Exit(1)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	// load template
	tplBytes, err := ioutil.ReadFile(tplPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read tpl error: %v\n", err)
		os.Exit(1)
	}
	tpl := string(tplBytes)

	// load translations from i18n directory
	langs := Translations{}
	files, err := ioutil.ReadDir(i18nDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read i18n dir error: %v\n", err)
		os.Exit(1)
	}
	for _, fi := range files {
		if fi.IsDir() || !strings.HasSuffix(fi.Name(), ".json") {
			continue
		}
		code := strings.TrimSuffix(fi.Name(), ".json")
		path := filepath.Join(i18nDir, fi.Name())
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read lang file %s error: %v\n", path, err)
			os.Exit(1)
		}
		var dict map[string]string
		if err := json.Unmarshal(data, &dict); err != nil {
			fmt.Fprintf(os.Stderr, "unmarshal lang file %s error: %v\n", path, err)
			os.Exit(1)
		}
		langs[code] = dict
	}

	// connect DB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "db open error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// query zones
	rows, err := db.Query(`
        SELECT id_zone, COALESCE(zone_state,'') AS state, COALESCE(zone_country,'') AS country,
               DATE_FORMAT(zone_effective, '%M %d, %Y') AS effective
        FROM zones
    `)
	if err != nil {
		fmt.Fprintf(os.Stderr, "query error: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var zones []Zone
	for rows.Next() {
		var z Zone
		if err := rows.Scan(&z.ID, &z.State, &z.Country, &z.Effective); err != nil {
			fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
			os.Exit(1)
		}
		zones = append(zones, z)
	}

	// constants
	const (
		signupURL = "https://example.com/signup"
		oauthURL  = "https://example.com/oauth"
		exitURL   = "https://example.com/exit"
	)

	// ensure output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir cache error: %v\n", err)
		os.Exit(1)
	}

	// render per language and zone
	for code, dict := range langs {
		for _, z := range zones {
			vars := make(map[string]string, len(dict)+5)
			for k, v := range dict {
				vars[k] = v
			}
			vars["SIGNUPURL"] = signupURL
			vars["OAUTHURL"] = oauthURL
			vars["EXITURL"] = exitURL
			parts := []string{}
			if z.State != "" {
				parts = append(parts, z.State)
			}
			if z.Country != "" {
				parts = append(parts, z.Country)
			}
			vars["LOCATION"] = strings.Join(parts, ", ")
			vars["EFFECTIVE"] = z.Effective

			out := replaceMarkers(tpl, vars)
			fname := filepath.Join(outputDir, fmt.Sprintf("%s.%d.html", code, z.ID))
			if err := ioutil.WriteFile(fname, []byte(out), 0644); err != nil {
				fmt.Fprintf(os.Stderr, "write file error: %v\n", err)
				continue
			}
			fmt.Printf("Wrote %s\n", fname)
		}
	}
}

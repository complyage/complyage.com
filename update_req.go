package main

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import (
	"bufio"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// ||------------------------------------------------------------------------------------------------||
// || .env Loader (minimal)
// ||------------------------------------------------------------------------------------------------||
func loadDotEnv(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// support comments after value using # (keeps your #|| style safe)
		if i := strings.Index(line, "#"); i >= 0 {
			line = strings.TrimSpace(line[:i])
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		if _, exists := os.LookupEnv(k); !exists {
			_ = os.Setenv(k, v)
		}
	}
	return s.Err()
}

// ||------------------------------------------------------------------------------------------------||
// || Map: ZoneRequirement -> VerifyType code
// ||------------------------------------------------------------------------------------------------||
var zoneToVerifyCode = map[string]string{
	"ID_UPLOAD":     "IDEN",
	"GOV_ID":        "IDEN",
	"DIGITAL_ID":    "IDEN",
	"BIOMETRIC":     "FACE",
	"FACIAL_EST":    "FACE",
	"HAND_ANALYSIS": "FACE",
	"CREDIT_CARD":   "CRCD",
	"TXN_DATA":      "CRCD",
	"OPEN_BANKING":  "CRCD",
}

// ||------------------------------------------------------------------------------------------------||
// || Normalize & Map helpers
// ||------------------------------------------------------------------------------------------------||
func normalizeKey(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}

// mapOne returns either the mapped VerifyType code (when we know it), or the
// ORIGINAL token (preserving the original casing/spelling) when unknown.
// didMap == true only when we converted to one of our codes.
func mapOne(original string) (mapped string, didMap bool) {
	key := normalizeKey(original)
	if key == "" {
		return "", false
	}
	if code, ok := zoneToVerifyCode[key]; ok {
		return code, true
	}
	// leave as-is (same name) when it's not one of ours
	return strings.TrimSpace(original), false
}

// mapValue maps a raw value that may contain a single token or a comma list.
// - Known requirements → our codes (MAIL/PHNE/ADDR/CRCD/IDEN/FACE/USER etc.)
// - Unknown requirements → preserved EXACTLY as provided (same name / casing)
// - De-dupes while preserving first occurrence order
// Returns: joined string, changed flag, and any warnings
func mapValue(raw string) (string, bool, []string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", false, nil
	}

	// split on commas (support single token or list)
	var parts []string
	if strings.Contains(raw, ",") {
		for _, p := range strings.Split(raw, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				parts = append(parts, p)
			}
		}
	} else {
		parts = []string{raw}
	}

	out := make([]string, 0, len(parts))
	seen := make(map[string]struct{}, len(parts))
	changed := false
	var warns []string
	duplicatesRemoved := false

	for _, p := range parts {
		mapped, did := mapOne(p)
		// Warn only if not mapped AND not already a known VerifyType code
		if !did {
			switch strings.ToUpper(strings.TrimSpace(mapped)) {
			case "MAIL", "PHNE", "ADDR", "CRCD", "IDEN", "FACE", "USER":
				// already a code; fine
			default:
				// preserved as provided
				warns = append(warns, fmt.Sprintf("unmapped requirement '%s' (left unchanged)", p))
			}
		}

		// de-dup preserving first occurrence
		if _, ok := seen[mapped]; ok {
			duplicatesRemoved = true
			continue
		}
		seen[mapped] = struct{}{}
		out = append(out, mapped)

		if did { // only count as changed when we actually remapped to one of ours
			changed = true
		}
	}

	// consider de-dup effect as a change too
	if duplicatesRemoved {
		changed = true
	}

	// Join with commas (no extra spaces) to keep column tidy
	return strings.Join(out, ","), changed, warns
}

// ||------------------------------------------------------------------------------------------------||
// || Main
// ||------------------------------------------------------------------------------------------------||
func main() {
	// flags
	var (
		envPath string
		dryRun  bool
		table   string
		idCol   string
		reqCol  string
	)
	flag.StringVar(&envPath, "env", ".env", "Path to .env file")
	flag.BoolVar(&dryRun, "dry", false, "Dry run (no writes)")
	flag.StringVar(&table, "table", "zones", "Table name")
	flag.StringVar(&idCol, "idcol", "id_zone", "ID column")
	flag.StringVar(&reqCol, "reqcol", "zone_requirements", "Requirements column")
	flag.Parse()

	// load .env (if present)
	if envPath != "" {
		if _, err := os.Stat(envPath); err == nil {
			if err := loadDotEnv(envPath); err != nil {
				log.Fatalf("failed to load .env: %v", err)
			}
			log.Printf("Loaded env from %s", filepath.Clean(envPath))
		}
	}

	// build DSN from env
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || name == "" {
		log.Fatal("missing DB envs (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4,utf8&collation=utf8mb4_0900_ai_ci",
		user, pass, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ping db: %v", err)
	}
	log.Printf("Connected to MySQL %s:%s/%s (dry=%v)", host, port, name, dryRun)

	// read all rows
	sel := fmt.Sprintf("SELECT %s, %s FROM %s", idCol, reqCol, table)
	rows, err := db.Query(sel)
	if err != nil {
		log.Fatalf("select: %v", err)
	}
	defer rows.Close()

	type row struct {
		id  int64
		req sql.NullString
	}

	var scanned, updated int
	for rows.Next() {
		var r row
		if err := rows.Scan(&r.id, &r.req); err != nil {
			log.Fatalf("scan: %v", err)
		}
		scanned++

		old := strings.TrimSpace(r.req.String)
		newVal, changed, warns := mapValue(old)

		for _, w := range warns {
			log.Printf("WARN id=%d: %s", r.id, w)
		}

		if !changed {
			continue
		}

		log.Printf("id=%d: '%s' -> '%s'", r.id, old, newVal)

		if dryRun {
			continue
		}

		if err := updateOne(db, table, reqCol, idCol, newVal, r.id); err != nil {
			log.Fatalf("update id=%d: %v", r.id, err)
		}
		updated++
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("row err: %v", err)
	}

	log.Printf("Scanned %d rows; %d updated.", scanned, updated)
}

// ||------------------------------------------------------------------------------------------------||
// || Write Back
// ||------------------------------------------------------------------------------------------------||
func updateOne(db *sql.DB, table, reqCol, idCol, newVal string, id int64) error {
	if strings.TrimSpace(newVal) == "" {
		// avoid blanking out accidentally
		return errors.New("refusing to write empty zone_requirements")
	}
	q := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?", table, reqCol, idCol)
	_, err := db.Exec(q, newVal, id)
	return err
}

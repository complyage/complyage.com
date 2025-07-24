package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/helpers"
	"base/loaders"
	"base/responses"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Serves the HTML file with dynamic replacements
//||------------------------------------------------------------------------------------------------||

var (
	gateHTMLCache string
)

//||------------------------------------------------------------------------------------------------||
//|| Init
//||------------------------------------------------------------------------------------------------||

func init() {
	if UseInMemory {
		b, err := ioutil.ReadFile("assets/gate.html")
		if err != nil {
			panic(fmt.Sprintf("failed to preload gate.html: %v", err))
		}
		gateHTMLCache = string(b)
	}
}

// ||------------------------------------------------------------------------------------------------||
// || Serves the HTML file with dynamic replacements
// ||------------------------------------------------------------------------------------------------||

func ServeAgeGateHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Serves the HTML file with dynamic replacements
	//||------------------------------------------------------------------------------------------------||

	apiKey := r.URL.Query().Get("apiKey")
	if apiKey == "" {
		responses.ErrorHTML(w, "apiKey is required")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Site By API Key
	//||------------------------------------------------------------------------------------------------||

	site := loaders.GetSiteByPublic(apiKey)
	if site == nil {
		responses.ErrorHTML(w, "Invalid apiKey")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Zone By IP
	//||------------------------------------------------------------------------------------------------||

	ip := helpers.GetClientIP(r)
	loc, err := helpers.GetLocationByIP(ip)
	if err != nil {
		responses.ErrorHTML(w, "Geo lookup failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Zone
	//||------------------------------------------------------------------------------------------------||

	zone, found := loaders.FindZoneByLocation(loc.City, loc.State)
	if !found || zone == nil {
		responses.NoEnforce(w)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Serves the HTML file with dynamic replacements
	//||------------------------------------------------------------------------------------------------||

	lang := "en"
	if al := r.Header.Get("Accept-Language"); al != "" {
		primary := strings.SplitN(al, ",", 2)[0]
		if code := strings.SplitN(primary, "-", 2)[0]; code != "" {
			lang = strings.ToLower(code)
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Translations
	//||------------------------------------------------------------------------------------------------||

	translations, _ := loaders.GetTranslations(lang)

	//||------------------------------------------------------------------------------------------------||
	//|| Get The Template
	//||------------------------------------------------------------------------------------------------||

	tpl := ""
	if UseInMemory {
		tpl = gateHTMLCache
	} else {
		b, err := ioutil.ReadFile("assets/gate.html")
		if err != nil {
			http.Error(w, "template not found", http.StatusInternalServerError)
			fmt.Printf("read template error: %v\n", err)
			return
		}
		tpl = string(b)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate Template  Markers
	//||------------------------------------------------------------------------------------------------||

	vars := make(map[string]string, len(translations)+5)
	//||------------------------------------------------------------------------------------------------||
	//|| Add Translations
	//||------------------------------------------------------------------------------------------------||
	for key, val := range translations {
		vars[key] = val
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Add Translations
	//||------------------------------------------------------------------------------------------------||
	vars["IPADDRESS"] = ip
	vars["SIGNUPURL"] = os.Getenv("COMPLYAGE_UI_URL") + "/signup?apiKey=" + apiKey
	vars["OAUTHURL"] = os.Getenv("COMPLYAGE_UI_URL") + "/oauth?apiKey=" + apiKey
	vars["EXITURL"] = os.Getenv("COMPLYAGE_UI_URL") + "/exit"
	//||------------------------------------------------------------------------------------------------||
	//|| Logo URL
	//||------------------------------------------------------------------------------------------------||
	hashSecret := os.Getenv("MINIO_HASH")
	host := os.Getenv("VITE_COMPLYAGE_MINIO_URL")
	bucket := os.Getenv("VITE_MINIO_BUCKET")
	hashInput := fmt.Sprintf("%s%d%s", hashSecret, site.IDSite, hashSecret)
	hashBytes := sha256.Sum256([]byte(hashInput))
	hashHex := fmt.Sprintf("%d_%s", site.IDSite, hex.EncodeToString(hashBytes[:]))
	logoURL := fmt.Sprintf("%s/%s/sites/logos/%s.webp", host, bucket, hashHex)
	//||------------------------------------------------------------------------------------------------||
	//|| City/State
	//||------------------------------------------------------------------------------------------------||
	var parts []string
	if zone.ZoneState != nil && *zone.ZoneState != "" {
		parts = append(parts, *zone.ZoneState)
	}
	if zone.ZoneCountry != nil && *zone.ZoneCountry != "" {
		parts = append(parts, *zone.ZoneCountry)
	}
	vars["LOCATION"] = strings.Join(parts, ", ")
	//||----------------------------------------------------------------------------------------------||
	//|| Effective Date
	//||----------------------------------------------------------------------------------------------||
	if zone.ZoneEffective != nil {
		vars["EFFECTIVE"] = helpers.FormatMonthYear(*zone.ZoneEffective, lang)
	} else {
		vars["EFFECTIVE"] = ""
	}
	//||----------------------------------------------------------------------------------------------||
	//|| Site URL/Zone
	//||----------------------------------------------------------------------------------------------||
	vars["SITE_URL"] = site.SiteURL
	vars["SITE_NAME"] = site.SiteName
	vars["SITE_LOGO"] = logoURL
	vars["ZONE_ID"] = strconv.Itoa(int(zone.IDZone))
	//||----------------------------------------------------------------------------------------------||
	//|| Local
	//||----------------------------------------------------------------------------------------------||
	vars["COMPLYAGE_UI_URL"] = os.Getenv("SITE_URL")
	vars["COMPLYAGE_CLIENT_URL"] = os.Getenv("LOCAL_URL")
	//||----------------------------------------------------------------------------------------------||
	//|| Single pass replacement of [%%KEY%%] → value
	//||----------------------------------------------------------------------------------------------||
	for key, val := range vars {
		marker := fmt.Sprintf("[%%%%%s%%%%]", key)
		tpl = strings.ReplaceAll(tpl, marker, val)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Replace Zone Markers
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(tpl))

}

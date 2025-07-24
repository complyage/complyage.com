// translate_i18n.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// Response structs for Google Translate API
// Contains the translated text
// @see https://cloud.google.com/translate/docs/reference/rest/v2/translate
type translateResponse struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func main() {
	// load .env
	_ = godotenv.Load()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	langsEnv := os.Getenv("LANGUAGES")
	if apiKey == "" || langsEnv == "" {
		fmt.Fprintln(os.Stderr, "❌ .env must contain GOOGLE_API_KEY and LANGUAGES (e.g. es,fr,de)")
		os.Exit(1)
	}
	targets := strings.Split(langsEnv, ",")

	// determine project root (cwd)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "❌ cannot get working dir:", err)
		os.Exit(1)
	}

	// paths to i18n directory
	i18nDir := filepath.Join(wd, "i18n")
	enPath := filepath.Join(i18nDir, "en.json")

	// read English source
	enBytes, err := ioutil.ReadFile(enPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed reading %s: %v\n", enPath, err)
		os.Exit(1)
	}
	var en map[string]string
	if err := json.Unmarshal(enBytes, &en); err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed parsing %s: %v\n", enPath, err)
		os.Exit(1)
	}

	// ensure i18n directory exists before writing
	if err := os.MkdirAll(i18nDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "❌ cannot ensure i18n dir: %v\n", err)
		os.Exit(1)
	}

	// translate each language
	for _, lang := range targets {
		lang = strings.TrimSpace(lang)
		if lang == "" || lang == "en" {
			continue
		}
		fmt.Printf("🔄 Translating to %s...\n", lang)
		out := make(map[string]string, len(en))
		for key, text := range en {
			tr, err := translateText(apiKey, text, lang)
			if err != nil {
				fmt.Fprintf(os.Stderr, "   ⚠️ error translating %s: %v\n", key, err)
				tr = text
			}
			out[key] = tr
		}
		// write result
		destPath := filepath.Join(i18nDir, lang+".json")
		data, _ := json.MarshalIndent(out, "", "  ")
		if err := ioutil.WriteFile(destPath, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "❌ failed writing %s: %v\n", destPath, err)
			continue
		}
		fmt.Printf("✅ Written %s\n", destPath)
	}
}

// translateText calls Google Translate API
func translateText(apiKey, text, target string) (string, error) {
	endpoint := "https://translation.googleapis.com/language/translate/v2"
	params := url.Values{}
	params.Set("key", apiKey)
	params.Set("q", text)
	params.Set("target", target)
	params.Set("format", "text")

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("api error: %s", string(body))
	}
	var tr translateResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", err
	}
	if len(tr.Data.Translations) == 0 {
		return "", fmt.Errorf("no translations returned")
	}
	return tr.Data.Translations[0].TranslatedText, nil
}

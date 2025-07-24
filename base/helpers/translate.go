package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type translateResponse struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func TranslateI18nHandler(w http.ResponseWriter, r *http.Request) {

	if os.Getenv("ENV_MODE") == "production" {
		http.Error(w, "This endpoint is not available in production mode", http.StatusForbidden)
		return
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")
	langsEnv := os.Getenv("LANGUAGES")
	if apiKey == "" || langsEnv == "" {
		http.Error(w, "Missing GOOGLE_API_KEY or LANGUAGES in .env", http.StatusInternalServerError)
		return
	}
	targets := strings.Split(langsEnv, ",")

	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Unable to get working directory", http.StatusInternalServerError)
		return
	}

	assetsDir := filepath.Join(wd, "assets")
	i18nDir := filepath.Join(assetsDir, "lang")
	enPath := filepath.Join(assetsDir, "en.json")

	enBytes, err := os.ReadFile(enPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read en.json: %v", err), http.StatusInternalServerError)
		return
	}

	var en map[string]string
	if err := json.Unmarshal(enBytes, &en); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse en.json: %v", err), http.StatusInternalServerError)
		return
	}

	if err := os.MkdirAll(i18nDir, 0755); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create i18n dir: %v", err), http.StatusInternalServerError)
		return
	}

	result := make(map[string]map[string]string)

	for _, lang := range targets {
		lang = strings.TrimSpace(lang)
		if lang == "" || lang == "en" {
			continue
		}

		out := make(map[string]string, len(en))
		for key, text := range en {
			tr, err := translateText(apiKey, text, lang)
			if err != nil {
				out[key] = text // fallback
			} else {
				out[key] = tr
			}
		}

		destPath := filepath.Join(i18nDir, lang+".json")
		if data, err := json.MarshalIndent(out, "", "  "); err == nil {
			_ = os.WriteFile(destPath, data, 0644)
		}

		result[lang] = out
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

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
		body, _ := io.ReadAll(resp.Body)
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

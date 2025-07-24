package loaders

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

const defaultLang = "en"

//||------------------------------------------------------------------------------------------------||
//|| In‑memory map of translations
//||   key: language code (filename without “.json”)
//||   val: map of translation keys→values
//||------------------------------------------------------------------------------------------------||

var Translations map[string]map[string]string

//||------------------------------------------------------------------------------------------------||
//|| LoadTranslations
//||------------------------------------------------------------------------------------------------||

func LoadTranslations() {
	dir := "assets/lang"
	Translations = make(map[string]map[string]string)

	// 1) Load default lang
	defPath := filepath.Join("assets", defaultLang+".json")
	data, err := ioutil.ReadFile(defPath)
	if err != nil {
		log.Fatalf("Default translation file %q missing or unreadable: %v", defPath, err)
	}
	var defMap map[string]string
	if err := json.Unmarshal(data, &defMap); err != nil {
		log.Fatalf("Default translation JSON invalid (%q): %v", defPath, err)
	}
	Translations[defaultLang] = defMap

	// 2) Load any other languages
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		log.Printf("Warning: could not scan translations directory %q: %v", dir, err)
		return
	}
	for _, path := range files {
		name := filepath.Base(path)
		key := strings.TrimSuffix(name, filepath.Ext(name))
		if key == defaultLang {
			continue // already loaded
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("Skipping %q: read error: %v", path, err)
			continue
		}
		var m map[string]string
		if err := json.Unmarshal(b, &m); err != nil {
			log.Printf("Skipping %q: JSON parse error: %v", path, err)
			continue
		}
		Translations[key] = m
	}

	log.Printf("Loaded %d translation files (default=%q)", len(Translations), defaultLang)
}

//||------------------------------------------------------------------------------------------------||
//|| GetTranslations
//||------------------------------------------------------------------------------------------------||

func GetTranslations(lang string) (map[string]string, bool) {
	if m, ok := Translations[lang]; ok {
		return m, true
	}
	// fallback to default
	return Translations[defaultLang], true
}

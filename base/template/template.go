package template

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/loaders"
	"fmt"
	"net/http"
	"os"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Vars
//||------------------------------------------------------------------------------------------------||

var TemplateFiles []TemplateFile

//||------------------------------------------------------------------------------------------------||
//|| Template Structures
//||------------------------------------------------------------------------------------------------||

type TemplateMarkers struct {
	Marker string
	Value  string
}

type TemplateFile struct {
	Name string
	Path string
	Data string
}

type TemplateInstance struct {
	Name      string
	Data      string
	Markers   []TemplateMarkers
	Add       func(marker string, value string) TemplateInstance
	Compile   func() (string, error)
	Translate func(r *http.Request) TemplateInstance
}

//||------------------------------------------------------------------------------------------------||
//|| Register Template File
//||------------------------------------------------------------------------------------------------||

func Register(name string, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Template file not found: " + path)
	}

	TemplateFiles = append(TemplateFiles, TemplateFile{
		Name: name,
		Path: path,
		Data: string(data),
	})
}

//||------------------------------------------------------------------------------------------------||
//|| Create a New Instance from Registered Template
//||------------------------------------------------------------------------------------------------||

func Create(name string) TemplateInstance {
	var tmplData string
	var path string

	//||------------------------------------------------------------------------------------------------||
	//|| Find registered template path
	//||------------------------------------------------------------------------------------------------||
	for _, t := range TemplateFiles {
		if t.Name == name {
			tmplData = t.Data
			path = t.Path
			break
		}
	}
	if path == "" {
		panic(fmt.Sprintf("Template not found: %s", name))
	}

	//||------------------------------------------------------------------------------------------------||
	//|| In DEV mode, reload template file from disk
	//||------------------------------------------------------------------------------------------------||
	if os.Getenv("ENV_MODE") == "development" {
		data, err := os.ReadFile(path)
		if err == nil {
			tmplData = string(data) // override cached content
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Instance
	//||------------------------------------------------------------------------------------------------||

	instance := TemplateInstance{
		Name:    name,
		Data:    tmplData,
		Markers: []TemplateMarkers{},
	}

	instance.Add = func(marker string, value string) TemplateInstance {
		instance.Markers = append(instance.Markers, TemplateMarkers{
			Marker: marker,
			Value:  value,
		})
		return instance
	}

	instance.Compile = func() (string, error) {
		data := instance.Data
		for _, m := range instance.Markers {
			data = replaceMarker(data, m.Marker, m.Value)
		}
		return data, nil
	}

	instance.Translate = func(r *http.Request) TemplateInstance {
		lang := "en"
		if al := r.Header.Get("Accept-Language"); al != "" {
			primary := strings.SplitN(al, ",", 2)[0]
			if code := strings.SplitN(primary, "-", 2)[0]; code != "" {
				lang = strings.ToLower(code)
			}
		}
		translations, _ := loaders.GetTranslations(lang)
		for key, val := range translations {
			instance.Markers = append(instance.Markers, TemplateMarkers{
				Marker: key,
				Value:  val,
			})
		}
		return instance
	}

	return instance
}

//||------------------------------------------------------------------------------------------------||
//|| Replace Marker in String
//||------------------------------------------------------------------------------------------------||

func replaceMarker(data string, marker string, value string) string {
	placeholder := fmt.Sprintf("[%%%%%s%%%%]", marker)
	return strings.ReplaceAll(data, placeholder, value)
}

//||------------------------------------------------------------------------------------------------||
//|| Helper
//||------------------------------------------------------------------------------------------------||

func ReplaceMarker(data string, marker string, value string) string {
	return replaceMarker(data, marker, value)
}

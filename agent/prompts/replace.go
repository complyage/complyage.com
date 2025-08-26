package prompts

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Static
//||------------------------------------------------------------------------------------------------||

func Replace(alias, content1 string) string {
	tpl := template.Create(alias)
	tpl.Add("CONTENT1", content1)
	content, err := tpl.Compile()
	if err != nil {
		panic("Failed to compile static prompt: " + alias + " Error: " + err.Error())
	}
	return content
}

//||------------------------------------------------------------------------------------------------||
//|| Static
//||------------------------------------------------------------------------------------------------||

func ReplaceTwo(alias, content1, content2 string) string {
	tpl := template.Create(alias)
	tpl.Add("CONTENT1", content1)
	tpl.Add("CONTENT2", content2)
	content, err := tpl.Compile()
	if err != nil {
		panic("Failed to compile static prompt: " + alias + " Error: " + err.Error())
	}
	return content
}

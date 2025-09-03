package prompts

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"github.com/ralphferrara/aria/base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Static
//||------------------------------------------------------------------------------------------------||

func Replace(alias, content1 string) string {
	tpl := template.Create(alias, "en")
	tpl.Add("CONTENT1", content1)
	content := tpl.Compile()
	return content
}

//||------------------------------------------------------------------------------------------------||
//|| Static
//||------------------------------------------------------------------------------------------------||

func ReplaceTwo(alias, content1, content2 string) string {
	tpl := template.Create(alias, "en")
	tpl.Add("CONTENT1", content1)
	tpl.Add("CONTENT2", content2)
	content := tpl.Compile()
	return content
}

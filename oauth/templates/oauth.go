package templates

import (
	"net/http"

	"github.com/ralphferrara/aria/base/template"
)

func OAuth(r *http.Request) template.TemplateInstance {

	tpl := template.Create("oauth")
	tpl.Request(r)
	return tpl
}

package templates

import "github.com/ralphferrara/aria/base/template"

func SubPermissionHTML(code, statusClass string) string {
	tpl := template.Create("sub.permission")
	tpl.Add("PERMCODE", code)
	return tpl.Compile()
}

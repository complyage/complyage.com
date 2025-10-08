package templates

import (
	"os"

	"github.com/complyage/base/enforce"

	"github.com/complyage/base/scopes"
	"github.com/ralphferrara/aria/base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Sub
//||------------------------------------------------------------------------------------------------||

func SubPermissionHTML(s enforce.EnforcedScope, loggedIn bool, isAge bool) string {
	//||------------------------------------------------------------------------------------------------||
	//|| Sub
	//||------------------------------------------------------------------------------------------------||
	var tpl template.TemplateInstance
	if isAge {
		tpl = template.Create("sub.permission.age")
	} else {
		tpl = template.Create("sub.permission")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Sub
	//||------------------------------------------------------------------------------------------------||
	tpl.Add("PERMCODE", s.Code.String())
	tpl.Add("PERM_PERMCODE_TITLE", scopes.Title(s.Code.String()))
	tpl.Add("PERM_PERMCODE_DESC", scopes.Description(s.Code.String()))
	tpl.Add("PERM_PERMCODE_ICON", scopes.Icon(s.Code.String()))
	//||------------------------------------------------------------------------------------------------||
	//|| Button/Badge
	//||------------------------------------------------------------------------------------------------||
	if !loggedIn {
		btn := template.Create("verify.login")
		tpl.Add("STATUS", btn.Compile())
		tpl.Add("SHOWPREVIEW", "hidden")
		tpl.Add("MARKER_VERIFICATION_STATUS", "{{::DEFAULT:IS_NOT_VERIFIED}}")
		tpl.Add("CLASS_VERIFICATION_STATUS", "text-red-400")
	} else if s.Verified {
		tpl.Add("SHOWPREVIEW", "flex")
		tpl.Add("MARKER_VERIFICATION_STATUS", "{{::DEFAULT:IS_VERIFIED}}")
		tpl.Add("CLASS_VERIFICATION_STATUS", "text-yellow-400")
		if isAge {
			badge := template.Create("verify.use")
			tpl.Add("STATUS", badge.Compile())
		} else {
			badge := template.Create("verify.badge")
			tpl.Add("STATUS", badge.Compile())
		}
	} else {
		btn := template.Create("verify.button")
		tpl.Add("STATUS", btn.Compile())
		tpl.Add("SHOWPREVIEW", "hidden")
		tpl.Add("MARKER_VERIFICATION_STATUS", "{{::DEFAULT:IS_NOT_VERIFIED}}")
		tpl.Add("CLASS_VERIFICATION_STATUS", "text-red-400")
	}
	tpl.Data = tpl.Compile()
	tpl.Add("PERMCODE", s.Code.String())
	tpl.Add("VITE_COMPLYAGE_UI_URL", os.Getenv("VITE_COMPLYAGE_UI_URL"))
	tpl.Add("VITE_COMPLYAGE_API_URL", os.Getenv("VITE_COMPLYAGE_API_URL"))
	//||------------------------------------------------------------------------------------------------||
	//|| Return Compiled Template
	//||------------------------------------------------------------------------------------------------||
	return tpl.Compile()
}

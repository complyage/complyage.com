package templates

import (
	"fmt"
	"os"

	"github.com/complyage/base/enforce"

	"github.com/ralphferrara/aria/base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Sub
//||------------------------------------------------------------------------------------------------||

func SubAgeHTML(enforcement *enforce.Enforcement) string {
	//||------------------------------------------------------------------------------------------------||
	//|| Sub
	//||------------------------------------------------------------------------------------------------||
	tpl := template.Create("sub.age")
	//||------------------------------------------------------------------------------------------------||
	//|| Sub
	//||------------------------------------------------------------------------------------------------||
	tpl.Add("PERMCODE", "AGE")
	//||------------------------------------------------------------------------------------------------||
	//|| Of Age
	//||------------------------------------------------------------------------------------------------||
	if enforcement.Age.ZoneVerified {
		tpl.Add("MARKER_AGE_APPROVED", "{{::DEFAULT:CONGRATS_AGE_VERIFIED}}")
	} else {
		tpl.Add("MARKER_AGE_APPROVED", "{{::DEFAULT:AGE_VERIFICATION_REQUIRED}}")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| DOB/Age
	//||------------------------------------------------------------------------------------------------||
	tpl.Add("DOB", enforcement.Age.DOB.String())
	if enforcement.Age.DOB.Age() == 18 {
		tpl.Add("USERAGE", fmt.Sprintf("%d+", enforcement.Age.DOB.Age()))
	} else {
		tpl.Add("USERAGE", fmt.Sprintf("%d", enforcement.Age.DOB.Age()))
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Button/Badge
	//||------------------------------------------------------------------------------------------------||
	if enforcement.User.Level == 0 {
		btn := template.Create("verify.login")
		tpl.Add("STATUS", btn.Compile())
	} else if enforcement.Age.Verified {
		badge := template.Create("verify.badge")
		tpl.Add("STATUS", badge.Compile())
	} else {
		btn := template.Create("verify.button")
		tpl.Add("STATUS", btn.Compile())
	}
	tpl.Data = tpl.Compile()
	tpl.Add("PERMCODE", "AGE")
	tpl.Add("VITE_COMPLYAGE_UI_URL", os.Getenv("VITE_COMPLYAGE_UI_URL"))
	tpl.Add("VITE_COMPLYAGE_API_URL", os.Getenv("VITE_COMPLYAGE_API_URL"))
	//||------------------------------------------------------------------------------------------------||
	//|| Return Compiled Template
	//||------------------------------------------------------------------------------------------------||
	return tpl.Compile()
}

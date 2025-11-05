package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"net/http"
	"oauth/access"
	"oauth/templates"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/complyage/base/enforce"

	"github.com/ralphferrara/aria/base/template"
	"github.com/ralphferrara/aria/log"
)

// ||------------------------------------------------------------------------------------------------||
// || Serves the HTML file with dynamic replacements
// ||------------------------------------------------------------------------------------------------||

func ServeOAuthHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| OAuth
	//||------------------------------------------------------------------------------------------------||

	enforcement, err := enforce.LoadEnforcementScopes(r)
	if err != nil {
		fmt.Println(err.Error())
		templates.ErrorHTML(r, w, err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Serves the HTML file with dynamic replacements
	//||------------------------------------------------------------------------------------------------||

	tpl := template.Create("oauth")

	//||------------------------------------------------------------------------------------------------||
	//|| Status
	//||------------------------------------------------------------------------------------------------||

	var statusTemplate template.TemplateInstance
	if enforcement.User.Level > 0 && enforcement.User.Status == "ACTV" {
		statusTemplate = template.Create("status.loggedin")
		statusTemplate.Add("VITE_COMPLYAGE_UI_URL", os.Getenv("VITE_COMPLYAGE_UI_URL"))
		statusTemplate.Add("USERNAME", enforcement.User.Username)
	} else {
		statusTemplate = template.Create("status.loggedout")
		statusTemplate.Add("VITE_COMPLYAGE_UI_URL", os.Getenv("VITE_COMPLYAGE_UI_URL"))
	}
	tpl.Add("STATUS_LOGIN", statusTemplate.Compile())

	//||------------------------------------------------------------------------------------------------||
	//|| OAuth
	//||------------------------------------------------------------------------------------------------||

	htmlPermissions := strings.Builder{}
	for _, s := range enforcement.Scopes {

		fmt.Println("Processing Scope:", s.Code)
		log.PrettyPrint(s)

		//||------------------------------------------------------------------------------------------------||
		//|| Permissions Template Markers
		//||------------------------------------------------------------------------------------------------||

		html := templates.SubPermissionHTML(s, enforcement.User.Level > 0, false)

		//||------------------------------------------------------------------------------------------------||
		//|| Add to Main HTML
		//||------------------------------------------------------------------------------------------------||

		htmlPermissions.WriteString(html)
	}

	tpl.Data = tpl.Compile()

	//||------------------------------------------------------------------------------------------------||
	//|| Age
	//||------------------------------------------------------------------------------------------------||

	ageCheck := templates.SubAgeHTML(&enforcement)
	tpl.Add("AGECHECK", ageCheck)

	//||------------------------------------------------------------------------------------------------||
	//|| Cache Bust
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("CACHEBUST", strconv.FormatInt(time.Now().Unix(), 10))

	//||------------------------------------------------------------------------------------------------||
	//|| User
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("USERNAME", enforcement.User.Username)

	//||------------------------------------------------------------------------------------------------||
	//|| Permissions
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("PERMISSIONS", htmlPermissions.String())

	//||------------------------------------------------------------------------------------------------||
	//|| Site Data
	//||------------------------------------------------------------------------------------------------||

	if enforcement.Site.Logo == "" {
		enforcement.Site.Logo = "/static/img/complyage-w.webp"
	}
	tpl.Add("SITE_NAME", enforcement.Site.Name)
	tpl.Add("SITE_LOGO", os.Getenv("VITE_COMPLYAGE_MINIO_URL")+"/sites/"+enforcement.Site.Logo)
	tpl.Add("SITE_URL", enforcement.Site.URL)

	//||------------------------------------------------------------------------------------------------||
	//|| Statuses
	//||------------------------------------------------------------------------------------------------||

	loginStatus := "loggedOut"
	if enforcement.User.Level > 0 && enforcement.User.Status == "ACTV" {
		loginStatus = "loggedIn"
	}
	tpl.Add("LOGINSTATUS", loginStatus)

	//||------------------------------------------------------------------------------------------------||
	//|| Zone
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("IP", enforcement.Zone.IPAddress)
	tpl.Add("REGION", enforcement.Zone.Region)
	tpl.Add("COUNTRY", enforcement.Zone.Country)
	if enforcement.Zone.Enforced {
		tpl.Add("ENFORCED", "Enforced")
	} else {
		tpl.Add("ENFORCED", "Not Enforced")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Footer
	//||------------------------------------------------------------------------------------------------||

	var footer template.TemplateInstance
	if !enforcement.Zone.Enforced {
		//||------------------------------------------------------------------------------------------------||
		//|| No Enforcement Zone - Shouldn't really happen
		//||------------------------------------------------------------------------------------------------||
		oa, err := access.Create(enforcement)
		if err != nil {
			templates.ErrorHTML(r, w, "Failed to create access key: "+err.Error())
			return
		}
		//||------------------------------------------------------------------------------------------------||
		//|| Footer Template
		//||------------------------------------------------------------------------------------------------||
		footer = template.Create("footer.notrequired")
		footer.Add("TOKEN", oa.Token)
		footer.Add("BYPASSKEY", oa.BypassKey)
	} else {
		if enforcement.User.Level == 0 {
			footer = template.Create("footer.loggedout")
		} else {
			if enforcement.Missing > 0 {
				footer = template.Create("footer.requires")
				footer.Add("MISSING", strconv.Itoa(enforcement.Missing))
			} else {
				if !enforcement.HasPrivate {
					footer = template.Create("footer.private")
				} else {
					//||------------------------------------------------------------------------------------------------||
					//|| We Are Ready To Go
					//||------------------------------------------------------------------------------------------------||
					oa, err := access.Create(enforcement)
					if err != nil {
						templates.ErrorHTML(r, w, "Failed to create access key: "+err.Error())
						return
					}
					//||------------------------------------------------------------------------------------------------||
					//|| We Are Ready To Go
					//||------------------------------------------------------------------------------------------------||
					footer = template.Create("footer.verified")
					footer.Add("TOKEN", oa.Token)
					footer.Add("AUTHORIZEKEY", oa.AuthorizeKey)
					footer.Add("DENYKEY", oa.DenyKey)
				}
			}
		}
	}
	tpl.Add("FOOTER", footer.Compile())
	tpl.Data = tpl.Compile()

	//||------------------------------------------------------------------------------------------------||
	//|| URLS
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("CLIENTID", enforcement.Site.ClientId)
	tpl.Add("VITE_COMPLYAGE_UI_URL", os.Getenv("VITE_COMPLYAGE_UI_URL"))
	tpl.Add("VITE_COMPLYAGE_API_URL", os.Getenv("VITE_COMPLYAGE_API_URL"))

	log.PrettyPrint(enforcement)

	//||------------------------------------------------------------------------------------------------||
	//|| Of Age Class
	//||------------------------------------------------------------------------------------------------||

	if enforcement.Age.ZoneVerified {
		tpl.Add("OFAGECLASS", "ofAge")
	} else {
		tpl.Add("OFAGECLASS", "notOfAge")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Replace Zone Markers
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(tpl.Compile()))

}

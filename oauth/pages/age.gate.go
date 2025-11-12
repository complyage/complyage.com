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
	"github.com/complyage/base/vpns"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/template"
	"github.com/ralphferrara/aria/log"
)

// ||------------------------------------------------------------------------------------------------||
// || Serves the HTML file with dynamic replacements
// ||------------------------------------------------------------------------------------------------||

func ServeAgeGateHandler(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("ServeAgeGateHandler")

	//||------------------------------------------------------------------------------------------------||
	//|| Load the Enforcement Data
	//||------------------------------------------------------------------------------------------------||

	enforcement, err := enforce.LoadEnforcementAge(r)
	if err != nil {
		fmt.Println(err.Error())
		templates.ErrorHTML(r, w, err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Site Data
	//||------------------------------------------------------------------------------------------------||

	if enforcement.Site.Logo == "" {
		enforcement.Site.Logo = "/static/img/complyage-w.webp"
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Template
	//||------------------------------------------------------------------------------------------------||

	tpl := template.Create("oauth")

	//||------------------------------------------------------------------------------------------------||
	//|| VPN Template
	//||------------------------------------------------------------------------------------------------||

	vpn := vpns.GetRandomVPN()
	vpnTemplate := template.Create("sub.vpn")
	if vpn != nil {
		vpnTemplate.Add("VPN_NAME", vpn.Name)
		vpnTemplate.Add("VPN_URL", vpn.URL)
		vpnTemplate.Add("VPN_BLURB", vpn.Blurb)
		tpl.Add("SUB_VPN", vpnTemplate.Compile())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Logged In / Logged Out
	//||------------------------------------------------------------------------------------------------||

	var statusTemplate template.TemplateInstance
	if enforcement.User.Level > 0 && enforcement.User.Status == "ACTV" {
		statusTemplate = template.Create("status.loggedin")
	} else {
		statusTemplate = template.Create("status.loggedout")
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
		html := templates.SubPermissionHTML(s, enforcement.User.Level > 0, true)
		//||------------------------------------------------------------------------------------------------||
		//|| Add to Main HTML
		//||------------------------------------------------------------------------------------------------||
		htmlPermissions.WriteString(html)
	}

	tpl.Data = tpl.Compile()

	//||------------------------------------------------------------------------------------------------||
	//|| Permissions
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("PERMISSIONS", htmlPermissions.String())

	//||------------------------------------------------------------------------------------------------||
	//|| Footer
	//||------------------------------------------------------------------------------------------------||

	var footer template.TemplateInstance
	shouldGenerateToken := false

	if !enforcement.Zone.Enforced {
		shouldGenerateToken = true
		footer = template.Create("footer.notrequired")
	} else {
		if enforcement.User.Level == 0 {
			footer = template.Create("footer.loggedout")
		} else {
			if enforcement.Missing > 0 {
				footer = template.Create("footer.notmet")
				footer.Add("MISSING", strconv.Itoa(enforcement.Missing))
			} else {
				if !enforcement.HasPrivate {
					footer = template.Create("footer.private")
				} else {
					shouldGenerateToken = true
					footer = template.Create("footer.verified.age")
				}
			}
		}
	}
	tpl.Add("FOOTER", footer.Compile())

	//||------------------------------------------------------------------------------------------------||
	//|| Zone Enforced
	//||------------------------------------------------------------------------------------------------||

	if enforcement.Zone.Enforced {
		tpl.Add("ENFORCED", "{{::DEFAULT:ENFORCED}}")
	} else {
		tpl.Add("ENFORCED", "{{::DEFAULT:NOT_ENFORCED}}")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Pre-compile the Template before adding the global markers
	//||------------------------------------------------------------------------------------------------||

	tpl.Data = tpl.Compile()

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Access Token
	//||------------------------------------------------------------------------------------------------||

	if shouldGenerateToken {
		oa, err := access.Create(enforcement)
		if err != nil {
			templates.ErrorHTML(r, w, "Failed to create access key: "+err.Error())
			return
		}
		tpl.Add("TOKEN", oa.Token)
		tpl.Add("PREVIEWKEY", oa.PreviewKey)
		tpl.Add("AUTHORIZEKEY", oa.AuthorizeKey)
		tpl.Add("DENYKEY", oa.DenyKey)
		tpl.Add("BYPASSKEY", oa.BypassKey)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Global Variables
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("CLIENTID", enforcement.Site.ClientId)
	tpl.Add("USERNAME", enforcement.User.Username)
	tpl.Add("VITE_COMPLYAGE_UI_URL", os.Getenv("VITE_COMPLYAGE_UI_URL"))
	tpl.Add("VITE_COMPLYAGE_API_URL", os.Getenv("VITE_COMPLYAGE_API_URL"))
	tpl.Add("SITE_NAME", enforcement.Site.Name)
	tpl.Add("SITE_LOGO", os.Getenv("VITE_COMPLYAGE_MINIO_URL")+"/sites/"+enforcement.Site.Logo)
	tpl.Add("SITE_URL", enforcement.Site.URL)
	tpl.Add("IP", enforcement.Zone.IPAddress)
	tpl.Add("REGION", enforcement.Zone.Region)
	tpl.Add("COUNTRY", enforcement.Zone.Country)
	tpl.Add("CACHEBUST", strconv.FormatInt(time.Now().Unix(), 10))

	//||------------------------------------------------------------------------------------------------||
	//|| Replace Zone Markers
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(tpl.Compile()))

}

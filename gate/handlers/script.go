package handlers

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/version"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Load the Templates
//||------------------------------------------------------------------------------------------------||

func init() {
	template.Register("complyage.dev.js", "loader/dist/complyage.js")
	template.Register("complyage.prod.js", "loader/dist/complyage.min.js")
}

//||------------------------------------------------------------------------------------------------||
//|| Generate the Script
//||------------------------------------------------------------------------------------------------||

func GenerateScriptHandler(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||
	clientId := r.URL.Query().Get("client_id")
	debug := false
	debugStr := "false"
	if r.URL.Query().Get("debug") == "true" {
		debug = true
		debugStr = "true"
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Create Template Container
	//||------------------------------------------------------------------------------------------------||
	var tpl template.TemplateInstance
	if debug {
		tpl = template.Create("complyage.dev.js")
	} else {
		tpl = template.Create("complyage.prod.js")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Markers
	//||------------------------------------------------------------------------------------------------||
	tpl.Add("CLIENT_ID", clientId)
	tpl.Add("DEBUG", debugStr)
	tpl.Add("VERSION", fmt.Sprintf("%.1f", version.VERSION_FLOAT))
	tpl.Add("GATEURL", app.Config.HTTP["gate"].URL)
	tpl.Add("OAUTHURL", app.Config.HTTP["oauth"].URL)
	//||------------------------------------------------------------------------------------------------||
	//|| Template
	//||------------------------------------------------------------------------------------------------||
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(tpl.Compile()))
}

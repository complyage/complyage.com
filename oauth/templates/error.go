package templates

import (
	"net/http"

	"github.com/ralphferrara/aria/base/template"
)

func ErrorHTML(r *http.Request, w http.ResponseWriter, errorMessage string) string {
	tmp := template.Create("error")
	tmp.Request(r)
	tmp.Add("ERROR_MESSAGE", errorMessage)
	return tmp.Show(w)
}

package static_pages

import (
    // "appengine"
    // "appengine/datastore"
    // "appengine/user"
    "html/template"
    "net/http"
    // "time"
)

func init() {
  http.HandleFunc("/info/", infoHandler)
}

var infoTmpl = template.Must(template.ParseFiles(
  "apps/main/templates/base.html",
  "apps/info/templates/info.html"
))

func infoHandler(w http.ResponseWriter, r *http.Request) {
  if err := infoTmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

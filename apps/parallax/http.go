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
  http.HandleFunc("/parallax/", infoHandler)
}

var infoTmpl = template.Must(template.ParseFiles(
  "apps/main/templates/base.html",
  "apps/parallax/templates/parallax.html"
))

func infoHandler(w http.ResponseWriter, r *http.Request) {
  if err := infoTmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

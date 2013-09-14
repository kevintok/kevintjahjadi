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
  http.HandleFunc("/parallax/", handler)
}

var parallaxTemplate = template.Must(template.ParseFiles(
  "apps/main/templates/base.html",
  "apps/parallax/templates/parallax.html",
))

func handler(w http.ResponseWriter, r *http.Request) {
  if err := parallaxTemplate.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

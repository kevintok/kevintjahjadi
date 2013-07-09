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
  http.HandleFunc("/products/", listHandler)
}

var listTmpl = template.Must(template.ParseFiles("static/templates/base.html",
  "apps/products/templates/list.html"))

func listHandler(w http.ResponseWriter, r *http.Request) {
  if err := listTmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

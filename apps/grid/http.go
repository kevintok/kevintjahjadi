package static_pages

import (
    "html/template"
    "net/http"
)

func init() {
  http.HandleFunc("/grid/", handler)
}

var gridTemplate = template.Must(template.ParseFiles(
  "apps/main/templates/base.html",
  "apps/grid/templates/grid.html",
))

func handler(w http.ResponseWriter, r *http.Request) {
  if err := gridTemplate.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

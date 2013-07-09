package test
// http://stackoverflow.com/questions/9573644/go-appengine-how-to-structure-templates-for-application
import (
    // "appengine"
    // "appengine/datastore"
    // "appengine/user"
    "html/template"
    "net/http"
    // "time"
)

// type Greeting struct {
//     Author  string
//     Content string
//     Date    time.Time
// }

var index = template.Must(template.ParseFiles(
  "static/templates/base.html",
  "static/templates/index.html",
))

var page404 = template.Must(template.ParseFiles(
  "static/templates/base.html",
  "static/templates/static_pages/404.html",
))

func init() {
    http.HandleFunc("/", root)
    // http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    if err := index.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        if err := page404.Execute(w, nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

// func root(w http.ResponseWriter, r *http.Request) {
//     c := appengine.NewContext(r)
//     q := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
//     greetings := make([]Greeting, 0, 10)
//     if _, err := q.GetAll(c, &greetings); err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     if err := guestbookTemplate.Execute(w, greetings); err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//     }
// }

// var guestbookTemplate = template.Must(template.New("book").Parse(guestbookTemplateHTML))

// const guestbookTemplateHTML = `
// <html>
//   <head>
//     <link type="text/css" rel="stylesheet" href="/stylesheets/main.css" />
//   </head>
//   <body>
//     {{range .}}
//       {{with .Author}}
//         <p><b>{{.}}</b> wrote:</p>
//       {{else}}
//         <p>An anonymous person wrote:</p>
//       {{end}}
//       <pre>{{.Content}}</pre>
//     {{end}}
//     <form action="/sign" method="post">
//       <div><textarea name="content" rows="3" cols="60"></textarea></div>
//       <div><input type="submit" value="Sign Guestbook"></div>
//     </form>
//   </body>
// </html>
// `

// func sign(w http.ResponseWriter, r *http.Request) {
//     c := appengine.NewContext(r)
//     g := Greeting{
//         Content: r.FormValue("content"),
//         Date:    time.Now(),
//     }
//     if u := user.Current(c); u != nil {
//         g.Author = u.String()
//     }
//     _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Greeting", nil), &g)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     http.Redirect(w, r, "/", http.StatusFound)
// }

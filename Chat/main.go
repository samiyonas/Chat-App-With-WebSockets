package main

import (
    "log"
    "net/http"
    "html/template"
    "sync"
)

type templateHandler struct {
    once sync.Once
    filename string
    tmpl *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.tmpl = template.Must(template.ParseFiles("templates/"+t.filename))
    })
    t.tmpl.Execute(w, nil)
}

func main() {
    http.Handle("/", &templateHandler{filename: "chat.html"})

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

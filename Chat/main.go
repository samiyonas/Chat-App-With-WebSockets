package main

import (
    "log"
    "net/http"
    "html/template"
    "sync"
)

// Handles templates  for the home page
type templateHandler struct {
    // Makes sure that template compiling is done once for all goroutines
    once sync.Once
    // template filename
    filename string
    // Template type
    tmpl *template.Template
}

// Using the above type as a handler by implementing http.Handler interface(ServeHTTP method)
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.tmpl = template.Must(template.ParseFiles("templates/"+t.filename))
    })
    t.tmpl.Execute(w, nil)
}

func main() {
    // Registers the handler for the home page which is templateHandler type
    http.Handle("/", &templateHandler{filename: "chat.html"})

    // Error handling
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

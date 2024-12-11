package main

import (
    "log"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`
    <html>
    <head>
    <title>Chat</title>
    </head>
    <body>
    Let's chat!
    </body>
    </html>
    `))
}

func main() {
    http.HandleFunc("/", Home)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

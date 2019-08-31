package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    
   http.HandleFunc("/about me", func(w http.ResponseWriter, r *http.Request){
            fmt.Fprintf(w, "Hello,Stop knowing about and wasting time do something usefull ", html.EscapeString(r.URL.Path))
})
   http.HandleFunc("/about Us", func(w http.ResponseWriter, r *http.Request){
            fmt.Fprintf(w, "Hello,Stop knowing about and wasting time do something useful and tech", html.EscapeString(r.URL.Path))
})


    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}

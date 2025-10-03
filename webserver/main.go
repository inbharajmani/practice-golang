package main

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    // Start the webserver on port 8440
    go func() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "Hello, you've reached the webserver on port 8440!")
        })
        fmt.Println("Starting webserver on port 8440")
        if err := http.ListenAndServe(":8440", nil); err != nil {
            fmt.Println("Error starting webserver:", err)
        }
    }()

    // Start the proxy server on port 8442
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Host == "example.com" {
            proxyURL, _ := url.Parse("http://localhost:8440")
            proxy := httputil.NewSingleHostReverseProxy(proxyURL)
            proxy.ServeHTTP(w, r)
        } else {
            w.WriteHeader(http.StatusForbidden)
            io.WriteString(w, "403 - Forbidden")
        }
    })
    fmt.Println("Starting proxy server on port 8442")
    if err := http.ListenAndServe(":8442", nil); err != nil {
        fmt.Println("Error starting proxy server:", err)
    }


    
}
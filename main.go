package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    now := time.Now().Format(time.RFC3339)
    fmt.Fprintf(w, "Current time: %s\n", now)
}

func main() {
    fmt.Println("Server started on :8080")
    fmt.Println("Current time:", time.Now().Format(time.RFC3339))
    http.HandleFunc("/", handler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  // API
  http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    fmt.Fprintln(w, "Hello from Go API3")
  })

  // 靜態檔
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)

  fmt.Println("Server running on port", port)
  http.ListenAndServe(":"+port, nil)
}


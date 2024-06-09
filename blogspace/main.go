package main

import (
  "log"
  "net/http"

  "github.com/schmeekygeek/wecart/api"
)

func main() {
  http.HandleFunc("/", api.SayHello)

  log.Println("Starting server on port :8080")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Println("Error starting server:", err)
  }
}

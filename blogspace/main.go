package main

import (
	"log"
	"net/http"

	"github.com/schmeekygeek/blogspace/api"
	"github.com/schmeekygeek/blogspace/data"
)

func main() {
  http.HandleFunc("/", api.SayHello)
  http.HandleFunc("/person/create", api.CreatePerson)

  log.Println("Starting server on port :8080")
  err := data.Init()
  if err != nil {
    panic(err)
  }
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Println("Error starting server:", err)
  }
}

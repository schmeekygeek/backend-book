package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/schmeekygeek/blogspace/data"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, "Method not allowed", 405)
    return
  }
  w.Header().Add("Content-Type", "application/json")
  w.Write([]byte("Hello"))
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w, "Method not allowed", 405)
    return
  }
  p := data.Person{}
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&p)
  if err != nil {
    log.Printf("Something went wrong: %v", err)
  }
  log.Printf("Got person to save: %v", p)
}

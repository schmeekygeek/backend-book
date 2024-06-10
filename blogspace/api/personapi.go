package api

import (
	"encoding/json"
	"fmt"
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
  fmt.Println(p)
  if err != nil {
    log.Printf("Error decoding data %v", err)
  }

  err = data.CreatePerson(&p)
  if err != nil {
    log.Printf("Error creating person %v", err)
  }
}

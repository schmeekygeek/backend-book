package api

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello"))
}

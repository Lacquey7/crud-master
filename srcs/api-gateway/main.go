package main

import (
	"api-gateway/internal/server"
	"net/http"
)

func HandleCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	r := http.NewServeMux()

	r.Handle("/", http.HandlerFunc(HandleCheck))

	s := server.New("localhost", "8080", r)
	err := s.Run()
	if err != nil {
		return
	}
}

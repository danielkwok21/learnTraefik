package main

import (
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[loginHandler] incoming request")
	if r.Method != http.MethodGet {
		log.Println("[loginHandler] method not allowed: ", r.Method)

		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Authorization") == "" {
		log.Println("[loginHandler] No authorization header found, returning 401")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("[loginHandler] Authorization header found, returning 200")

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/login", loginHandler)

	log.Println("listening on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

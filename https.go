package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/.well-known/acme-challenge/", challengeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var challenges = map[string]string{
	"<KEY>": "<VALUE>",
}

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	challenge := strings.Split(r.RequestURI, "/.well-known/acme-challenge/")[1]
	if responseToChallenge, ok := challenges[challenge]; !ok {
		http.Error(w, "Error", http.StatusNotFound)
		return
	} else {
		fmt.Fprintf(w, "%s", responseToChallenge)
	}
}

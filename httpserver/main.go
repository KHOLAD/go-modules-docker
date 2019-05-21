package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.Handle("/", logger(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "http package main")
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("URI: %s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

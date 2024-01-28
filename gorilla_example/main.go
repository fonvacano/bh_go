package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, fprintf := fmt.Fprintf(writer, "Hello man")
		if fprintf != nil {
			return
		}
	}).Methods("GET")

	r.HandleFunc("/users/{user:[a-z]+}", func(writer http.ResponseWriter, request *http.Request) {
		user := mux.Vars(request)["user"]
		_, fprintf := fmt.Fprintf(writer, "low case Hello %s\n", user)
		if fprintf != nil {
			return
		}
	}).Methods("GET")

	r.HandleFunc("/users/{user}", func(writer http.ResponseWriter, request *http.Request) {
		user := mux.Vars(request)["user"]
		_, fprintf := fmt.Fprintf(writer, "Hello %s\n", user)
		if fprintf != nil {
			return
		}
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}

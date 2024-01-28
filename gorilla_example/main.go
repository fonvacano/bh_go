package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

type badAuth struct {
	user string
	pass string
}

// do not use it, just demonstration
func (a *badAuth) serveHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")

	if user != a.user || pass != a.pass {
		http.Error(w, "auth error", 401)
		return
	}

	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)
	next(w, r)
}

func main() {
	r := mux.NewRouter()
	n := negroni.Classic()
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

	n.UseHandler(r)

	http.ListenAndServe(":8080", n)
}

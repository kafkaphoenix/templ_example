package main

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
)

func FuncTimeHandler(now func() time.Time) TimeHandler {
	return TimeHandler{Now: now}
}

type TimeHandler struct {
	Now func() time.Time
}

func (nh TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timeComponent(nh.Now()).Render(r.Context(), w)
}

func FuncHelloHandler() HelloHandler {
	return HelloHandler{}
}

type HelloHandler struct {
}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	helloComponent(name).Render(r.Context(), w)
}

func main() {
	http.Handle("/", FuncTimeHandler(time.Now))
	http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))
	http.Handle("/hello/{name}", FuncHelloHandler())

	http.ListenAndServe(":8080", nil)
}

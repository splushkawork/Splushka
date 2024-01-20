package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("ui/html/*.html"))
	mux := mux.NewRouter()
	mux.HandleFunc("/", landing).Methods("GET")
	mux.HandleFunc("/forum", forum).Methods("GET")
	mux.HandleFunc("/account", account).Methods("GET")

	fs := http.FileServer(http.Dir("./ui/static"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	log.Println("Запуск веб-сервера на http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

/*
var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("ui/html/*.html"))

	log.Println("Server will start at http://localhost:8080/")

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/forum", forum).Methods("GET")
	r.HandleFunc("/adviсe", adviсe).Methods("GET")

	fs := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("ui/html/*.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", landing)
	mux.HandleFunc("/forum", forum)
	mux.HandleFunc("/adviсe", adviсe)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск веб-сервера на http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

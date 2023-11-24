package main

import (
    //"fmt"
	"html/template" 
    "log"
    "net/http"
    //"strconv"
)

var tpl *template.Template


func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home_page.html",
		"./ui/html/base_layout.html",
		"./ui/html/footer_partial.html",
		"./ui/html/header_partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
 
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func forum(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path!= "/forum" {
        http.NotFound(w, r)
        return
    }

    files := []string{
		"./ui/html/forum.html",
        "./ui/html/base_layout.html",
        "./ui/html/footer_partial.html",
        "./ui/html/header_partial.html",
    }
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
 
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}}
 
func adviсe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path!= "/adviсe" {
        http.NotFound(w, r)
        return
    }

    files := []string{
		"./ui/html/adviсe.html",
        "./ui/html/base_layout.html",
        "./ui/html/footer_partial.html",
        "./ui/html/header_partial.html",
    }
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
 
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

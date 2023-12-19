package main

import (
	//"fmt"

	"net/http"
	//"strconv"
)

func landing(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "landing.html", r)

}

func forum(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "forum.html", r)
}

func adviсe(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "adviсe.html", r)
}

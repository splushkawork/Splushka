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

func account(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("type") == "signup" {
		templates.ExecuteTemplate(w, "SignUp.html", r)
	} else if r.URL.Query().Get("type") == "login" {
		templates.ExecuteTemplate(w, "LogIn.html", r)
	} else {
		w.Write([]byte("<h1>Страница не найдена</h1>"))
	}

}

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

func SignUp(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "SignUp.html", r)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "SignIn.html", r)
}

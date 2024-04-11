package main

import (
	//"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var templates *template.Template

/*type product struct{
    id int
    model string
    company string
    price int
}*/

func main() {
	templates = template.Must(template.ParseGlob("ui/html/*.html"))
	mux := mux.NewRouter()
	mux.HandleFunc("/", landing).Methods("GET")
	mux.HandleFunc("/forum", forum).Methods("GET")
	mux.HandleFunc("/account", account).Methods("GET")

	fs := http.FileServer(http.Dir("./ui/static"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	/*connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}*/

	log.Println("Запуск веб-сервера на http://localhost:8000")
	errServ := http.ListenAndServe(":8000", mux)
	log.Fatal(errServ)
}

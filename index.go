package main

import (
	"os"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Description string
}

func webHandler(w http.ResponseWriter, r *http.Request){
	td := Page{"ConfigChecker", "Cet outil permet de vérifier vos configurations !"}

  	t, err := template.ParseFiles("templates.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, td)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":456", nil))
}
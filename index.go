package main

import (
	"html/template"
	"net/http"
	"fmt"
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
	fmt.Println("Server Up and Running")
	http.HandleFunc("/", webHandler)
    	http.ListenAndServe(":8080", nil)
}

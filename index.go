package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type Home struct {
	Title string
	Description string
}

type Verif struct {
	Title string
	Description string
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	td := Home{"ConfigChecker", "Cet outil permet de vérifier vos configurations !"}

  	t, err := template.ParseFiles("templates.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, td)
	if err != nil {
		panic(err)
	}
}

func verifHandler(w http.ResponseWriter, r *http.Request){
	td := Verif{"Vérification de votre configuration","",""}
}

func main() {

	fmt.Println("Server Up and Running ...")
	http.HanndlerFunc("/",homeHandler)
    http.ListenAndServe( "0.0.0.0:8181",nil)
}

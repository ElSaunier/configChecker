package main

import (
	"html/template"
	"net/http"
	"fmt"

	"https://github.com/gorilla/mux"
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

	/* Router ?*/
	r := mux.newRouter()
	r.Handler(homeHandler)
	r.methods("POST").Handler(verifHandler)

	fmt.Println("Server Up and Running ...")
	srv := &http.Server{Addr: "0.0.0.0:8181", Handler: r}
    srv.ListenAndServe()
}

package main

import (
	"html/template"
	"net/http"
	"log"
	"configChecker"

	"github.com/bmizerany/pat"
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

func createTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
	  log.Println(err)
	  http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
  
	err = tmpl.Execute(w, data)
	if err != nil {
	  log.Println(err)
	  http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
  }  

func homeHandler(w http.ResponseWriter, r *http.Request){
	createTemplate(w, "templates/home.html", nil)
}

func sendHandler(w http.ResponseWriter, r *http.Request){

}

func verifHandler(w http.ResponseWriter, r *http.Request){
	createTemplate(w, "templates/home.html", nil)
}

func main() {

	mux := pat.New()
	mux.Get("/",http.HandlerFunc(homeHandler))
	mux.Post("/",http.HandlerFunc(sendHandler))
	mux.Get("/verif",http.HandlerFunc(verifHandler))

	var cfg ConfigFile
	cfg.Content = "Content"

	Validate(cfg)

	log.Println("Server Up and Running ...")
	err := http.ListenAndServe("0.0.0.0:8181",mux)
	if err != nil {
		log.Fatal(err)
	}
}

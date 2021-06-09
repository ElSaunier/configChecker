package main

import (
	"html/template"
	"net/http"
	"log"

	"github.com/bmizerany/pat"
)

type Home struct {
	Title string
	Description string
}

type Verif struct {
	Title string
	Content string
	Result string
	Valid bool

}

var first Home
var res Verif

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
	first.Title = "ConfigChecker"
	first.Description = "Outil de vérification de vos configurations Promtool et Alertmanagertool"
	createTemplate(w, "templates/home.html", first)
}

func sendHandler(w http.ResponseWriter, r *http.Request){
		var cfg ConfigFile
		cfg.Content = r.PostFormValue("config")

		if r.PostFormValue("identifier") == 1
			res.Content, res.Result, res.Valid = cfg.ValidatePromtool()
		else
			res.Content, res.Result, res.Valid = cfg.ValidateAlertManager()

		http.Redirect(w, r, "/verif", http.StatusSeeOther)
}

func verifHandler(w http.ResponseWriter, r *http.Request){
	res.Title = "Résultat de la vérification"
	createTemplate(w, "templates/verif.html", res)
}




func main() {

	mux := pat.New()
	mux.Get("/",http.HandlerFunc(homeHandler))
	mux.Post("/",http.HandlerFunc(sendHandler))
	mux.Get("/verif",http.HandlerFunc(verifHandler))

	
	log.Println("Server Up and Running ...")
	err := http.ListenAndServe("0.0.0.0:8181",mux)
	
	if err != nil {
		log.Fatal(err)
	}
}

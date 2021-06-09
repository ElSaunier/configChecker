package main

import (
	"html/template"
	"net/http"
	"log"
	"strconv"

	"github.com/bmizerany/pat"
)

type Verif struct {
	Content string
	Result string
}

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
	createTemplate(w, "templates/home.html", nil)
}

func sendHandler(w http.ResponseWriter, r *http.Request){
		var cfg ConfigFile
		cfg.Content = r.PostFormValue("config")
		i, err := strconv.Atoi(r.PostFormValue("idenfitier")) 
		if err != nil{
			log.Fatal("Impossible to define used tool")
		}
		if i == 1 {
			res.Content, res.Result = cfg.ValidatePromtool()
		} else {
			res.Content, res.Result = cfg.ValidateAlertManager()
		}

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

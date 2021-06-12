package main

import (
	"html/template"
	"net/http"
	"log"
	"strconv"

	"github.com/gorilla/mux"
)

type Verif struct {
	Content string
	Result string
	Color string
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
				
		if cfg.Content != "" {
			i, err := strconv.Atoi(r.PostFormValue("identifier"))
			if err != nil{
				log.Fatal("Can't define used tool")
			}
    		log.Println("[+] Checking ...")
			if i == 1 {
				res.Content, res.Result, res.Color = cfg.ValidatePromtool()
			} else {
				res.Content, res.Result, res.Color = cfg.ValidateAlertManager()
			}
			log.Println("[-] Checking Finished")
		} else {
			res.Content = ""
			res.Result = "Empty uploaded file"
			res.Color = "red"
		}
		http.Redirect(w, r, "/verif", http.StatusSeeOther)
}

func verifHandler(w http.ResponseWriter, r *http.Request){
	createTemplate(w, "templates/verif.html", res)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/",homeHandler).Methods("GET")
	r.HandleFunc("/",sendHandler).Methods("POST")
	r.HandleFunc("/verif",verifHandler).Methods("GET")
	
	log.Println("Server Up and Running ...")
	err := http.ListenAndServe("0.0.0.0:8181",r)
	
	if err != nil {
		log.Fatal(err)
	}
}

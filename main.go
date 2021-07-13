package main

import (
	"html/template"
	"net/http"
	"log"
	"strconv"
	"io/ioutil"
	"fmt"

	"github.com/gorilla/mux"
)

type Verif struct {
	Content string
	Result string
	Color string
	Identifier int
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
	createTemplate(w, "/app/templates/home.html", nil)
}

func sendHandler(w http.ResponseWriter, r *http.Request){
		var cfg ConfigFile
		cfg.Content = r.PostFormValue("config")
				
		if cfg.Content != "" {
			i, err := strconv.Atoi(r.PostFormValue("identifier"))
			if err != nil{
				log.Fatal("Can't define used tool")
			}
			res.Identifier = i;
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
			res.Identifier = 0
		}
		http.Redirect(w, r, "/verif", http.StatusSeeOther)
}

func verifHandler(w http.ResponseWriter, r *http.Request){
	createTemplate(w, "/app/templates/verif.html", res)
}

func apiHandler(w http.ResponseWriter, r *http.Request){
	var cfg ConfigFile

	// Taille de 10MB maximum
    r.ParseMultipartForm(10 << 20)

    file, _, err := r.FormFile("config")
	if r.FormFile("config") == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing config file")
		return
	}
	
    if err != nil {
        log.Println("Error Retrieving the File")
        log.Println(err)
        return
    }
    defer file.Close()

    fileBytes, err2 := ioutil.ReadAll(file)
    if err2 != nil {
        log.Println(err2)
    }

	fileString := string(fileBytes)
	cfg.Content = fileString


	if cfg.Content != "" {
		i, err := strconv.Atoi(r.PostFormValue("identifier"))
		if err != nil{
			log.Fatal("Can't define used tool")
		}
		res.Identifier = i;

		if i == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Missing identifier (1 for Promtool | 2 for Amtool")
			return
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
		res.Identifier = 0
	}
	
	if res.Color == "red" {
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write([]byte(res.Result))
	return
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/",homeHandler).Methods("GET")
	r.HandleFunc("/",sendHandler).Methods("POST")
	r.HandleFunc("/verif",verifHandler).Methods("GET")
	r.HandleFunc("/verif",sendHandler).Methods("POST")
	r.HandleFunc("/api",apiHandler).Methods("POST")
	
	log.Println("Server Up and Running ...")
	err := http.ListenAndServe("0.0.0.0:8181",r)
	
	if err != nil {
		log.Fatal(err)
	}
}

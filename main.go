package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"

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
}

type ConfigFile struct{
	Content string
	Errors map[string]string
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
	var cfg ConfigFile
	cfg.Content = 
`scrape_configs:
	- job_name: blackbox
	  params:
		module:
		  - http_2xx
	  scrape_interval: 1m
	  scrape_timeout: 10s
	  metrics_path: /probe
	  scheme: http
  
	  static_configs:
		- targets:
			- http://host/metrics
	  ec2_sd_configs:
		- endpoint: http://host
		  port: 8080
		  region: us-east-1
		  refresh_interval: 1m
		  filters:
			- name: "tag:prometheus:tag"
			  values:
			  - xyz
  
	  gce_sd_configs:
		- project: example-project
		  zone: us-east1-a
		  port: 8181
		- project: example-project
		  zone: us-east1-b
		  port: 8181
  
	  relabel_configs:
		- source_labels: [__address__]
		  target_label: __param_target
		- source_labels: [__param_target]
		  target_label: instance
		- source_labels: [__param_target]
		  target_label: node_name`

		//cfg.content = r.PostFormValue("config")
		cfg.Validate()

		http.Redirect(w, r, "/verif", http.StatusSeeOther)
}

func verifHandler(w http.ResponseWriter, r *http.Request){
	createTemplate(w, "templates/verif.html", nil)
}




func main() {

	mux := pat.New()
	mux.Get("/",http.HandlerFunc(homeHandler))
	mux.Post("/",http.HandlerFunc(sendHandler))
	mux.Get("/verif",http.HandlerFunc(verifHandler))

	
	log.Println("Server Up and Running ...")
	err = http.ListenAndServe("0.0.0.0:8181",mux)
	
	if err != nil {
		log.Fatal(err)
	}
}

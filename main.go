package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
	"os/exec"

	"github.com/bmizerany/pat"
	//"github.com/prometheus/prometheus/cmd/promtool/"

	//pconfig "github.com/percona/promconfig"
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
	
	fmt.Println("[+] Checking ...")
	//CheckConfig("promconfig-main/testdata/test1.yml")

	cmd := exec.Command("promtool","check","config","testdata/test.yml")
   	out, err := cmd.CombinedOutput()
    	if err != nil {
        	log.Fatalf("cmd.Run() failed with %s\n", err)
   	}
   	fmt.Printf("Combined out:\n%s\n", string(out))
	fmt.Println("[-] Checking finished")

	log.Println("Server Up and Running ...")
	err = http.ListenAndServe("0.0.0.0:8181",mux)
	
	if err != nil {
		log.Fatal(err)
	}
}

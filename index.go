package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"

	"github.com/bmizerany/pat"
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

func (config *ConfigFile) Validate() bool{
	config.Errors = make(map[string]string)
	b := []byte(config.Content)

	fmt.Println("[+] Checking ...")

	fmt.Println(b)
	/*
	matches, err := filepath.Glob("testdata/*.yml")
	require.NoError(t, err)
	require.NotEmpty(t, matches)

	for _, yf := range matches {
		b, err := ioutil.ReadFile(yf) //=> returns bytes
		require.NoError(t, err) //=> Similar to err != nil

		var cfg Config
		err = yaml.Unmarshal(b, &cfg)
		require.NoError(t, err)
		actualB, err := json.MarshalIndent(cfg, "", "  ")
		require.NoError(t, err)
		actualB = append(actualB, '\n')

		jf := strings.TrimSuffix(yf, filepath.Ext(yf)) + ".json"

		if *goldenF {
			err = ioutil.WriteFile(jf, actualB, 0644)
			require.NoError(t, err)
		}

		expectedB, err := ioutil.ReadFile(jf)
		require.NoError(t, err)

		expectedS := string(expectedB)
		actualS := string(actualB)
		assert.Equal(t, expectedS, actualS)
	}
	*/
	fmt.Println("[-] Checking finished")

	return len(config.Errors) == 0
}


func main() {

	mux := pat.New()
	mux.Get("/",http.HandlerFunc(homeHandler))
	mux.Post("/",http.HandlerFunc(sendHandler))
	mux.Get("/verif",http.HandlerFunc(verifHandler))

	var cfg ConfigFile
	cfg.Content = "Content"

	cfg.Validate()

	log.Println("Server Up and Running ...")
	err := http.ListenAndServe("0.0.0.0:8181",mux)
	if err != nil {
		log.Fatal(err)
	}
}

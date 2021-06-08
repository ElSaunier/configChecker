package main 

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
	"bufio"
	"strings"

	//pconfig "github.com/percona/promconfig"
)


func (config *ConfigFile) Validate() bool{

	global := map[string]bool{"global" : false,"rule_files" : false,"scrape_configs" : false,"alerting" : false,"remote_write" : false,"remote_read" : false}
	config.Errors = make(map[string]string)
	str := config.Content

	fmt.Println("[+] Checking ...")

	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		curLine := scanner.Text()
		/*if curLine[0] != " "
			if global[curLine]
				return false
			else
				global[curLine] = ValidateGlobal()*/

		if (!(strings.Contains(curLine," ") || strings.Compare(curLine,"") == 0))
			return cfg.ValidateGlobal()		 
	}

	return false
}

func (config *ConfigFile)
package main

import (
	"os"
	"fmt"
	"log"
	"os/exec"
)

type ConfigFile struct{
	Content string
}

func (config ConfigFile) ValidatePromtool() (string, string, bool){
/*
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
*/


	fmt.Println("[+] Checking ...")

	file,errFile := os.Create("testdata/test.yml")
	if errFile != nil {
		log.Fatal("Cannot create File : %s\n",errFile)
	}
	_, errWrite := file.Write([]byte(config.Content))
	if errWrite != nil {
		log.Fatal("Cannot write in File : %s\n",errWrite)
	}

	cmd := exec.Command("promtool","check","config",file.Name())
   	out, err := cmd.CombinedOutput()
    if err != nil {
    	fmt.Println("cmd.Run() failed with %s\n", err)
   	} else {
		fmt.Printf("Combined out:\n%s\n", string(out))
	}
   	
	file.Close()
	fmt.Println("[-] Checking finished")

	return config.Content, string(out), err == nil
}

func (config ConfigFile) ValidateAlertManager() (string, string, bool){
	/*
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
	*/
	
	
		fmt.Println("[+] Checking ...")
	
		file,errFile := os.Create("testdata/test.yml")
		if errFile != nil {
			log.Fatal("Cannot create File : %s\n",errFile)
		}
		_, errWrite := file.Write([]byte(config.Content))
		if errWrite != nil {
			log.Fatal("Cannot write in File : %s\n",errWrite)
		}
	
		cmd := exec.Command("amtool","check-config",file.Name())
		   out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("cmd.Run() failed with %s\n", err)
		   } else {
			fmt.Printf("Combined out:\n%s\n", string(out))
		}
		   
		file.Close()
		fmt.Println("[-] Checking finished")
	
		return config.Content, string(out), err == nil
	}

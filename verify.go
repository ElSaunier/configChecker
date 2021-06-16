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

func (config ConfigFile) ValidatePromtool() (string, string, string){
	color := "green"
	file,errFile := os.OpenFile("/data/verif.yml",os.O_RDWR|os.O_TRUNC,0666)
	if errFile != nil {
		log.Fatal("Cannot open File : ",errFile)
	}
	_, errWrite := file.Write([]byte(config.Content))
	if errWrite != nil {
		log.Fatal("Cannot write in File : ",errWrite)
	}

	cmd := exec.Command("promtool","check","config",file.Name())
   	out, err := cmd.CombinedOutput()
    if err != nil {
    	fmt.Printf("cmd.Run() failed with %s\n", err)
		color = "red"
   	} 
	fmt.Printf("Combined out:\n%s\n", string(out))
	file.Close()

	return config.Content, string(out), color
}

func (config ConfigFile) ValidateAlertManager() (string, string, string){
		color := "green"
		file,errFile := os.OpenFile("/data/verif.yml",os.O_RDWR|os.O_TRUNC,0666)
		if errFile != nil {
			log.Fatal("Cannot open File : ",errFile)
		}
		_, errWrite := file.Write([]byte(config.Content))
		if errWrite != nil {
			log.Fatal("Cannot write in File :",errWrite)
		}
	
		cmd := exec.Command("amtool","check-config",file.Name())
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("cmd.Run() failed with %s\n", err)
			color = "red"
		}
		fmt.Printf("Combined out:\n%s\n", string(out))		   
		file.Close()

		return config.Content, string(out), color
	}

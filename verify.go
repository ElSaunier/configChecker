package main

import (
	"os/exec"
)

type ConfigFile struct{
	Content string
}

func (config *ConfigFile) Validate() bool{

	fmt.Println("[+] Checking ...")

	file,errFile := os.Create("testdata/test.yml")
	if errFile != nil {
		log.Fatal("Cannot create File : %s\n",errFile)
	}
	_, errWrite := File.Write([]byte(config.Content))
	if errWrite != nil {
		log.Fatal("Cannot write in File : %s\n",errWrite)
	}

	cmd := exec.Command("promtool","check","config",file)
   	out, err := cmd.CombinedOutput()
    if err != nil {
    	fmt.Println("cmd.Run() failed with %s\n", err)
   	} else {
		fmt.Printf("Combined out:\n%s\n", string(out))
	}
   	
	file.close()
	fmt.Println("[-] Checking finished")

	return err == nil
}

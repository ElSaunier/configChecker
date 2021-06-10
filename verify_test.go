package main

import (
	"testing"
	"os"
	"os/exec"
)

func TestValidatePromtool(t *testing.T) {
	file,err := os.Open("testdata/test.yml")
	if err != nil {
		t.Errorf("Cannot open File : %s\n",err)
	}
	
	cmd := exec.Command("promtool","check","config",file.Name())
   	out, errc := cmd.CombinedOutput()
    if errc != nil {
    	t.Errorf("cmd.Run() failed with %s\n %s\n", errc, out)
   	}
}

func TestValidateAlertManager(t *testing.T) {
	file,err := os.Open("testdata/testam.yml")
	if err != nil {
		t.Errorf("Cannot open File : %s\n",err)
	}

	cmd := exec.Command("amtool","check-config",file.Name())
   	out, errc := cmd.CombinedOutput()
    if errc != nil {
    	t.Errorf("cmd.Run() failed with %s\n %s\n", errc, out)
   	}    
}


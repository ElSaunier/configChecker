package main

import (
	"fmt"

	pconfig "github.com/percona/promconfig"
)

type ConfigFile struct{
	Content string
	Errors map[string]string
}

func (config *ConfigFile) Validate() bool{
	config.Errors = make(map[string]string)

	fmt.Println("Checking ...")
	pconfig.TestGoldenData()

	return len(config.Errors) == 0
}

package main

import (
	"github.com/percona/promconfig"
)

type ConfigFile {
	Content string
	Errors map[string]string
}

func (config *ConfigFile) Validate() bool{
	config.Errors = make(map[string]string)


	return len(config.Errors) == 0
}
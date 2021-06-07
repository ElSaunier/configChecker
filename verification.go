import (
	pconfig "promconfig-main/promconfig"
)

type ConfigFile {
	Content string
	Errors map[string]string
}

func (config *ConfigFile) Validate() bool{
	config.Errors = make(map[string]string)


	promconfig.TestGoldenData()

	return len(config.Errors) == 0
}

package configChecker

import (
	"fmt"
	//pconfig "github.com/percona/promconfig"
)

type ConfigFile struct{
	Content string
	Errors map[string]string
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

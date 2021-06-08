// promconfig
// Copyright 2020 Percona LLC
//
// Based on Prometheus systems and service monitoring server.
// Copyright 2015 The Prometheus Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package promconfig // import "github.com/percona/promconfig"

import (
	"regexp"
)

// Config is the top-level configuration for Prometheus's config files.
type Config struct {
	GlobalConfig       GlobalConfig         `yaml:"global"`
	AlertingConfig     AlertingConfig       `yaml:"alerting,omitempty"`
	RuleFiles          []string             `yaml:"rule_files,omitempty"`
	ScrapeConfigs      []*ScrapeConfig      `yaml:"scrape_configs,omitempty"`
	RemoteWriteConfigs []*RemoteWriteConfig `yaml:"remote_write,omitempty"`
	RemoteReadConfigs  []*RemoteReadConfig  `yaml:"remote_read,omitempty"`
}

func (Config cfg) VerifCfg (bool) {

	verifRules bool
	// Pas de caractères spécieux dans les noms de fichiers
	r , _ := regexp.Compile("[a-zA-Z0-9-_\/]+")


	for _,s := range cfg.RuleFiles {
		if (!r.MatchString(s))
			verifRules = false
	}


	return cfg.VerifGCfg() &&
			cfg.AlertingConfig.VerifACfg() &&
			verifRules &&
			cfg.ScrapeConfigs.VerifSCfg() && 
			cfg.RemoteWriteConfigs.VerifRWCfg() &&
			cfg.RemoteReadConfigs.VerifRRCfg()
}
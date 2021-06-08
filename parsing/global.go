package parser

type Global struct{
	
	ScrapeInterval Duration `yaml:"scrape_interval,omitempty"`
	ScrapeTimeout Duration `yaml:"scrape_timeout,omitempty"`
	EvaluationInterval Duration `yaml:"evaluation_interval,omitempty"`
	QueryLogFile string `yaml:"query_log_file,omitempty"`
	ExternalLabels map[string]string `yaml:"external_labels,omitempty"`
}
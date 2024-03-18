package config

type InternalConfig struct {
	AppName                            string
	Port                               string
	DefaultTimeoutHTTPExternalRequests int
}

type Config struct {
	Internal *InternalConfig
}

func Get() *Config {
	return &Config{
		Internal: &InternalConfig{
			AppName:                            "AMSVault",
			Port:                               "8080",
			DefaultTimeoutHTTPExternalRequests: 10,
		},
	}
}

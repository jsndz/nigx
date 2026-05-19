package config

type Config struct {
	Route   string
	Proxies []string
}

func NewConfig(route string, proxies []string) *Config {
	return &Config{
		Route:   route,
		Proxies: proxies,
	}
}

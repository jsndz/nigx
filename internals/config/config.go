package config

type Config struct {
	Route string
	Proxy string
}

func NewConfig(route string, proxy string) *Config {
	return &Config{
		Route: route,
		Proxy: proxy,
	}
}

package config

import "flag"

// Config holds application configuration
type Config struct {
	Port int
	Host string
}

// Load loads configuration from command line flags
func Load() *Config {
	cfg := &Config{}
	
	flag.IntVar(&cfg.Port, "port", 8080, "Port number to run the server on")
	flag.StringVar(&cfg.Host, "host", "localhost", "Host to bind the server to")
	flag.Parse()
	
	return cfg
}

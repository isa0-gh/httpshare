package config

import "flag"

// Config holds application configuration
type Config struct {
	Port      int
	Host      string
	Directory string
	LogFile   string
}

// Load loads configuration from command line flags
func Load() *Config {
	cfg := &Config{}

	flag.IntVar(&cfg.Port, "port", 8080, "Port number to run the server on")
	flag.StringVar(&cfg.Host, "host", "0.0.0.0", "Host to bind the server to")
	flag.StringVar(&cfg.Directory, "directory", ".", "Path to directory to serve files")
	flag.StringVar(&cfg.LogFile, "log", "", "Write logs to a file")
	flag.Parse()

	return cfg
}

var Cfg *Config = Load()

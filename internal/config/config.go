package config

import "os"

const (
	httpSrvAddrEnv = "HTTP_SERVER_ADDR" // host:port
	logLevelEnv    = "LOG_LEVEL"
)

type Config struct {
	LogLevel       string
	HTTPServerAddr string
}

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	cfg = &Config{}

	if ll, ok := os.LookupEnv(logLevelEnv); ok {
		cfg.LogLevel = ll
	}

	if addr, ok := os.LookupEnv(httpSrvAddrEnv); ok {
		cfg.HTTPServerAddr = addr
	}

	return *cfg
}

package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Config holds the application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logging  LoggingConfig
	CORS     CORSConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port         string `mapstructure:"port"`
	Environment  string `mapstructure:"environment"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	IdleTimeout  int    `mapstructure:"idle_timeout"`
	JWTSecret    string `mapstructure:"jwt_secret"`
}

// DatabaseConfig holds database-specific configuration
type DatabaseConfig struct {
	URL            string `default:"postgres://postgres:postgres@localhost:5432/apm"`
	MaxConns       int    `envconfig:"MAX_CONNS" default:"10"`
	MinConns       int    `envconfig:"MIN_CONNS" default:"2"`
	MaxConnLife    int    `envconfig:"MAX_CONN_LIFE" default:"3600"` // 1 hour in seconds
	MaxConnIdle    int    `envconfig:"MAX_CONN_IDLE" default:"300"`  // 5 minutes in seconds
	ConnectTimeout int    `envconfig:"CONNECT_TIMEOUT" default:"10"` // 10 seconds
}

// LoggingConfig holds logging-specific configuration
type LoggingConfig struct {
	Level  string `default:"info"`
	Format string `default:"json"`
}

// CORSConfig holds CORS-specific configuration
type CORSConfig struct {
	AllowedOrigins []string `envconfig:"ALLOWED_ORIGINS" default:"*"`
	AllowedMethods []string `envconfig:"ALLOWED_METHODS" default:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders []string `envconfig:"ALLOWED_HEADERS" default:"Content-Type,Authorization"`
}

// Load loads the application configuration from environment variables
// using the envconfig library
func Load() (Config, error) {
	var cfg Config
	err := envconfig.Process("APM", &cfg)
	return cfg, err
}

// LoadOrFatal loads configuration or panics if there's an error
func LoadOrFatal() Config {
	cfg, err := Load()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
	return cfg
}

// PrintUsage prints the available environment variables and their defaults
func PrintUsage() {
	var cfg Config
	if err := envconfig.Usage("APM", &cfg); err != nil {
		panic("Failed to print usage: " + err.Error())
	}
}

// PrintUsageToWriter prints the available environment variables and their defaults to the specified writer
func PrintUsageToWriter(w *os.File) {
	var cfg Config
	if err := envconfig.Usagef("APM", &cfg, w, envconfig.DefaultTableFormat); err != nil {
		panic("Failed to print usage: " + err.Error())
	}
}

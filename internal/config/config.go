package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string        `yaml:"env" env-required:"true"`
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
	Grpc     GRPCConfig    `yaml:"grpc"`
	Storage  StorageConfig `yaml:"database"`
}

type StorageConfig struct {
	Type     string `yaml:"type"`
	Path     string `yaml:"path"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load() *Config {
	path := fetchConfigPath()
	if path == "" {
		log.Fatal("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("Config file not found: ", path, err.Error())
	}

	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatal("Failed to read config file", err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var result string

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// --config="path/to/config.yaml"
	flag.StringVar(&result, "config", "", "path to config file")
	flag.Parse()

	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}

	return result
}

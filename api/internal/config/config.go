package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ListenAddr string
	Database   DatabaseConfig
	Consul     ConsulConfig
	Vault      VaultConfig
	RabbitMQ   RabbitMQConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

type ConsulConfig struct {
	Address string
	Token   string
}

type VaultConfig struct {
	Address string
	Token   string
}

type RabbitMQConfig struct {
	URL string
}

func Load() (*Config, error) {
	viper.SetDefault("listen_addr", ":8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.sslmode", "disable")
	
	viper.AutomaticEnv()
	
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	
	return &config, nil
} 
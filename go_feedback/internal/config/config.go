package config

import (
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSchema   string `mapstructure:"DB_SCHEMA"`
	DBPort     string `mapstructure:"DB_PORT"`

	ServerPort      string        `mapstructure:"SERVER_PORT"`
	RedisHost       string        `mapstructure:"REDIS_HOST"`
	RedisPort       string        `mapstructure:"REDIS_PORT"`
	RedisPassword   string        `mapstructure:"REDIS_PASSWORD"`
	RedisExpiration time.Duration // Stored as time.Duration after parsing

	EmailHost     string `mapstructure:"EMAIL_HOST"`
	EmailPort     int    `mapstructure:"EMAIL_PORT"`
	EmailUsername string `mapstructure:"EMAIL_USERNAME"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
	EmailFrom     string `mapstructure:"EMAIL_FROM"`
	SwaggerHost   string `mapstructure:"SWAGGER_HOST"`

	BasicAuthUsername string `mapstructure:"BASIC_AUTH_USERNAME"`
	BasicAuthPassword string `mapstructure:"BASIC_AUTH_PASSWORD"`
	BearerSignerKey   string `mapstructure:"BEARER_TOKEN_SIGNER_KEY"`

	AccessExpiryDays  int `mapstructure:"ACCESS_EXPIRY_DAYS"`
	RefreshExpiryDays int `mapstructure:"REFRESH_EXPIRY_DAYS"`

	DisableAuth bool `mapstructure:"DISABLE_AUTH"`

	STSRefreshMins time.Duration
}

func LoadConfig() (Config, error) {
	return LoadGivenConfig(".", ".env")
}

func LoadGivenConfig(configPath, configName string) (Config, error) {
	var config Config

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetDefault("DB_SCHEMA", "user_management")
	viper.SetDefault("DISABLE_AUTH", false)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	config.RedisExpiration = parseDuration("Redis expiration seconds", "REDIS_EXPIRATION", time.Second, -1)

	return config, nil
}

func MustLoadConfig() Config {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return config
}

// ParseDuration - parse string into Duration, def will contain the default value, -1 means fail on error
func parseDuration(logTitle, envVar string, duration time.Duration, def int) time.Duration {
	expirationStr := viper.GetString(envVar)
	expiration, err := strconv.Atoi(expirationStr)
	if err != nil {
		if def == -1 {
			log.Fatalf("Invalid expiration value in %s: %v", envVar, err)
		} else {
			expiration = def
		}
	}
	return time.Duration(expiration) * duration
}

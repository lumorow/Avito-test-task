package config

import "github.com/spf13/viper"

type (
	Config struct {
		PostgresDB `mapstructure:"PSQL_CONF"`
		Logger     `mapstructure:"LOGGER"`
		HTTP       `mapstructure:"HTTP"`
	}

	PostgresDB struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Dbname   string `mapstructure:"DB_NAME"`
		SSLMode  string `mapstructure:"SSLMODE"`
	}

	Logger struct {
		Level string `mapstructure:"LEVEL"`
	}

	HTTP struct {
		Port string `mapstructure:"HTTP_PORT"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	return cfg, nil
}

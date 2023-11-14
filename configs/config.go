package configs

import (
	"os"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUsername        string `mapstructure:"DB_USERNAME"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfigLocal() *conf {
	var cfg *conf

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg
}

func LoadConfigDocker() *conf {
	return &conf{
		DBDriver:          os.Getenv("DB_DRIVER"),
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            os.Getenv("DB_PORT"),
		DBUsername:        os.Getenv("DB_USERNAME"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		WebServerPort:     os.Getenv("WEB_SERVER_PORT"),
		GRPCServerPort:    os.Getenv("GRPC_SERVER_PORT"),
		GraphQLServerPort: os.Getenv("GRAPHQL_SERVER_PORT"),
	}
}

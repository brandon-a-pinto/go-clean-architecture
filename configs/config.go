package configs

import (
	"os"
	"strconv"
)

type Conf struct {
	DBDriver          string
	DBHost            string
	DBPort            string
	DBUsername        string
	DBPassword        string
	DBName            string
	WebServerPort     string
	GRPCServerPort    string
	GraphQLServerPort string
	JWTSecret         string
	JWTExpiresIn      int
}

func LoadConfig() *Conf {
	exp, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	return &Conf{
		DBDriver:          os.Getenv("DB_DRIVER"),
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            os.Getenv("DB_PORT"),
		DBUsername:        os.Getenv("DB_USERNAME"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		WebServerPort:     os.Getenv("WEB_SERVER_PORT"),
		GRPCServerPort:    os.Getenv("GRPC_SERVER_PORT"),
		GraphQLServerPort: os.Getenv("GRAPHQL_SERVER_PORT"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTExpiresIn:      exp,
	}
}

package config

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"strconv"
)

func Load(files ...string) {
	err := gotenv.Load(files...)

	if err != nil {
		log.Println(err)
	}
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetEnv() string {
	return os.Getenv("ENVIRONMENT")
}

func GetDBDsn() string {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	SSLMode := os.Getenv("DB_SSL")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		username,
		password,
		name,
		port,
		SSLMode,
	)

	return dsn
}

func GetAuthServiceGRPCurl() string {
	url := fmt.Sprintf("%s:%s", os.Getenv("AUTH_SERVICE_GRPC_HOST"), os.Getenv("AUTH_SERVICE_GRPC_PORT"))

	return url
}

func GetEmail() string {
	return os.Getenv("EMAIL")
}

func GetEmailPass() string {
	return os.Getenv("EMAIL_PASS")
}

func GetSMTPHost() string {
	return os.Getenv("SMTP_HOST")
}

func GetSMTPPort() string {
	return os.Getenv("SMTP_PORT")
}

type AMQPConfigs struct {
	User, Host, Pass string
	Port             int
}

func toInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}

func GetRabbitmq() AMQPConfigs {
	return AMQPConfigs{
		User: os.Getenv("AMQP_USER"),
		Host: os.Getenv("AMQP_HOST"),
		Pass: os.Getenv("AMQP_PASS"),
		Port: toInt(os.Getenv("AMQP_PORT")),
	}
}

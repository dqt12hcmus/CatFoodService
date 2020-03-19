package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	TestEnv = "test"
	DevEnt  = "dev"
)

var Env string
var AMQP = struct {
	Host                   string
	Port                   string
	UserName               string
	Password               string
	ConnectionString       string
	QueueName              string
	NumOfConnectionRetries int
	RetryDelay             time.Duration
}{
	QueueName:              "order",
	NumOfConnectionRetries: 20,
	RetryDelay:             time.Second,
}
var DB = struct {
	Host                   string
	Port                   string
	UserName               string
	Password               string
	DBName                 string
	Collection             string
	ConnectionString       string
	Driver                 string
	NumOfConnectionRetries int
}{
	NumOfConnectionRetries: 20,
}

func getOptionalEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.WithFields(log.Fields{"err": err}).Info("Skip environment file")
	}

	AMQP.Host = getOptionalEnv("RABBITMQ_HOST", "rabbitmq")
	AMQP.Port = getOptionalEnv("RABBITMQ_PORT", "5672")
	AMQP.UserName = getOptionalEnv("RABBITMQ_USER", "test")
	AMQP.Password = getOptionalEnv("RABBITMQ_PASS", "test")
	AMQP.ConnectionString = fmt.Sprintf("amqp://%v:%v@%v:%v/", AMQP.UserName, AMQP.Password, AMQP.Host, AMQP.Port)

	Env = getOptionalEnv("ENV", DevEnt)
	DB.Host = getOptionalEnv("MONGO_HOST", "localhost")
	DB.Port = getOptionalEnv("MONGO_PORT", "5432")
	DB.UserName = getOptionalEnv("MONGO_USERNAME", "admin")
	DB.Password = getOptionalEnv("MONGO_PASSWORD", "admin")
	DB.DBName = getOptionalEnv("ORDER_DATABASE", "order-service")
	DB.Collection = getOptionalEnv("ORDER_COLLECTION", "orders")

	DB.ConnectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", DB.UserName, DB.Password, DB.Host, DB.Port, DB.DBName)
}

package tool

import (
	"context"
	"errors"
	"order-service/config"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConnector interface {
	Connect() (*mongo.Database, error)
	NumTries() int
	BreakTime() time.Duration
}

type dbConnector struct{}

func (*dbConnector) Connect() (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DB.ConnectionString))
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return
	}
	db = client.Database(config.DB.DBName)

	return
}

func (*dbConnector) BreakTime() time.Duration {
	return time.Second
}

func (*dbConnector) NumTries() int {
	return config.DB.NumOfConnectionRetries
}

func NewDBConnector() DBConnector {
	return &dbConnector{}
}

func ConnectDB(connector DBConnector) (db *mongo.Database, err error) {
	errorCount := 0
	for {
		db, err = connector.Connect()
		if err == nil {
			return
		}

		errorCount++
		if errorCount < connector.NumTries() {
			log.WithFields(log.Fields{
				"err":      err,
				"numTries": errorCount,
			}).Info("Error connecting to DB. Retrying...")
			time.Sleep(connector.BreakTime())
		} else {
			return nil, errors.New("cannot connect to DB")
		}
	}
}

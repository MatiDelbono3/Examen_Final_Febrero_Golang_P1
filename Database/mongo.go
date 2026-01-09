package database

import (
	"context"
	"fmt"
	"log"

	event "go.mongodb.org/mongo-driver/event"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	Connect() error
	Disconnect() error
	GetClient() *mongo.Client
}
type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB crea una nueva instancia de MongoDB
func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (db *MongoDB) Connect() error {
	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			log.Printf("MongoDB query started: %s %v", evt.CommandName, evt.Command)
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			log.Printf("MongoDB query succeeded: %s took %dms", evt.CommandName, evt.DurationNanos/1e6)
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
			log.Printf("MongoDB query failed: %s, error: %v", evt.CommandName, evt.Failure)
		},
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetMonitor(monitor)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	db.client = client
	fmt.Printf("Client Connect: %v\n", db.client)

	return nil
}

func (db *MongoDB) Disconnect() error {
	fmt.Printf("Client Disconnect: %v\n", db.client)

	if db.client == nil {
		return nil
	}

	return db.client.Disconnect(context.Background())
}

func (db *MongoDB) GetClient() *mongo.Client {
	return db.client
}

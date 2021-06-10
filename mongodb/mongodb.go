package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	dsn := fmt.Sprintf("mongodb://root:toor@localhost:27017")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	//defer func() {
	//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//	defer cancel()
	//
	//	err := client.Disconnect(ctx)
	//	if err != nil {
	//		log.Printf("Disconnect: %s\n", err)
	//	}
	//}()
	return client, client.Ping(context.Background(), readpref.Primary())
}

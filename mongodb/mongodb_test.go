package mongodb

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestNewMongoDB(t *testing.T) {
	client, err := NewMongoDB()
	if err != nil {
		t.Fatalf("NewMongoDB: %s", err)
	}

	c := client.Database("test_mongodb").Collection("test_collection")
	_, err = c.InsertOne(context.Background(), bson.D{{Key: "hello", Value: "world"}})
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	type A struct {
		Name string
		Age  int
	}
	a := &A{
		Name: "hello",
		Age:  12,
	}
	//data, err := bson.Marshal(a)
	//if err != nil {
	//	t.Fatalf("%s\n", err)
	//}
	_, err = c.InsertOne(context.Background(), a)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	//log.Printf("data: %s\n", data)
}

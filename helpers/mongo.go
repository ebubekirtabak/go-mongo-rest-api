package helpers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDatabase() *mongo.Client {

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB_URI"))

	ctx, _ := context.WithTimeout(context.Background(), 35 * time.Second)
	var err error
	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Ping Exception ", err)
		return nil
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func InsertOne(collectionName string, query bson.M) bool {
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, query)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

func IsExistsObject(collectionName string, query bson.M) bool {
	collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
	err := collection.FindOne(context.TODO(), query)
	if err.Err() != nil {
		return false
	}

	return true
}

func FindOne(collectionName string, query bson.M) (*mongo.SingleResult, error) {
	collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
	singleResult := collection.FindOne(context.TODO(), query)
	return singleResult, singleResult.Err()
}

func FindAll(collectionName string, query bson.M) (*mongo.Cursor, error) {
	collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
	singleResult, err := collection.Find(context.TODO(), query)
	return singleResult, err
}

func DeleteOne(collectionName string, query bson.M) (int64, bool) {
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
	result, err := collection.DeleteOne(ctx, query)
	if err != nil {
		return 0, false
	}

	return result.DeletedCount, true
}

func UpdateOne(collectionName string, query bson.M, data bson.M) (*mongo.UpdateResult, bool) {
	collection := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
	result, err := collection.UpdateOne(
		context.TODO(),
		query,
		data,
	)
	if err != nil {
		fmt.Println("UpdateOne -> ", err)
		return nil, false
	}

	return result, true
}
package database

//
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://morphatend:morph234@cluster0.ilg5d1g.mongodb.net/"))

	if err != nil {
		log.Fatal(err)
	}

	defer cancel()

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("Failed to connect to mongodb")
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to mongodb")

	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("VShop").Collection(CollectionName)
	return collection

}

func ProductData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("VShop").Collection(CollectionName)
	return productcollection
}

func ShopData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var shopcollection *mongo.Collection = client.Database("VShop").Collection(CollectionName)
	return shopcollection
}

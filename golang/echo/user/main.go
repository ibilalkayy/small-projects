package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Phone     int                `json:"phone" bson:"phone"`
	Password  string             `json:"password" bson:"password"`
	City      string             `json:"city" bson:"city"`
	Address   string             `json:"address" bson:"address"`
}

func Load() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load the .env file", err)
	}

	mongo_uri := os.Getenv("MONGO_URI")
	return mongo_uri
}

func Connect(MONGO_URI string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, errors.New("failed to connect to client")
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, errors.New("the connection is failed")
	} else {
		fmt.Println("Connected to MongoDB database")
	}

	return client, nil
}

func CreateCollection(u User, client *mongo.Client) error {
	collection := client.Database("user").Collection("app")
	fmt.Println("Collection is created")

	doc := bson.D{
		{Key: "firstname", Value: u.FirstName},
		{Key: "lastname", Value: u.LastName},
		{Key: "email", Value: u.Email},
		{Key: "phone", Value: u.Phone},
		{Key: "password", Value: u.Password},
		{Key: "city", Value: u.City},
		{Key: "address", Value: u.Address},
	}
	_, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatal("Failed to insert document: ", err)
	}

	fmt.Println("Document inserted into the collection")
	return nil
}

func main() {
	mongo_uri := Load()
	client, err := Connect(mongo_uri)
	if err != nil {
		log.Fatal("failed to connect to the database", err)
	}

	e := echo.New()

	e.POST("/create", func(ctx echo.Context) error {
		var newUser User
		if err := ctx.Bind(&newUser); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		if err := CreateCollection(newUser, client); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return ctx.JSON(http.StatusCreated, newUser)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

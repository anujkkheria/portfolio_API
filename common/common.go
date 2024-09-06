package common

import (
	"context"

	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/options"
)

var db *mongo.Database

func GetDBCollection(col string)*mongo.Collection{
	return	db.Collection(col)
}

func InitDB() error{
	uri:=os.Getenv("MONGO_URI")
	if uri == " "{
		return errors.New("db uri missing")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	fmt.Println(client)
	if err!= nil{
		return err
	}

	db = client.Database("Myinfo")

	return nil
}

func CloseDB() error{

	return db.Client().Disconnect(context.Background())

}
package main

import (
	"context"
	"encoding/csv"
	"eventials/database"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Inserting companies from csv file
func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(database.DatabaseURIEnvVariable("URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	YawoenDatabase := client.Database("Yawoen")
	companiesCollection := YawoenDatabase.Collection("companies")

	file, err := os.Open("q1_catalog.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println("Inserting companies into database from q1_catalog")

	for i := 1; i < len(records); i++ {

		zipSize := len(strings.Split(records[i][0], ";")[1])

		if zipSize == 5 {
			_, err := companiesCollection.InsertOne(context.TODO(), bson.D{
				{Key: "name", Value: strings.ToUpper(strings.Split(records[i][0], ";")[0])},
				{Key: "zip", Value: strings.Split(records[i][0], ";")[1]},
			})

			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("Companies added to database!")

	if err != nil {
		log.Fatal(err)
	}
}

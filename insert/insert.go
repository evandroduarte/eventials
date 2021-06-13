package insert

import(
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"os"
	"fmt"
	"strings"
	"encoding/csv"
	"context"
	"time"
)

//Inserting companies from csv file
func InsertCompanies() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://evandro:YWQdeZjCTlP8kka6@eventials.q0i7q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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

	for i := 1; i < len(records); i++ {
		companyResult, err := companiesCollection.InsertOne(ctx, bson.D{
			{Key: "name", Value: strings.ToUpper(strings.Split(records[i][0], ";")[0])},
			{Key: "zip", Value: strings.Split(records[i][0], ";")[1]},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(companyResult)
	}

	if err != nil {
		log.Fatal(err)
	}
}
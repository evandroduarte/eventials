package update

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"
	"strings"
	"time"
)

func UpdateCompanies(records [][]string) {
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

	for i := 1; i < len(records); i++ {
		filter := bson.M{"name": strings.ToUpper(strings.Split(records[i][0], ";")[0]),
			"zip": strings.Split(records[i][0], ";")[1],
		}

		update := bson.M{
			"$set": bson.M{"website": strings.Split(records[i][0], ";")[2]},
		}

		companiesCollection.FindOneAndUpdate(ctx, filter, update)
	}

}
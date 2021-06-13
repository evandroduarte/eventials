package get

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Company struct {
	Id string
	Name string
	Zip  string
	Website string
}

//API request for getting company by name and zip code
func GetCompany(w http.ResponseWriter, r *http.Request) {
	var company Company

	err := json.NewDecoder(r.Body).Decode(&company)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	regex := `(?i).*` + company.Name

	filterCursor, err := companiesCollection.Find(ctx, bson.M{
		"$and": []bson.M{
			{"name": bson.M{"$regex": regex}},
			{"zip": company.Zip},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	var companiesFiltered []bson.M
	if err = filterCursor.All(ctx, &companiesFiltered); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(companiesFiltered)

}
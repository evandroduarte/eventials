package main

import (
	"eventials/get"
	"eventials/insert"
	"eventials/upload_csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	//By running main.go the companies are automatically added to de db
	insert.InsertCompanies()
	fmt.Println("Server listening on port 8080")
	setRoutes()
}

func DatabaseURIEnvVariable(URI string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(URI)
}

//Routes function
func setRoutes() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/upload", upload_csv.UploadCSV)
	router.HandleFunc("/company", get.GetCompany)
	log.Fatal(http.ListenAndServe(":8080", router))
}
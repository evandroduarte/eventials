package main

import (
	"eventials/get"
	"eventials/insert"
	"eventials/upload_csv"
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	//By running main.go the companies are automatically added to de db
	insert.InsertCompanies()
	fmt.Println("Server listening on port 8080")
	setRoutes()
}

//Routes function
func setRoutes() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/upload", upload_csv.UploadCSV)
	router.HandleFunc("/company", get.GetCompany)
	log.Fatal(http.ListenAndServe(":8080", router))
}
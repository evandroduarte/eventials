package main

import (
	"log"
	"net/http"
	"eventials/get"
	"eventials/upload_csv"
	"eventials/insert"

	"github.com/gorilla/mux"
)

func main() {
	insert.InsertCompanies()
	setRoutes()

}

//Routes function
func setRoutes() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/upload", upload_csv.UploadCSV)
	router.HandleFunc("/company", get.GetCompany)
	log.Fatal(http.ListenAndServe(":8080", router))
}
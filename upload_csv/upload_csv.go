package upload_csv

import(
	"eventials/update_company"
	"net/http"
	"fmt"
	"encoding/csv"
	"encoding/json"
)

//API POST request for CSV upload
func UploadCSV(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	if handler.Filename[len(handler.Filename)-3:] == "csv" && records[0][0] == "name;addresszip;website" {
		update.UpdateCompanies(records)
		json.NewEncoder(w).Encode("Documents updated")
		
	} else if records[0][0] != "name;addresszip;website" {
		fmt.Println("Wrong parameters file")
	} else {
		fmt.Println("Incorrect type of file!")
	}
}
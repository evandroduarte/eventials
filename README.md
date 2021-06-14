# Eventials data integration teste - Yawoen Company 

This is an aplication to integrate data from csv files into a Mongo database.

## Installation

Create a M0 Cluster on Mongo DB and retrieve the URI for connection, inside the root folder create a .env file with the following structure:

URI=<your_connection_URI>

## Usage

After the creation of .env file with connection URI to de Mongo database cluster use the command:

```Bash
go run main.go
```

OR

```Bash
go build main.go
```

⚠️ Running the aplication multiple times without clearing the database will result in multiple repeated entries

## API

The API will be accessible through port *8080*

The routes listed in the API are:

POST (/upload) - Recieve a multipart with key "myFile" and if the file is a well formated csv (name;address zip; website) start the data integration
between the companies in the csv and the companies already added to the database

GET (/company) - Recieve a json with the keys "name" and "zip" and their values for the company, returns a json containing all the information of the company in  
following structure:
```Bash
[
  {
    "_id": "60c75cfe56ebd3aa68b3f196",
    "name": "TOLA SALES GROUP",
    "website": "http://repsources.com",
    "zip": "78229"
  }
]
```

## Tests

Tests were made using Insomnia API Client


## License
[MIT](https://choosealicense.com/licenses/mit/)

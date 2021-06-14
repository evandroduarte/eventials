# Eventials data integration challenge - Yawoen Company 

This application was developed for the eventials data integration challenge, developed with Go v1.16 and Mongo DB

## Installation

Create a M0 Cluster on Mongo DB and retrieve the URI for connection, inside the root folder create a .env file with the following structure:

URI=<your_connection_URI>

## Usage

After the creation of .env file with connection URI to de Mongo database cluster use the command:

```Bash
make
```
Make will run the insert of the contents of q1_catalog.csv into the database and start the API

⚠️ Running the "make" command multiple times without clearing the database will result in multiple repeated entries

## API

To start the API without inserting the components of the q1_catalog.csv file to the database use the command:

```Bash
make start
```

The API will be accessible through port *8080*

The routes listed in the API are:

POST (/upload) - Recieve a multipart with key "myFile" and if the file is a well formated csv (name;address zip; website) start the data integration
between the companies in the csv and the companies already added to the database

GET (/company) - Recieve a json with the keys "name" and "zip" and their values for the company, returns a json containing all the information of the company 

response example:
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

Upload test:

![api_update](https://user-images.githubusercontent.com/39135867/121963895-1becbc80-cd41-11eb-888b-6c0c64c9a5b1.png)

Company get test:

![api_get](https://user-images.githubusercontent.com/39135867/121963890-1abb8f80-cd41-11eb-8e58-52baf1970238.png)


## License
[MIT](https://choosealicense.com/licenses/mit/)

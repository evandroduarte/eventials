#Challenge Makefile

all: setup start

start:
#TODO: commands necessary to start the API
	go run main.go

check:
#TODO: include command to test the code and show the results

setup:
#Adding the companies from the csv file
	go run insert/main.go
package main

import (
	"log"

	"github.com/iamseki/csv-to-db/app/mongousecase"
	"github.com/iamseki/csv-to-db/infra/csv"
	"github.com/iamseki/csv-to-db/infra/db/mongodb"
)

func main() {
	parser := csv.NewCDIFromCSV()
	repository := mongodb.NewInsertCDIsMongoRepository()
	usecase := mongousecase.NewCDIsToMongoDB(repository, parser)
	err := usecase.ConvertCdiCsvToMongodb()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("CDIs CSV converted into mongo database successfully")
}

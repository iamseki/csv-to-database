package main

import (
	"flag"
	"log"
	"os"

	"github.com/iamseki/csv-to-db/app/mongousecase"
	"github.com/iamseki/csv-to-db/infra/csv"
	"github.com/iamseki/csv-to-db/infra/db/mongodb"
)

func main() {
	var filename string
	dir, _ := os.Getwd()
	scriptsDir := dir + "/scripts/"
	flag.StringVar(&filename, "filename", scriptsDir+"CDI_Prices.csv", "csv filename to convert, e.g --filename=CDI_Prices.csv")
	flag.Parse()

	parser := csv.NewCDIFromCSV(filename)
	repository := mongodb.NewInsertCDIsMongoRepository()
	usecase := mongousecase.NewCDIsToMongoDB(repository, parser)

	err := usecase.ConvertCdiCsvToMongodb()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("CDIs CSV converted into mongo database successfully")
}
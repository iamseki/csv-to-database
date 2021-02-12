package main

import (
	"log"
)

func main() {
	options := parseProgramFlags()

	cdiUsecase := newCDIToMongoDBUseCase(options)

	err := cdiUsecase.ConvertCdiCsvToMongodb()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("CDIs CSV converted into mongo database successfully")
}

package main

import (
	"github.com/iamseki/csv-to-db/app/mongousecase"
	"github.com/iamseki/csv-to-db/domain"
	"github.com/iamseki/csv-to-db/infra/csv"
	"github.com/iamseki/csv-to-db/infra/db/mongodb"
)

func newCDIToMongoDBUseCase(options flags) domain.CSVToDatabaseWriter {
	parser := csv.NewCDIFromCSV(options.filename)
	repository := mongodb.NewInsertCDIsMongoRepository()
	return mongousecase.NewCDIsToMongoDB(repository, parser)
}

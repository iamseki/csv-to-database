package mongousecase

import (
	"github.com/iamseki/csv-to-db/contracts"
)

type CDIsToMongoDB struct {
	Repository contracts.InsertCDIsRepository
	Reader     contracts.CDIReader
}

func NewCDIsToMongoDB(repository contracts.InsertCDIsRepository, reader contracts.CDIReader) *CDIsToMongoDB {
	return &CDIsToMongoDB{
		Repository: repository,
		Reader:     reader,
	}
}

func (cdiToMongo *CDIsToMongoDB) ConvertCdiCsvToMongodb() error {
	cdis := cdiToMongo.Reader.ConvertCDIFromCSV()
	err := cdiToMongo.Repository.InsertMany(cdis)
	if err != nil {
		return err
	}
	return nil
}

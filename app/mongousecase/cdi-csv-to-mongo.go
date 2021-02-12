package mongousecase

import (
	"github.com/iamseki/csv-to-db/app/protocols"
)

type CDIsToMongoDB struct {
	Repository protocols.InsertCDIsRepository
	Reader     protocols.CDIReader
}

func NewCDIsToMongoDB(repository protocols.InsertCDIsRepository, reader protocols.CDIReader) *CDIsToMongoDB {
	return &CDIsToMongoDB{
		Repository: repository,
		Reader:     reader,
	}
}

func (cdiToMongo *CDIsToMongoDB) ConvertCSVIntoDB() error {
	cdis := cdiToMongo.Reader.ConvertCDIFromCSV()
	err := cdiToMongo.Repository.InsertMany(cdis)
	if err != nil {
		return err
	}
	return nil
}

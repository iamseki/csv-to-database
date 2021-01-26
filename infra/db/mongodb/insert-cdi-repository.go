package mongodb

import (
	"log"

	"github.com/iamseki/csv-to-db/contracts"
)

type InsertCDIsMongoRepository struct {
	Client *Mongo
}

func NewInsertCDIsMongoRepository() *InsertCDIsMongoRepository {
	return &InsertCDIsMongoRepository{
		Client: newMongoConnection(),
	}
}

func (m *InsertCDIsMongoRepository) InsertMany(cdis []contracts.CDI) error {
	cdiCollection := m.Client.getCollection("cdi")

	docsToTestTreshold := int64(100)
	numberOfDocs, err := cdiCollection.EstimatedDocumentCount(m.Client.Ctx)
	if numberOfDocs > docsToTestTreshold {
		log.Fatalln("CDIs already inserted into CDI mongo collection")
	}

	docs := []interface{}{}
	for _, cdi := range cdis {
		docs = append(docs, cdi)
	}

	_, err = cdiCollection.InsertMany(m.Client.Ctx, docs)
	if err != nil {
		return err
	}
	return nil
}

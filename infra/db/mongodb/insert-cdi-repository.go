package mongodb

import "github.com/iamseki/csv-to-db/contracts"

type InsertCDIsMongoRepository struct {
	Client *Mongo
}

func NewInsertCDIsMongoRepository() *InsertCDIsMongoRepository {
	return &InsertCDIsMongoRepository{
		Client: newMongoConnection(),
	}
}

func (m *InsertCDIsMongoRepository) InsertMany(cdis []contracts.CDI) error {
	resultCollection := m.Client.getCollection("cdi")

	docs := []interface{}{}
	for _, cdi := range cdis {
		docs = append(docs, cdi)
	}

	_, err := resultCollection.InsertMany(m.Client.Ctx, docs)
	if err != nil {
		return err
	}
	return nil
}

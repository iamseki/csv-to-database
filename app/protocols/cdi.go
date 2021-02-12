package protocols

import "github.com/iamseki/csv-to-db/domain"

type CDIReader interface {
	ConvertCDIFromCSV() []domain.CDI
}

type InsertCDIsRepository interface {
	InsertMany([]domain.CDI) error
}

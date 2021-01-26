package contracts

import "time"

type CDI struct {
	Value float32   `bson:"value" json:"value"`
	Date  time.Time `bson:"date" json:"date"`
}

type CDIReader interface {
	ConvertCDIFromCSV() []CDI
}

type InsertCDIsRepository interface {
	InsertMany([]CDI) error
}

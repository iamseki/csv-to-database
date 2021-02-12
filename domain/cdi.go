package domain

import "time"

type CDI struct {
	Value float32   `bson:"value" json:"value"`
	Date  time.Time `bson:"date" json:"date"`
}

type CSVToDatabaseWriter interface {
	ConvertCSVIntoDB() error
}

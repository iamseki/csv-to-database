package domain

type CSVToDatabaseWriter interface {
	ConvertCSVIntoDB() error
}

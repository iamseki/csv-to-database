package csv

import (
	"log"

	"github.com/iamseki/csv-to-db/contracts"
)

type CDIParser struct {
	filename string
}

func NewCDIFromCSV(filename string) *CDIParser {
	return &CDIParser{filename}
}

func (c *CDIParser) ConvertCDIFromCSV() []contracts.CDI {
	var CDIs []contracts.CDI
	standardSampleLayout := "02/01/2006"
	DATE_INDEX, VALUE_INDEX := 1, 2

	r, os := getCSVReader(c.filename)
	defer os.Close()

	header := getCSVNextRecord(r)
	log.Println(header)
	for {
		record := getCSVNextRecord(r)
		if record == nil {
			break
		}

		date := parseDate(record[DATE_INDEX], standardSampleLayout)
		cdi := parseFloat32(record[VALUE_INDEX])
		CDIs = append(CDIs, contracts.CDI{Value: cdi, Date: date})
	}

	return CDIs
}

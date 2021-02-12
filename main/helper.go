package main

import "flag"

type flags struct {
	cdiFilename string
}

func parseProgramFlags() flags {
	var options flags
	flag.StringVar(&options.cdiFilename, "cdi", "CDI_Prices.csv", "csv filename to convert, e.g --cdi=CDI_Prices.csv")
	flag.Parse()

	return options
}

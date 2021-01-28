package main

import "flag"

type flags struct {
	filename string
}

func parseProgramFlags() flags {
	var options flags
	flag.StringVar(&options.filename, "filename", "CDI_Prices.csv", "csv filename to convert, e.g --filename=CDI_Prices.csv")
	flag.Parse()

	return options
}

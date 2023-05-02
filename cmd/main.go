package main

import (
	"log"

	"github.com/qazaqpyn/uk-crimes/pkg/ukcrimes"
)

const (
	latitude  = "52.640961"
	longitude = "-1.126371"
)

func main() {
	uc, err := ukcrimes.New()
	if err != nil {
		log.Fatal(err)
	}

	report, err := uc.ScanCoordinates(longitude, latitude)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(report.String(5))
}

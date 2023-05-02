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

	// Init the client
	uc, err := ukcrimes.New()
	if err != nil {
		log.Fatal(err)
	}

	// Send coordinates for crime check
	report, err := uc.ScanCoordinates(longitude, latitude)
	if err != nil {
		log.Fatal(err)
	}

	// Print readable report
	log.Println(report.String(5))
}

package client

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qazaqpyn/uk-crimes/internal/responses"
)

// Default settings
const (
	ukApisUrl              = "https://data.police.uk/api"
	requestTimeout         = 2 * time.Second
	okMsg                  = "[OK]"
	errorMsg               = "[ERROR]"
	sucessSubmit           = "successfully submitted coordinates for crime checks!"
	coordinateNotSpecified = "coordinate numbers not specified"
)

// Client is the custom UKCrimes client containing apiUrl
type Client struct {
	client *http.Client
	apiUrl string
}

// NewClient return UKCrimes client with passed default settings
func NewClient() *Client {
	return &Client{
		client: &http.Client{},
		apiUrl: ukApisUrl,
	}
}

// ScanCoordinates send coordinates for checking and gets report back
func (c Client) ScanCoordinates(longitude, latitude string) (*responses.Report, error) {
	if longitude == "" && latitude == "" {
		return nil, fmt.Errorf("$s $s", errorMsg, coordinateNotSpecified)
	}

	coordinates := responses.Location{
		Longitude: longitude,
		Latitude:  latitude,
	}
	reportResp, err := c.postCoordinate(coordinates)
	if err != nil {
		return &responses.Report{}, fmt.Errorf("%s %s", errorMsg, err)
	}

	log.Printf("%s %s,%s %s", okMsg, longitude, latitude, sucessSubmit)

	// not required to wait, can be removed
	time.Sleep(requestTimeout)
	return reportResp, nil
}

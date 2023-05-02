package client

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qazaqpyn/uk-crimes/internal/responses"
)

const (
	ukApisUrl              = "https://data.police.uk/api"
	requestTimeout         = 2 * time.Second
	okMsg                  = "[OK]"
	errorMsg               = "[ERROR]"
	sucessSubmit           = "successfully submitted coordinates for crime checks!"
	coordinateNotSpecified = "coordinate numbers not specified"
)

type Client struct {
	client *http.Client
	apiUrl string
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
		apiUrl: ukApisUrl,
	}
}

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
	time.Sleep(requestTimeout)
	return reportResp, nil
}

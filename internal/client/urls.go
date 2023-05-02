package client

import (
	"encoding/json"
	"io"

	"github.com/qazaqpyn/uk-crimes/internal/responses"
)

const (
	urlStreetCrime       = "/crimes-street/all-crime"
	paramLat             = "?lat="
	paramLon             = "&lng="
	requestLimitExceeded = "request rate limit exceeded"
)

func (c Client) postCoordinate(coordinates responses.Location) (*responses.Report, error) {
	requestUrl := c.apiUrl + urlStreetCrime
	params := paramLat + coordinates.Latitude + paramLon + coordinates.Longitude

	resp, err := c.client.Get(requestUrl + params)
	if err != nil {
		return &responses.Report{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &responses.Report{}, err
	}

	//parse response
	var scan []responses.Crime
	err = json.Unmarshal(body, &scan)
	if err != nil {
		return &responses.Report{}, err
	}

	report := responses.Report{
		Location: coordinates,
		Crime:    scan,
	}

	return &report, nil
}

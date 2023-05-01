package responses

import (
	"fmt"
	"time"
)

type Location struct {
	Latitude  string
	Longitude string
	Name      string
}

type Crime struct {
	Category string
	Location
	Date time.Time
	ID   int
}

type Report struct {
	Location
	Crime []Crime
}

func (r Report) String() string {
	if len(r.Crime) == 0 {
		return fmt.Sprintf("[NO CRIMES] LONGITUDE: %s, LATITUDE: %s", r.Location.Longitude, r.Location.Latitude)
	}

	report := fmt.Sprintf("[CRIMES] LONGITUDE: %s, LATITUDE: %s\n", r.Location.Longitude, r.Location.Latitude)
	for i, crime := range r.Crime {
		report += fmt.Sprintf("\t[CRIME%d] CATEGORY: %s DATE: %s\n", i, crime.Category, crime.Date)
		report += fmt.Sprintf("\t\tLOCATION: %s \t ID: %d\n", crime.Location.Name, crime.ID)
	}

	return report
}

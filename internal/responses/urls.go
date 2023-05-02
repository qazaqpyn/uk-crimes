package responses

import (
	"fmt"
)

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Street    `json:"street"`
}

type Street struct {
	Name string `json:"name"`
}

type Crime struct {
	Category string `json:"category"`
	Location `json:"location"`
	Date     string `json:"month"`
	ID       int    `json:"id"`
}

type Report struct {
	Location
	Crime []Crime
}

func (r Report) String(top uint) string {
	if len(r.Crime) == 0 {
		return fmt.Sprintf("[NO CRIMES] LONGITUDE: %s, LATITUDE: %s", r.Location.Longitude, r.Location.Latitude)
	}

	report := fmt.Sprintf("[CRIMES] LONGITUDE: %s, LATITUDE: %s\n", r.Location.Longitude, r.Location.Latitude)
	for i, crime := range r.Crime[:top] {
		report += fmt.Sprintf("\t[CRIME%d] CATEGORY: %s DATE: %s\n", i, crime.Category, crime.Date)
		report += fmt.Sprintf("\t\tLOCATION: %s \t ID: %d\n", crime.Location.Street.Name, crime.ID)
	}

	return report
}

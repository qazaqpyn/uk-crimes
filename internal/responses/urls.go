package responses

import (
	"fmt"
)

// Location is the sending and receiving types
type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Street    `json:"street"`
}

// Street is the naming of crime location
type Street struct {
	Name string `json:"name"`
}

// Crime is the response that received
type Crime struct {
	Category string `json:"category"`
	Location `json:"location"`
	Date     string `json:"month"`
	ID       int    `json:"id"`
}

// Report is the response that contains all Crimes information
type Report struct {
	Location
	Crime []Crime
}

// String return Report in readable format and receives the arguments of Top number, if top = 0 => all data. Fe: top = 5 => 5 latest crimes at radious of coordinate
func (r Report) String(top uint) string {
	if len(r.Crime) == 0 {
		return fmt.Sprintf("[NO CRIMES] LONGITUDE: %s, LATITUDE: %s", r.Location.Longitude, r.Location.Latitude)
	}

	report := fmt.Sprintf("[CRIMES] LONGITUDE: %s, LATITUDE: %s\n", r.Location.Longitude, r.Location.Latitude)
	if top != 0 {
		for i, crime := range r.Crime[:top] {
			report += fmt.Sprintf("\t[CRIME%d] CATEGORY: %s \t DATE: %s\n", i, crime.Category, crime.Date)
			report += fmt.Sprintf("\t\t LOCATION: %s \t ID: %d\n", crime.Location.Street.Name, crime.ID)
		}
	} else {
		for i, crime := range r.Crime {
			report += fmt.Sprintf("\t[CRIME%d] CATEGORY: %s \t DATE: %s\n", i, crime.Category, crime.Date)
			report += fmt.Sprintf("\t\t LOCATION: %s \t ID: %d\n", crime.Location.Street.Name, crime.ID)
		}
	}

	return report
}

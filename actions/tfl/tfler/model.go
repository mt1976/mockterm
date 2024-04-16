package tfler

type Line struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Code        string   //`json:"code"`
	Status      string   //`json:"status"`
	StatusCode  string   //`json:"status_code"`
	Type        []string //`json:"type"`
	Disruptions []string //`json:"disruptions"`
}

type LineDetail struct {
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Status     string    `json:"status"`
	StatusCode string    `json:"status_code"`
	Type       []string  `json:"type"`
	Stations   []Station `json:"stations"`
	Narrative  string    `json:"narrative"`
}

type Station struct {
	Name       string   `json:"name"`
	Code       string   `json:"code"`
	Status     string   `json:"status"`
	StatusCode string   `json:"status_code"`
	Type       []string `json:"type"`
	LineCode   string   `json:"line_code"`
}

type StationDetail struct {
	Name       string   `json:"name"`
	Code       string   `json:"code"`
	Status     string   `json:"status"`
	StatusCode string   `json:"status_code"`
	Type       []string `json:"type"`
	LineCode   string   `json:"line_code"`
	LineName   string   `json:"line_name"`
	Narrative  string   `json:"narrative"`
}

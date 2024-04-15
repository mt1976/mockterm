package tfler

type Line struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Line       string `json:"line"`
	StatusCode string `json:"status_code"`
}

type LineDetail struct {
	Name     string    `json:"name"`
	Code     string    `json:"code"`
	Stations []Station `json:"stations"`
}

type Station struct {
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Status   string   `json:"status"`
	Type     []string `json:"type"`
	LineCode string   `json:"line_code"`
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

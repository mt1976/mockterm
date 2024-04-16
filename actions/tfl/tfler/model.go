package tfler

import "time"

type Line struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Code        string   `json:"name"`
	Status      string   //`json:"status"`
	StatusCode  string   //`json:"status_code"`
	Type        []string //`json:"type"`
	Disruptions []string //`json:"disruptions"`
	LastUpdated time.Time
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

type LineImport []struct {
	Type          string        `json:"$type"`
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	ModeName      string        `json:"modeName"`
	Disruptions   []interface{} `json:"disruptions"`
	Created       time.Time     `json:"created"`
	Modified      time.Time     `json:"modified"`
	LineStatuses  []interface{} `json:"lineStatuses"`
	RouteSections []interface{} `json:"routeSections"`
	ServiceTypes  []struct {
		Type string `json:"$type"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"serviceTypes"`
	Crowding struct {
		Type string `json:"$type"`
	} `json:"crowding"`
}

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

type LineDetailImport []struct {
	Type          string   `json:"$type"`
	NaptanID      string   `json:"naptanId"`
	Modes         []string `json:"modes"`
	IcsCode       string   `json:"icsCode"`
	StopType      string   `json:"stopType"`
	StationNaptan string   `json:"stationNaptan"`
	Lines         []struct {
		LineType string `json:"$type"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		URI      string `json:"uri"`
		Type     string `json:"type"`
		Crowding struct {
			Type string `json:"$type"`
		} `json:"crowding"`
		RouteType string `json:"routeType"`
		Status    string `json:"status"`
	} `json:"lines"`
	LineGroup []struct {
		Type              string   `json:"$type"`
		NaptanIDReference string   `json:"naptanIdReference,omitempty"`
		StationAtcoCode   string   `json:"stationAtcoCode"`
		LineIdentifier    []string `json:"lineIdentifier"`
	} `json:"lineGroup"`
	LineModeGroups []struct {
		Type           string   `json:"$type"`
		ModeName       string   `json:"modeName"`
		LineIdentifier []string `json:"lineIdentifier"`
	} `json:"lineModeGroups"`
	Status               bool   `json:"status"`
	ID                   string `json:"id"`
	CommonName           string `json:"commonName"`
	PlaceType            string `json:"placeType"`
	AdditionalProperties []struct {
		Type            string `json:"$type"`
		Category        string `json:"category"`
		Key             string `json:"key"`
		SourceSystemKey string `json:"sourceSystemKey"`
		Value           string `json:"value"`
	} `json:"additionalProperties"`
	Children []struct {
		Type                 string        `json:"$type"`
		NaptanID             string        `json:"naptanId"`
		Modes                []interface{} `json:"modes"`
		IcsCode              string        `json:"icsCode"`
		StationNaptan        string        `json:"stationNaptan"`
		Lines                []interface{} `json:"lines"`
		LineGroup            []interface{} `json:"lineGroup"`
		LineModeGroups       []interface{} `json:"lineModeGroups"`
		Status               bool          `json:"status"`
		ID                   string        `json:"id"`
		CommonName           string        `json:"commonName"`
		PlaceType            string        `json:"placeType"`
		AdditionalProperties []interface{} `json:"additionalProperties"`
		Children             []interface{} `json:"children"`
		Lat                  float64       `json:"lat"`
		Lon                  float64       `json:"lon"`
	} `json:"children"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	HubNaptanCode string  `json:"hubNaptanCode,omitempty"`
}

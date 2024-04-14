package tfl

type listOfValidModes []struct {
	IsTflService       bool   `json:"isTflService"`
	IsFarePaying       bool   `json:"isFarePaying"`
	IsScheduledService bool   `json:"isScheduledService"`
	ModeName           string `json:"modeName"`
}

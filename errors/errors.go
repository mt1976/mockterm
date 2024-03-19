package support

import "errors"

const (
	ErrPlexInit           string = "unable to connect to server"
	ErrPlexConnectionTest string = "unable to test connection to server"
	ErrPlexConnect        string = "unable to connect to %v"
	ErrLibraryResponse    string = "unable to get libraries from %v"
	ErrDashboardNoHost    string = "dashboard: No default host set"
)

// Weather

const (
	ErrOpenWeather string = "failed to initialize OpenWeatherMap: %v"
)

var (
	ErrBaudRateError = errors.New("invalid baud rate")
)

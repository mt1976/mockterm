package support

import "errors"

const (
// ErrPlexInit           string = "unable to connect to server"
// ErrPlexConnectionTest string = "unable to test connection to server"
// ErrPlexConnect        string = "unable to connect to %v"
// ErrDashboardNoHost    string = "dashboard: No default host set"
)

var (
	ErrLibraryResponse    = errors.New("unable to get libraries from %v")
	ErrPlexInit           = errors.New("unable to connect to server")
	ErrPlexConnectionTest = errors.New("unable to test connection to server")
	ErrPlexConnect        = errors.New("unable to connect to %v")
	ErrDashboardNoHost    = errors.New("dashboard: No default host set")
)

// Weather

var (
	ErrOpenWeather = errors.New("failed to initialize OpenWeatherMap: %v")
)

var (
	ErrBaudRateError = errors.New("invalid baud rate")
)

var (
	ErrReadingFolder = errors.New("error reading folder %v")
	ErrOpeningFile   = errors.New("error opening file %v")
	ErrWritingFile   = errors.New("error writing file %v")
)

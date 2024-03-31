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
	ErrReadingFolder           = errors.New("error reading folder %v")
	ErrOpeningFile             = errors.New("error opening file %v")
	ErrWritingFile             = errors.New("error writing file %v")
	ErrFileInfo                = errors.New("error getting file info")
	ErrFileDoesNotExist        = errors.New("error file does not exist %v") //
	ErrDiskSizeUsage    string = "Prints file sizes in bytes, kilobytes, megabytes, and gigabytes\n\nUsage: sz <file> <file> <file>"
	ErrNoFilesSpecified string = "No files specified"
)

// CleanFileNames
var (
	ErrProcessingFiles  = errors.New("error processing files")
	ErrCleaningFileName = errors.New("error cleaning file name")
	ErrRenamingFile     = errors.New("failed to rename file %v")
)

// TidyFiles
var (
	ErrNoEmptyDirectories          = errors.New("unable to find empty directories %v")
	ErrUnableToRemoveDirectories   = errors.New("unable to remove empty directories %v")
	ErrUnableToFindFiles           = errors.New("unable to find files %v")
	ErrUnableToResolvePath         = errors.New("unable to resolve path %v")
	ErrNoPathSpecified             = errors.New("no path specification specified %v")
	ErrInvalidPath                 = errors.New("the path provided is not valid %v")
	ErrInvalidPathSpecialDirectory = errors.New("the path provided is the root or home directory")
	ErrFailedToChangeDirectory     = errors.New("failed to change directory to %v %v")
	ErrNotADirectory               = errors.New("%v is not a directory")
	ErrNotAFile                    = errors.New("%v is not a file")
	ErrDirectoryEmpty              = errors.New("directory %v is empty")
)

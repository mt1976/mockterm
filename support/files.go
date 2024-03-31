package support

import (
	"os"
	"time"

	term "github.com/mt1976/crt"
	cfgr "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

var config = cfgr.Configuration

// GetFilesList returns a list of files in the specified base folder.
// If an error occurs, an error message is displayed using the provided terminal.
func GetFilesList(t *term.ViewPort, baseFolder string) []os.DirEntry {
	files, err := os.ReadDir(baseFolder)
	if err != nil {
		t.Error(errs.ErrReadingFolder, baseFolder, err.Error())
		return nil
	}
	return files
}

// GetTimeStamp returns the current date in the format "20060102".
func GetTimeStamp() string {
	return time.Now().Format(config.TimeStampFormat)
}

// OpenFile opens a file for writing.
//
// Parameters:
// t - the terminal to use for output.
// filename - the name of the file to open.
//
// Returns:
// a file pointer, or an error if any occurred while opening the file.
func OpenFile(t *term.ViewPort, filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Error(errs.ErrOpeningFile, t.Formatters.Bold(filename), err.Error())
		return nil, err
	}
	return file, nil
}

// WriteStringSliceToFile writes a slice of strings to a file.
//
// Parameters:
// t - the terminal to use for output.
// file - the file to write to.
// content - the slice of strings to write to the file.
//
// Returns:
// an error if any occurred while writing to the file.
func WriteStringSliceToFile(t *term.ViewPort, file *os.File, content []string) error {
	for _, line := range content {
		_, err := file.WriteString(line + lang.SymNewline)
		if err != nil {
			t.Error(errs.ErrWritingFile, t.Formatters.Bold(file.Name()), err.Error())
			return err
		}
	}
	return nil
}

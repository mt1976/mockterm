package files

import (
	"io"
	"os"

	page "github.com/mt1976/crt/page"
	term "github.com/mt1976/crt/terminal"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

// GetList returns a list of files in the specified base folder.
// If an error occurs, an error message is displayed using the provided terminal.
func GetList(t *term.ViewPort, baseFolder string) []os.DirEntry {
	files, err := os.ReadDir(baseFolder)
	if err != nil {
		t.Error(errs.ErrReadingFolder, baseFolder, err.Error())
		return nil
	}
	return files
}

// Open opens a file for writing.
//
// Parameters:
// t - the terminal to use for output.
// filename - the name of the file to open.
//
// Returns:
// a file pointer, or an error if any occurred while opening the file.
func Open(p *page.Page, filename string) (*os.File, error) {
	file, err := os.Open(filename)
	//, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		p.Error(errs.ErrOpeningFile, filename, err.Error())
		return nil, err
	}

	defer file.Close() // Make sure to close the file when the function returns

	return file, nil
}

// WriteStringSlice writes a slice of strings to a file.
//
// Parameters:
// t - the terminal to use for output.
// file - the file to write to.
// content - the slice of strings to write to the file.
//
// Returns:
// an error if any occurred while writing to the file.
func WriteStringSlice(p *page.Page, file *os.File, content []string) error {
	for _, line := range content {
		// _, err := file.WriteString(line + lang.SymNewline)
		// if err != nil {
		// 	return err
		// }
		err := WriteString(p, file, line+lang.SymNewline)
		if err != nil {
			p.Error(errs.ErrWritingFile, file.Name(), err.Error())
			return err
		}
	}
	return nil
}

func WriteString(p *page.Page, file *os.File, content string) error {
	_, err := file.WriteString(content)
	if err != nil {
		p.Error(errs.ErrWritingFile, file.Name(), err.Error())
		return err
	}
	return nil
}

func ReadFile(p *page.Page, fileName string) ([]byte, error) {

	// // Open the file for reading
	// file, err := os.Open(fileName)
	// if err != nil {
	// 	p.Error(err, fileName)
	// 	//fmt.Println("Error opening file:", err)
	// 	return []byte{}, err
	// }
	// defer file.Close() // Make sure to close the file when the function returns

	file, err := Open(p, fileName)
	if err != nil {
		p.Error(err, fileName)
		return []byte{}, err
	}

	// read the file into a string
	data, err := io.ReadAll(file)
	if err != nil {
		p.Error(err, fileName)
		return []byte{}, err
	}
	return data, nil
}

func Read(p *page.Page, file *os.File) ([]byte, error) {
	// read the file into a string
	data, err := io.ReadAll(file)
	if err != nil {
		p.Error(err, file.Name())
		return []byte{}, err
	}
	return data, nil
}

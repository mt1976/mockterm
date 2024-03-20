package support

import (
	"os"
	"time"

	term "github.com/mt1976/crt"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

func GetFilesList(crt term.Crt, baseFolder string) []os.DirEntry {
	files, err := os.ReadDir(baseFolder)
	if err != nil {
		//fmt.Printf("%s Error reading folder %s: %v\n", chNormal, baseFolder, err)
		crt.Error(errs.ErrReadingFolder, baseFolder, err.Error())
		return nil
	}
	return files
}

func GetTimeStamp() string {
	return time.Now().Format("20060102")
}

func OpenFile(t term.Crt, filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Printf("%s Error opening file %s: %v\n", crt.CHnormal, filename, err)
		t.Error(errs.ErrOpeningFile, t.Formatters.Bold(filename), err.Error())
		return nil, err
	}
	return file, nil
}

func WriteStringSliceToFile(t term.Crt, file *os.File, content []string) error {
	for _, line := range content {
		_, err := file.WriteString(line + lang.SymNewline)
		if err != nil {
			//fmt.Printf("%s Error writing to file %s: %v\n", crt.CHnormal, file.Name(), err)
			t.Error(errs.ErrWritingFile, t.Formatters.Bold(file.Name()), err.Error())
			return err
		}
	}
	return nil
}

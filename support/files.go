package support

import (
	"os"
	"time"

	term "github.com/mt1976/crt"
	errs "github.com/mt1976/mockterm/errors"
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

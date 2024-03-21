package cleanfilenames

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/lithammer/shortuuid/v3"
	term "github.com/mt1976/crt"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	sppt "github.com/mt1976/mockterm/support"
)

var itemCount int = 0
var debugMode bool = false
var cfg = conf.Configuration

func Run(t term.Crt, debugModeIn bool, cleanPathIn, messageIn string) {
	//TODO add Input to get path to clean
	//TODO add Check to see if path exists, if not ask for input again
	//TODO add Input to select Mode, debug or normal
	//TODO re-present the path and mode selected, ask for permission to proceed
	//TODO if no permission go back, otherwise proceed.

	debugMode = debugModeIn

	t.Print(lang.TxtStartingCleanFileNames + t.Formatters.DQuote(t.Formatters.Bold(cleanPathIn)))
	t.Blank()

	baseFolder := "."

	fileList := sppt.GetFilesList(t, baseFolder)
	if len(fileList) == 0 {
		t.Shout(fmt.Sprintf(lang.TxtNoFilesFoundInFolder, baseFolder))
		return
	}

	t.Print(fmt.Sprintf(lang.TxtProcessingNFilesIn, len(fileList), messageIn))
	t.Blank()

	for _, file := range fileList {
		err := cleanFileName(t, file, baseFolder)
		if err != nil {
			t.Error(errs.ErrProcessingFiles, err.Error())
			return
		}
	}

	t.Break()

	if itemCount > 0 {
		t.Print(fmt.Sprintf(lang.TxtProcessedNFilesIn, itemCount, cleanPathIn))
	} else {
		t.Print(fmt.Sprintf(lang.TxtNoFilesProcessed, cleanPathIn))
	}
}

func cleanFileName(t term.Crt, info fs.DirEntry, path string) error {

	cleanName, err := getCleanName(t, info.Name())
	if err != nil {
		t.Error(errs.ErrCleaningFileName, info.Name(), err.Error())
		return errs.ErrCleaningFileName
	}

	if cleanName == lang.TxtOnlyFansFilename {
		// Rename the file to OnlyFans_Date_Time.mp4
		id := uuid.New()
		cleanName = lang.TxtOnlyFans + time.Now().Format(cfg.OnlyFansDateTimeFormat) + id + lang.FileExtensionMP4
	}

	if cleanName != info.Name() {
		renameFile(t, path, cleanName, info.Name())
		itemCount++
	}
	return nil
}

func getCleanName(t term.Crt, fileName string) (string, error) {
	//fmt.Printf("%s Cleaning file name '%s'\n", support.PFX, name)
	newFileName := fileName

	// Remove all characters that are not in the ValidChars list
	for _, c := range fileName {
		if !strings.Contains(strings.Join(cfg.ValidFileNameCharacters, ""), string(c)) {
			newFileName = strings.ReplaceAll(newFileName, string(c), "")
		}
	}
	newFileName = strings.ReplaceAll(newFileName, "_", " ")
	newFileName = strings.ReplaceAll(newFileName, "-", " ")

	// Remove all double spaces
	newFileName = t.Formatters.TrimRepeatingCharacters(newFileName, " ")
	newFileName = t.Formatters.TrimRepeatingCharacters(newFileName, ".")
	newFileName = t.Formatters.TrimRepeatingCharacters(newFileName, "-")
	newFileName = t.Formatters.TrimRepeatingCharacters(newFileName, "*")
	newFileName = strings.TrimLeft(newFileName, " ")
	newFileName = strings.TrimLeft(newFileName, "-")
	newFileName = strings.TrimLeft(newFileName, " ")
	newFileName = strings.TrimLeft(newFileName, "-")
	//fmt.Printf("%s New file name '%s'\n", support.PFX, newName)
	return newFileName, nil
}

func renameFile(t term.Crt, path string, newFileName string, oldFileName string) {
	newPath := filepath.Join(filepath.Dir(path), newFileName)
	oldPath := filepath.Join(filepath.Dir(path), oldFileName)
	err := error(nil)

	if !debugMode {
		err = os.Rename(oldPath, newPath)
	}

	if err != nil {
		t.Error(errs.ErrRenamingFile, path, err.Error())
	} else {
		t.Print(fmt.Sprintf(lang.TxtRemamedFile, oldFileName, newPath))
	}
}

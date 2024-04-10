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
	file "github.com/mt1976/crt/filechooser"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	sppt "github.com/mt1976/mockterm/support"
)

var itemCount int = 0
var debugMode bool = false
var cfg = conf.Configuration
var results = []string{}

func Run(t *term.ViewPort, debugModeIn bool, basePath string) {

	debugMode = debugModeIn

	p := t.NewPage(lang.TxtCleanFileNames)
	resultsAdd(upcase(p, lang.TxtCleanFileNamesReport))

	if basePath == "" {
		home, err := file.UserHome()
		if err != nil {
			p.Error(err)
		}
		basePath = home
	}

	if debugMode {
		p.AddFieldValuePair(lang.TxtMode, lang.TxtDebugMode)
	} else {
		p.AddFieldValuePair(lang.TxtMode, lang.TxtLiveMode)
	}

	baseFolder, _, err := file.FileChooser(basePath, file.DirectoriesOnly)
	if err != nil {
		p.Error(err, "file chooser error")
	}

	p.AddFieldValuePair(lang.TxtPath, baseFolder)
	resultsAdd(lang.TxtPath + " " + baseFolder)
	p.AddBlankRow()
	p.AddParagraph(lang.TxtCleanFileNamesDescription)

	msg := lang.TxtStartingCleanFileNames + t.Formatters.DQuote(t.Formatters.Bold(basePath))
	p.Info(msg)
	resultsAdd(msg)
	p.AddBlankRow()
	resultsAdd("")

	ok, err := p.Display_Confirmation(lang.TxtAreYouSureYouWantToProceed)
	if err != nil {
		p.Error(err, "unable to get user response")
	}
	if !ok {
		//fmt.Printf("%s Exiting\n", PFY)
		p.Info(lang.TxtQuittingMessage)
		return
	}

	fileList := sppt.GetFilesList(t, baseFolder)
	if len(fileList) == 0 {
		p.Error(errs.ErrDirectoryEmpty, baseFolder)
		resultsAdd(errs.ErrDirectoryEmpty.Error())
		resultsAdd(baseFolder)
		return
	}

	msg = fmt.Sprintf(lang.TxtProcessingNFilesIn, len(fileList), baseFolder)
	p.Info(msg)
	resultsAdd(msg)

	for _, file := range fileList {
		err := cleanFileName(p, file, baseFolder)
		if err != nil {
			p.Error(errs.ErrProcessingFiles, err.Error())
			resultsAdd(errs.ErrProcessingFiles.Error())
			resultsAdd(err.Error())
			return
		}
	}

	if itemCount > 0 {
		msg := fmt.Sprintf(lang.TxtProcessedNFilesIn, itemCount, basePath)
		p.Success(msg)
		resultsAdd(msg)
	} else {
		msg := fmt.Sprintf(lang.TxtNoFilesProcessed, basePath)
		p.Info(msg)
		resultsAdd(msg)
	}

	q := t.NewPage(lang.TxtCleanFileNamesResults)
	q.AddParagraph(results)
	q.Display_Actions()

}

func cleanFileName(p *term.Page, info fs.DirEntry, path string) error {
	resultsAdd("Processing : " + info.Name())
	cleanName, err := getCleanName(p, info.Name())
	if err != nil {
		p.Error(errs.ErrCleaningFileName, info.Name(), err.Error())
		resultsAdd(errs.ErrCleaningFileName.Error())
		resultsAdd(info.Name())
		resultsAdd(err.Error())
		return errs.ErrCleaningFileName
	}

	if cleanName == lang.TxtOnlyFansFilename {
		// Rename the file to OnlyFans_Date_Time.mp4
		id := uuid.New()
		cleanName = lang.TxtOnlyFans + time.Now().Format(cfg.OnlyFansDateTimeFormat) + id + lang.FileExtensionMP4
	}

	if cleanName != info.Name() {
		resultsAdd(fmt.Sprintf(lang.TxtRemamedFile, info.Name(), cleanName))
		renameFile(p, path, cleanName, info.Name())
		itemCount++
	}
	return nil
}

func getCleanName(p *term.Page, fileName string) (string, error) {
	//fmt.Printf("%s Cleaning file name '%s'\n", support.PFX, name)
	t := p.ViewPort()
	f := t.Formatters
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
	newFileName = f.TrimRepeatingCharacters(newFileName, " ")
	newFileName = f.TrimRepeatingCharacters(newFileName, ".")
	newFileName = f.TrimRepeatingCharacters(newFileName, "-")
	newFileName = f.TrimRepeatingCharacters(newFileName, "*")
	newFileName = strings.TrimLeft(newFileName, " ")
	newFileName = strings.TrimLeft(newFileName, "-")
	newFileName = strings.TrimLeft(newFileName, " ")
	newFileName = strings.TrimLeft(newFileName, "-")
	//fmt.Printf("%s New file name '%s'\n", support.PFX, newName)
	return newFileName, nil
}

func renameFile(p *term.Page, path string, newFileName string, oldFileName string) {
	newPath := filepath.Join(filepath.Dir(path), newFileName)
	oldPath := filepath.Join(filepath.Dir(path), oldFileName)
	err := error(nil)

	if !debugMode {
		err = os.Rename(oldPath, newPath)
	}

	if err != nil {
		p.Error(errs.ErrRenamingFile, path, err.Error())
	} else {
		p.Info(fmt.Sprintf(lang.TxtRemamedFile, oldFileName, newPath))
	}
}

func resultsAdd(msg string) {
	results = append(results, msg)
}

func upcase(t *term.Page, msg string) string {
	return t.ViewPort().Formatters.Upcase(msg)
}

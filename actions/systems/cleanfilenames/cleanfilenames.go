package cleanfilenames

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/lithammer/shortuuid/v3"
	file "github.com/mt1976/crt/filechooser"
	page "github.com/mt1976/crt/page"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/actions/systems/cleanfilenames/lang"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	clng "github.com/mt1976/mockterm/language"
	files "github.com/mt1976/mockterm/support/files"
	mode "github.com/mt1976/mockterm/support/modes"
)

var itemCount int = 0
var debugMode bool = false
var cfg = conf.Configuration
var results = []string{}

func Run(t *term.ViewPort, m mode.Modality, basePath string) {
	if m.Is(mode.DEBUG) {
		debugMode = true
	}
	p := page.NewPage(t, lang.Title)
	resultsAdd(upcase(p, lang.StorageReport.Text()))

	if basePath == "" {
		home, err := file.UserHome()
		if err != nil {
			p.Error(err)
		}
		basePath = home
	}

	p.AddFieldValuePair(lang.Mode.Text(), m.String())

	baseFolder, _, err := file.FileChooser(basePath, file.DirectoriesOnly)
	if err != nil {
		p.Error(err, "file chooser error")
		return
	}

	p.AddFieldValuePair(clng.Path.Text(), baseFolder)
	resultsAdd(clng.Path.Text() + " " + baseFolder)
	p.AddBlankRow()
	p.AddParagraph(lang.CleanFileNamesDescription.String())

	msg := lang.StartCleanFileNames.Text() + t.Formatters.DQuote(t.Formatters.Bold(basePath))
	p.Info(clng.New(msg))
	resultsAdd(msg)
	p.AddBlankRow()
	resultsAdd("")

	ok, err := p.Display_Confirmation(lang.AreYouSureYouWantToProceed)
	if err != nil {
		p.Error(err, "unable to get user response")
	}
	if !ok {
		//fmt.Printf("%s Exiting\n", PFY)
		p.Info(clng.Quitting)
		return
	}

	fileList := files.GetList(t, baseFolder)
	if len(fileList) == 0 {
		p.Error(errs.ErrDirectoryEmpty, baseFolder)
		resultsAdd(errs.ErrDirectoryEmpty.Error())
		resultsAdd(baseFolder)
		return
	}

	msg = fmt.Sprintf(lang.ProcessedNFilesIn.Text(), len(fileList), baseFolder)
	p.Info(clng.New(msg))
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
		msg := fmt.Sprintf(lang.ProcessedNFilesIn.Text(), itemCount, basePath)
		p.Success(clng.New(msg))
		resultsAdd(msg)
	} else {
		msg := fmt.Sprintf(lang.NoFilesProcessed.Text(), basePath)
		p.Info(clng.New(msg))
		resultsAdd(msg)
	}

	q := page.NewPage(t, lang.CleanFileNamesResults)
	q.AddParagraph(results)
	q.Display_Actions()

}

func cleanFileName(p *page.Page, info fs.DirEntry, path string) error {
	resultsAdd("Processing : " + info.Name())
	cleanName, err := getCleanName(p, info.Name())
	if err != nil {
		p.Error(errs.ErrCleaningFileName, info.Name(), err.Error())
		resultsAdd(errs.ErrCleaningFileName.Error())
		resultsAdd(info.Name())
		resultsAdd(err.Error())
		return errs.ErrCleaningFileName
	}

	if cleanName == lang.OnlyFans.Text() {
		// Rename the file to OnlyFans_Date_Time.mp4
		id := uuid.New()
		cleanName = lang.OnlyFans.Text() + time.Now().Format(cfg.OnlyFansDateTimeFormat) + id + lang.MP4.Text()
	}

	if cleanName != info.Name() {
		resultsAdd(fmt.Sprintf(lang.Renamed.Text(), info.Name(), cleanName))
		renameFile(p, path, cleanName, info.Name())
		itemCount++
	}
	return nil
}

func getCleanName(p *page.Page, fileName string) (string, error) {
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

func renameFile(p *page.Page, path string, newFileName string, oldFileName string) {
	newPath := filepath.Join(filepath.Dir(path), newFileName)
	oldPath := filepath.Join(filepath.Dir(path), oldFileName)
	err := error(nil)

	if !debugMode {
		err = os.Rename(oldPath, newPath)
	}

	if err != nil {
		p.Error(errs.ErrRenamingFile, path, err.Error())
	} else {
		p.Info(clng.New(fmt.Sprintf(lang.Renamed.Text(), oldFileName, newPath)))
	}
}

func resultsAdd(msg string) {
	results = append(results, msg)
}

func upcase(t *page.Page, msg string) string {
	return t.ViewPort().Formatters.Upcase(msg)
}

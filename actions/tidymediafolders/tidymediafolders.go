package tidymediafolders

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	term "github.com/mt1976/crt"
	file "github.com/mt1976/crt/filechooser"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	"github.com/ricochet2200/go-disk-usage/du"
)

var (
	fileExtensions = []string{"nfo", "jpeg", "jpg", "bif", "vob", "txt", "png", "me", "exe"}
)

var debugMode bool = true
var cfg = conf.Configuration
var results = []string{}

func Run(t *term.ViewPort, debugModeIn bool, pathIn string) {
	debugMode = debugModeIn

	//new page
	p := t.NewPage(lang.TxtTidyFilesTitle)

	//	t.Print(lang.TxtTidyFilesTitle)
	if pathIn == "" {
		home, err := file.UserHome()
		if err != nil {
			p.Error(err)
		}
		pathIn = home
	}
	path, err := file.ChooseDirectory(pathIn)
	if err != nil {
		p.Error(err, "file chooser error")
	}
	pathIn = path

	if len(pathIn) < 1 {
		p.Error(errs.ErrNoPathSpecified, pathIn)
		return
	}

	//path := os.Args[1]

	if _, err := os.Stat(pathIn); os.IsNotExist(err) {
		p.Error(errs.ErrInvalidPath, err.Error())
		return
	}

	if pathIn == "/" || pathIn == "~" {
		p.Error(errs.ErrInvalidPathSpecialDirectory, pathIn)
		return
	}

	if !debugMode {
		p.AddFieldValuePair(lang.TxtMode, lang.TxtDebugMode)
		//t.Special(t.Formatters.Bold(t.Underline(lang.TxtLiveRun)))
	} else {
		p.AddFieldValuePair(lang.TxtMode, lang.TxtLiveMode)
		//t.Print(t.Underline(lang.TxtTrailRun))
	}

	//t.Print(lang.TxtResolvedPath + realpath(t, pathIn))
	p.AddFieldValuePair(lang.TxtPath, realpath(p, pathIn))
	//fmt.Printf("%s File types to be removed: [%s]\n", support.CHnormal, strings.Join(types, " "))
	//t.Print(fmt.Sprintf(lang.TxtTidyFilesStart, t.Formatters.Bold(strings.Join(fileExtensions, " "))))
	p.AddBlankRow()
	p.Add(fmt.Sprintf(lang.TxtTidyFilesStart, t.Formatters.Bold(strings.Join(fileExtensions, " "))), "", "")
	//var userResponse string
	//fmt.Printf("%s Are you sure you want to proceed? %s(y/n) : %s", PFY, bold, normal)
	ok, err := p.Confirmation(lang.TxtAreYouSureYouWantToProceed)
	if err != nil {
		p.Error(err, "unable to get user response")
	}
	//userResponse := t.Input(lang.TxtAreYouSureYouWantToProceed, lang.OptAreYouSureYouWantToProceed)
	//fmt.Scanln(&userResponse)

	if !ok {
		//fmt.Printf("%s Exiting\n", PFY)
		p.Info(lang.TxtQuittingMessage)
		return
	}

	//	p.Clear()
	//	q := t.NewPage(lang.TxtTidyFilesTitle)
	//	t.Blank()
	//p.AddBlankRow()

	diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore := getDiskInfo(p, pathIn)

	//fmt.Printf("%s Changing directory to %s\n", PFY, path)
	p.Info(fmt.Sprintf(lang.TxtChangingDirectory, t.Formatters.Bold(pathIn)))
	//t.Blank()
	//	p.AddBlankRow()

	err = os.Chdir(pathIn)
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to change directory: %v", PFY, err))
		p.Error(errs.ErrFailedToChangeDirectory, pathIn, err.Error())
		return
	}

	//var wg sync.WaitGroup

	for _, fileExtension := range fileExtensions {
		//wg.Add(idx)
		//fmt.Printf("%s Operation on .%s files completed in %s seconds\n", support.CHspecial, fileExt, runtime)

		//defer wg.Done()
		processFileTypes(p, fileExtension)

	}
	//wg.Wait()

	p.Info(lang.TxtTidyFilesDeletingDirectories)
	startLoopIteration := time.Now()
	if !debugMode {
		removeEmptyDirectories(p)
	} else {
		findEmptyDirectories(p)
	}
	endLoopIteration := time.Now()
	runtime := endLoopIteration.Sub(startLoopIteration)
	p.Success(fmt.Sprintf(lang.TxtTidyFilesDeletingDirectoriesCompleted, t.Formatters.Bold(runtime.String())))
	diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter := getDiskInfo(p, pathIn)

	printStorageReport(p, diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore, diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter)
	q := t.NewPage(lang.TxtTidyFilesTitleResults)
	q.AddParagraph(results)
	q.DisplayWithActions()

}

func processFileTypes(p *term.Page, fileExtension string) {
	startTime := time.Now()

	if !debugMode {
		p.Info(lang.TxtRemovingFilesWithExt + fileExtension)
		removeFiles(p, fileExtension)
	} else {
		p.Info(lang.TxtFindingFilesWithExt + fileExtension)
		findFiles(p, fileExtension)
	}
	endTime := time.Now()
	runtime := endTime.Sub(startTime)

	msg := fmt.Sprintf(lang.TxtOperationComplete, fileExtension, runtime.String())
	p.Success(msg)
	resultsAdd(msg)
}

func realpath(p *term.Page, path string) string {

	realPathCmd := exec.Command("realpath", path)
	output, err := realPathCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to resolve path: %v", PFY, err))
		p.Error(errs.ErrUnableToResolvePath, err.Error())
		return ""
	}
	return strings.TrimSpace(string(output))
}

// The function "getDiskInfo" returns the total disk size, free disk space, and percentage of disk
// space used for a given path.
func getDiskInfo(p *term.Page, path string) (total, free, percentUsed string) {
	info := du.NewDiskUsage(path)
	total = p.ViewPort().Formatters.HumanDiskSize(info.Size())
	free = p.ViewPort().Formatters.HumanDiskSize(info.Available())
	percentUsed = p.ViewPort().Formatters.Human(info.Usage())
	return total, free, percentUsed
}

func removeFiles(p *term.Page, fileExtension string) {
	t := p.ViewPort()
	if debugMode {
		//p.ViewPort()
		t.Print(lang.TxtTidyFilesWouldHaveRemoved)
		return
	}
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExtension, "-exec", "rm", "-f", "{}", ";")
	err := findCmd.Run()
	if err != nil {
		p.Error(errs.ErrUnableToFindFiles, err.Error())
		return
	}
	msg := fmt.Sprintf(lang.TxtCommandRun, findCmd.String())
	p.Success(msg)
	resultsAdd(msg)
	output, _ := findCmd.Output()
	results = append(results, string(output))
}

func findFiles(p *term.Page, fileExt string) {
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExt)

	output, err := findCmd.Output()
	if err != nil {
		p.Error(errs.ErrUnableToFindFiles, err.Error())
		return
	}
	p.Success(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
	resultsAdd(string(output))
}

func removeEmptyDirectories(p *term.Page) {
	t := p.ViewPort()
	if debugMode {
		p.Info(lang.TxtTidyFilesWouldHaveRemoved)
		return
	}
	findCmd := exec.Command("find", ".", "-type", "d", "-exec", "rmdir", "{}", "+")
	err := findCmd.Run()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to remove empty directories: %v", PFY, err))
		p.Error(errs.ErrUnableToRemoveDirectories, err.Error())
		return
	}
	t.Println(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
	resultsAdd(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
	output, _ := findCmd.Output()
	outputTxt := string(output)
	resultsAdd(outputTxt)
}

func findEmptyDirectories(p *term.Page) {
	//t := p.ViewPort()
	findCmd := exec.Command("find", ".", "-type", "d", "-empty", "-print")
	output, err := findCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to find empty directories: %v", PFY, err))
		p.Error(errs.ErrNoEmptyDirectories, err.Error())
		return
	}
	resultsAdd(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
	resultsAdd(string(output))
}

func printStorageReport(p *term.Page, beforeDiskSizeTotal, beforeDiskSizeFree, beforeDiskPercentUsed, afterDiskSizeTotal, afterDiskSizeFree, afterDiskPercentUsed string) {
	t := p.ViewPort()
	mode := lang.TxtDebugMode
	if !debugMode {
		mode = lang.TxtLiveMode
	}

	resultsAdd(t.Formatters.Bold(t.Underline(lang.TxtStorageReportTitle)))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesBefore, t.Formatters.Bold(beforeDiskSizeFree), t.Formatters.Bold(beforeDiskSizeTotal), t.Formatters.Bold(beforeDiskPercentUsed)))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesAfter, t.Formatters.Bold(afterDiskSizeFree), t.Formatters.Bold(afterDiskSizeTotal), t.Formatters.Bold(afterDiskPercentUsed)))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesMachine, t.Helpers.GetSytemInfo()))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesHost, t.Helpers.GetHostName()))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesUser, t.Helpers.GetUsername()))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesMode, mode))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesTypes, strings.Join(fileExtensions, " ")))
	resultsAdd(fmt.Sprintf(lang.TxtTidyFilesEnd, time.Now().Format(cfg.TimeStampFormat)))
}

func resultsAdd(in string) {
	results = append(results, in)
}

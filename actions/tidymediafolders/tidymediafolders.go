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
		p.AddFieldValuePair("Mode", "Debug")
		//t.Special(t.Formatters.Bold(t.Underline(lang.TxtLiveRun)))
	} else {
		p.AddFieldValuePair("Mode", "Normal")
		//t.Print(t.Underline(lang.TxtTrailRun))
	}

	//t.Print(lang.TxtResolvedPath + realpath(t, pathIn))
	p.AddFieldValuePair("Path", realpath(p, pathIn))
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

	p.Success(fmt.Sprintf(lang.TxtOperationComplete, fileExtension, runtime.String()))

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
	total = t.Formatters.HumanDiskSize(info.Size())
	free = t.Formatters.HumanDiskSize(info.Available())
	percentUsed = t.Formatters.Human(info.Usage())
	return total, free, percentUsed
}

func removeFiles(p *term.Page, fileExtension string) {
	if debugMode {
		t.Print(lang.TxtTidyFilesWouldHaveRemoved)
		return
	}
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExtension, "-exec", "rm", "-f", "{}", ";")
	err := findCmd.Run()
	if err != nil {
		t.Error(errs.ErrUnableToFindFiles, err.Error())
		return
	}
	t.Println(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
}

func findFiles(p *term.Page, fileExt string) {
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExt)

	output, err := findCmd.Output()
	if err != nil {
		t.Error(errs.ErrUnableToFindFiles, err.Error())
		return
	}
	t.Println(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
	t.Spool(output)
}

func removeEmptyDirectories(p *term.Page) {
	if debugMode {
		t.Print(lang.TxtTidyFilesWouldHaveRemoved)
		return
	}
	findCmd := exec.Command("find", ".", "-type", "d", "-exec", "rmdir", "{}", "+")
	err := findCmd.Run()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to remove empty directories: %v", PFY, err))
		t.Error(errs.ErrUnableToRemoveDirectories, err.Error())
		return
	}
	t.Println(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
}

func findEmptyDirectories(p *term.Page) {
	findCmd := exec.Command("find", ".", "-type", "d", "-empty", "-print")
	output, err := findCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to find empty directories: %v", PFY, err))
		t.Error(errs.ErrNoEmptyDirectories, err.Error())
		return
	}
	t.Println(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
	t.Spool(output)
}

func printStorageReport(p *term.Page, beforeDiskSizeTotal, beforeDiskSizeFree, beforeDiskPercentUsed, afterDiskSizeTotal, afterDiskSizeFree, afterDiskPercentUsed string) {
	t := p.ViewPort()
	mode := lang.TxtDebugMode
	if !debugMode {
		mode = lang.TxtLiveMode
	}

	t.Break()
	t.Print(t.Formatters.Bold(t.Underline(lang.TxtStorageReportTitle)))
	t.Break()
	t.Print(fmt.Sprintf(lang.TxtTidyFilesBefore, t.Formatters.Bold(beforeDiskSizeFree), t.Formatters.Bold(beforeDiskSizeTotal), t.Formatters.Bold(beforeDiskPercentUsed)))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesAfter, t.Formatters.Bold(afterDiskSizeFree), t.Formatters.Bold(afterDiskSizeTotal), t.Formatters.Bold(afterDiskPercentUsed)))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesMachine, t.Helpers.GetSytemInfo()))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesHost, t.Helpers.GetHostName()))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesUser, t.Helpers.GetUsername()))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesMode, mode))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesTypes, strings.Join(fileExtensions, " ")))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesEnd, time.Now().Format(cfg.TimeStampFormat)))
}

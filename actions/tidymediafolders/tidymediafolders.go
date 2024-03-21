package tidymediafolders

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	term "github.com/mt1976/crt"
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

//var crt support.Crt

func Run(t term.Crt, debugModeIn bool, pathIn string) {
	debugMode = debugModeIn

	t.Print(lang.TxtTidyFilesTitle)

	if len(pathIn) < 1 {
		t.Error(errs.ErrNoPathSpecified, pathIn)
		return
	}

	//path := os.Args[1]

	if _, err := os.Stat(pathIn); os.IsNotExist(err) {
		t.Error(errs.ErrInvalidPath, err.Error())
		return
	}

	if pathIn == "/" || pathIn == "~" {
		t.Error(errs.ErrInvalidPathSpecialDirectory, pathIn)
		return
	}

	if !debugMode {
		t.Special(t.Formatters.Bold(t.Underline(lang.TxtLiveRun)))
	} else {
		t.Print(t.Underline(lang.TxtTrailRun))
	}

	t.Print(lang.TxtResolvedPath + realpath(t, pathIn))
	//fmt.Printf("%s File types to be removed: [%s]\n", support.CHnormal, strings.Join(types, " "))
	t.Print(fmt.Sprintf(lang.TxtTidyFilesStart, t.Formatters.Bold(strings.Join(fileExtensions, " "))))
	//var userResponse string
	//fmt.Printf("%s Are you sure you want to proceed? %s(y/n) : %s", PFY, bold, normal)
	userResponse := t.Input(lang.TxtAreYouSureYouWantToProceed, lang.OptAreYouSureYouWantToProceed)
	//fmt.Scanln(&userResponse)

	if strings.ToLower(userResponse) != "y" && strings.ToLower(userResponse) != "yes" {
		//fmt.Printf("%s Exiting\n", PFY)
		t.Print(t.Formatters.Bold(lang.TxtQuittingMessage))
		return
	}
	t.Blank()
	diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore := getDiskInfo(t, pathIn)

	//fmt.Printf("%s Changing directory to %s\n", PFY, path)
	t.Print(fmt.Sprintf(lang.TxtChangingDirectory, t.Formatters.Bold(pathIn)))
	t.Blank()
	err := os.Chdir(pathIn)
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to change directory: %v", PFY, err))
		t.Error(errs.ErrFailedToChangeDirectory, pathIn, err.Error())
		return
	}

	var wg sync.WaitGroup

	for idx, fileExtension := range fileExtensions {
		wg.Add(idx)
		//fmt.Printf("%s Operation on .%s files completed in %s seconds\n", support.CHspecial, fileExt, runtime)

		defer wg.Done()
		go processFileTypes(t, fileExtension)

	}
	wg.Wait()

	t.Special(lang.TxtTidyFilesDeletingDirectories)
	startLoopIteration := time.Now()
	if !debugMode {
		removeEmptyDirectories(t)
	} else {
		findEmptyDirectories(t)
	}
	endLoopIteration := time.Now()
	runtime := endLoopIteration.Sub(startLoopIteration)
	t.Special(fmt.Sprintf(lang.TxtTidyFilesDeletingDirectoriesCompleted, t.Formatters.Bold(runtime.String())))
	diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter := getDiskInfo(t, pathIn)

	printStorageReport(t, diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore, diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter)
}

func processFileTypes(t term.Crt, fileExtension string) {
	startTime := time.Now()

	if !debugMode {
		t.Special(lang.TxtRemovingFilesWithExt + t.Formatters.Bold(fileExtension))
		removeFiles(t, fileExtension)
	} else {
		t.Special(lang.TxtFindingFilesWithExt + t.Formatters.Bold(fileExtension))
		findFiles(t, fileExtension)
	}
	endTime := time.Now()
	runtime := endTime.Sub(startTime)

	t.Special(fmt.Sprintf(lang.TxtOperationComplete, t.Formatters.Bold(fileExtension), t.Formatters.Bold(runtime.String())))

}

func realpath(t term.Crt, path string) string {

	realPathCmd := exec.Command("realpath", path)
	output, err := realPathCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to resolve path: %v", PFY, err))
		t.Error(errs.ErrUnableToResolvePath, err.Error())
		return ""
	}
	return strings.TrimSpace(string(output))
}

// The function "getDiskInfo" returns the total disk size, free disk space, and percentage of disk
// space used for a given path.
func getDiskInfo(t term.Crt, path string) (total, free, percentUsed string) {
	info := du.NewDiskUsage(path)
	total = t.Formatters.HumanDiskSize(info.Size())
	free = t.Formatters.HumanDiskSize(info.Available())
	percentUsed = t.Formatters.Human(info.Usage())
	return total, free, percentUsed
}

func removeFiles(t term.Crt, fileExtension string) {
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

func findFiles(t term.Crt, fileExt string) {
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExt)

	output, err := findCmd.Output()
	if err != nil {
		t.Error(errs.ErrUnableToFindFiles, err.Error())
		return
	}
	t.Println(fmt.Sprintf(lang.TxtCommandRun, findCmd.String()))
	t.Spool(output)
}

func removeEmptyDirectories(t term.Crt) {
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

func findEmptyDirectories(t term.Crt) {
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

func printStorageReport(t term.Crt, beforeDiskSizeTotal, beforeDiskSizeFree, beforeDiskPercentUsed, afterDiskSizeTotal, afterDiskSizeFree, afterDiskPercentUsed string) {

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

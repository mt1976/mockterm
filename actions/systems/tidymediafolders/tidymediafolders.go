package tidymediafolders

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	file "github.com/mt1976/crt/filechooser"
	page "github.com/mt1976/crt/page"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/actions/systems/tidymediafolders/lang"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	clng "github.com/mt1976/mockterm/language"
	mode "github.com/mt1976/mockterm/support/modes"
	"github.com/ricochet2200/go-disk-usage/du"
)

var (
	fileExtensions = []string{"nfo", "jpeg", "jpg", "bif", "vob", "txt", "png", "me", "exe"}
)

var cfg = conf.Configuration
var results = []string{}

func Run(t *term.ViewPort, m mode.Modality, pathIn string) {

	//new page
	p := page.NewPage(t, lang.Title.Text())
	resultsAdd(upcase(p, lang.CleanFileNamesReport.Text()))

	if pathIn == "" {
		home, err := file.UserHome()
		if err != nil {
			p.Error(err)
		}
		pathIn = home
	}

	path, err := file.ChooseDirectory(pathIn)
	if err != nil {
		p.Error(errs.ErrTidyFiles, err.Error())
	}
	pathIn = path

	if len(pathIn) < 1 {
		p.Error(errs.ErrNoPathSpecified, pathIn)
		return
	}

	if _, err := os.Stat(pathIn); os.IsNotExist(err) {
		p.Error(errs.ErrInvalidPath, err.Error())
		return
	}

	if pathIn == "/" || pathIn == "~" {
		p.Error(errs.ErrInvalidPathSpecialDirectory, pathIn)
		return
	}

	if m.IsDebug() {
		p.AddFieldValuePair(clng.Mode.Text(), lang.DebugMode.Text())
	} else {
		p.AddFieldValuePair(clng.Mode.Text(), lang.LiveMode.Text())
	}

	p.AddFieldValuePair(clng.Path.Text(), realpath(p, pathIn))
	p.AddFieldValuePair(lang.FileTypesTxt, strings.Join(fileExtensions, " "))
	resultsAdd(clng.Path.Text() + " : " + pathIn)
	resultsAdd(lang.FileTypesTxt.Text() + " : " + strings.Join(fileExtensions, " "))

	p.AddBlankRow()
	msg := fmt.Sprintf(lang.Start.Text(), t.Formatters.Bold(strings.Join(fileExtensions, " ")))
	p.Add(msg, "", "")
	resultsAdd(msg)

	//var userResponse string
	//fmt.Printf("%s Are you sure you want to proceed? %s(y/n) : %s", PFY, bold, normal)

	ok, err := p.Display_Confirmation(lang.AreYouSureYouWantToProceed)
	if err != nil {
		p.Error(errs.ErrTidyFiles, err.Error())
	}

	if !ok {
		p.Info(clng.Quitting, "", "")
		return
	}

	diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore := getDiskInfo(p, pathIn)

	p.Info(lang.ChangingDirectory, t.Formatters.Bold(pathIn))

	err = os.Chdir(pathIn)
	if err != nil {
		p.Error(errs.ErrFailedToChangeDirectory, pathIn, err.Error())
		return
	}

	for _, fileExtension := range fileExtensions {
		processFileTypes(p, m, fileExtension)
	}

	p.Info(lang.DeletingDirectories)
	startLoopIteration := time.Now()
	if m.IsLive() {
		removeEmptyDirectories(p, m)
	} else {
		findEmptyDirectories(p, m)
	}
	endLoopIteration := time.Now()
	runtime := endLoopIteration.Sub(startLoopIteration)
	p.Success(lang.DeletingDirectoriesCompleted, t.Formatters.Bold(runtime.String()))
	diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter := getDiskInfo(p, pathIn)

	printStorageReport(p, m, diskSizeTotalBefore, diskSizeFreeBefore, diskPercentUsedBefore, diskSizeTotalAfter, diskSizeFreeAfter, diskPercentUsedAfter)
	q := page.NewPage(t, lang.TitleResults.Text())
	q.AddParagraph(results)
	q.Display_Actions()

}

func processFileTypes(p *page.Page, m mode.Modality, fileExtension string) {
	msg1 := fmt.Sprintf(lang.Processing.Text(), fileExtension)
	p.Info(lang.Processing, fileExtension)
	resultsAdd(msg1)

	startTime := time.Now()

	if m.IsLive() {
		p.Info(lang.RemovingFilesWithExt, fileExtension)
		removeFiles(p, m, fileExtension)
	} else {
		p.Info(lang.FindingFilesWithExt, fileExtension)
		findFiles(p, m, fileExtension)
	}
	endTime := time.Now()
	runtime := endTime.Sub(startTime)

	msg := fmt.Sprintf(lang.OperationComplete.Text(), fileExtension, runtime.String())
	p.Success(lang.OperationComplete, fileExtension, runtime.String())
	resultsAdd(msg)
}

func realpath(p *page.Page, path string) string {

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
func getDiskInfo(p *page.Page, path string) (total, free, percentUsed string) {
	info := du.NewDiskUsage(path)
	total = p.ViewPort().Formatters.HumanDiskSize(info.Size())
	free = p.ViewPort().Formatters.HumanDiskSize(info.Available())
	percentUsed = p.ViewPort().Formatters.Human(info.Usage())
	return total, free, percentUsed
}

func removeFiles(p *page.Page, m mode.Modality, fileExtension string) {
	msg1 := fmt.Sprintf(lang.RemovingFilesWithExt.Text(), fileExtension)
	p.Info(lang.RemovingFilesWithExt, fileExtension)
	resultsAdd(msg1)
	t := p.ViewPort()
	if m.IsDebug() {
		//p.ViewPort()
		t.Print(lang.FilesWouldHaveRemoved.Text())
		return
	}
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExtension, "-exec", "rm", "-f", "{}", ";")
	err := findCmd.Run()
	if err != nil {
		p.Error(errs.ErrUnableToFindFiles, err.Error())
		return
	}
	msg := fmt.Sprintf(lang.CommandRun.Text(), findCmd.String())
	p.Success(lang.CommandRun, findCmd.String())
	resultsAdd(msg)
	output, _ := findCmd.Output()
	results = append(results, string(output))
	msg2 := fmt.Sprintf(lang.OperationComplete.Text(), fileExtension, time.Now().Format(cfg.TimeStampFormat))
	p.Success(lang.OperationComplete, fileExtension, time.Now().Format(cfg.TimeStampFormat))
	resultsAdd(msg2)
}

func findFiles(p *page.Page, m mode.Modality, fileExt string) {
	msg1 := fmt.Sprintf(lang.FindingFilesWithExt.Text(), fileExt)
	p.Info(lang.FindingFilesWithExt, fileExt)
	resultsAdd(msg1)
	findCmd := exec.Command("find", ".", "-type", "f", "-name", "*."+fileExt)

	output, err := findCmd.Output()
	if err != nil {
		p.Error(errs.ErrUnableToFindFiles, err.Error())
		return
	}
	p.Success(lang.CommandRun, findCmd.String())
	resultsAdd(string(output))
	msg2 := fmt.Sprintf(lang.OperationComplete.Text(), fileExt, time.Now().Format(cfg.TimeStampFormat))
	p.Success(lang.OperationComplete, fileExt, time.Now().Format(cfg.TimeStampFormat))
	resultsAdd(msg2)
}

func removeEmptyDirectories(p *page.Page, m mode.Modality) error {
	//msg1 := lang.TxtRemovingEmptyDirectories
	p.Info(lang.RemovingEmptyDirectories)
	resultsAdd(lang.RemovingEmptyDirectories.Text())
	t := p.ViewPort()
	if m.IsDebug() {
		p.Info(lang.FilesWouldHaveRemoved)
		return nil
	}
	findCmd := exec.Command("find", ".", "-type", "d", "-exec", "rmdir", "{}", "+")
	err := findCmd.Run()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to remove empty directories: %v", PFY, err))
		p.Error(errs.ErrUnableToRemoveDirectories, err.Error())
		return err
	}
	t.Println(fmt.Sprintf(lang.CommandRun.Text(), findCmd.String()))
	resultsAdd(fmt.Sprintf(lang.CommandRun.Text(), findCmd.String()))
	output, _ := findCmd.Output()
	outputTxt := string(output)
	resultsAdd(outputTxt)
	msg2 := fmt.Sprintf(lang.OperationCompleteIncomplete.Text(), time.Now().Format(cfg.TimeStampFormat))
	p.Success(lang.OperationCompleteIncomplete, time.Now().Format(cfg.TimeStampFormat))
	resultsAdd(msg2)
	return nil
}

func findEmptyDirectories(p *page.Page, m mode.Modality) error {

	p.Info(lang.FindingEmptyDirectories)
	resultsAdd(lang.FindingEmptyDirectories.Text())
	//t := p.ViewPort()
	findCmd := exec.Command("find", ".", "-type", "d", "-empty", "-print")
	output, err := findCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to find empty directories: %v", PFY, err))
		p.Error(errs.ErrNoEmptyDirectories, err.Error())
		return errs.ErrNoEmptyDirectories
	}
	resultsAdd(fmt.Sprintf(lang.CommandRun.Text(), findCmd.String()))
	resultsAdd(string(output))
	p.Success(lang.OperationCompleteIncomplete, time.Now().Format(cfg.TimeStampFormat))
	resultsAdd(fmt.Sprintf(lang.OperationCompleteIncomplete.Text(), time.Now().Format(cfg.TimeStampFormat)))
	return nil
}

func printStorageReport(p *page.Page, m mode.Modality, beforeDiskSizeTotal, beforeDiskSizeFree, beforeDiskPercentUsed, afterDiskSizeTotal, afterDiskSizeFree, afterDiskPercentUsed string) {
	t := p.ViewPort()
	mode := lang.DebugMode
	if m.IsLive() {
		mode = lang.LiveMode
	}

	resultsAdd(t.Formatters.Bold(t.Underline(lang.StorageReport.Text())))
	resultsAdd("")
	resultsAdd(fmt.Sprintf(lang.FilesBefore.Text(), t.Formatters.Bold(beforeDiskSizeFree), t.Formatters.Bold(beforeDiskSizeTotal), t.Formatters.Bold(beforeDiskPercentUsed)))
	resultsAdd(fmt.Sprintf(lang.FilesAfter.Text(), t.Formatters.Bold(afterDiskSizeFree), t.Formatters.Bold(afterDiskSizeTotal), t.Formatters.Bold(afterDiskPercentUsed)))
	resultsAdd(fmt.Sprintf(lang.Machine.Text(), t.Helpers.GetSytemInfo()))
	resultsAdd(fmt.Sprintf(lang.Host.Text(), t.Helpers.GetHostName()))
	un, _ := t.Helpers.GetUserName()
	resultsAdd(fmt.Sprintf(lang.User.Text(), un))
	resultsAdd(fmt.Sprintf(lang.Mode.Text(), mode))
	resultsAdd(fmt.Sprintf(lang.FileTypes.Text(), strings.Join(fileExtensions, " ")))
	resultsAdd(fmt.Sprintf(lang.FilesEnd.Text(), time.Now().Format(cfg.TimeStampFormat)))
}

func resultsAdd(in string) {
	results = append(results, in)
}

func upcase(t *page.Page, msg string) string {
	return t.ViewPort().Formatters.Upcase(msg)
}

package disksize

import (
	"fmt"
	"os"

	term "github.com/mt1976/crt"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

func Run(t term.ViewPort, debug bool, args []string) {
	if len(args) < 2 {
		t.Shout(errs.ErrNoFilesSpecified)
		t.Print(errs.ErrDiskSizeUsage)
		return
	}
	t.Special(fmt.Sprintf(lang.TxtFileSizesOfNFiles, len(args)-1))
	t.Break()
	for _, file := range args[1:] {
		printFileSize(t, file)
	}
}

func printFileSize(t term.ViewPort, path string) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		//msg := fmt.Sprintf("File %s does not exist", path)
		t.Error(errs.ErrFileDoesNotExist, path)
		return
	}
	if err != nil {
		//fmt.Print("Something went wrong\n")
		t.Error(errs.ErrFileInfo, err.Error())
		return
	}
	size := fi.Size()
	sizeKb := float64(size) / 1024
	sizeMb := sizeKb / 1024
	sizeGb := sizeMb / 1024
	sizeTb := sizeGb / 1024
	fileName := t.Formatters.Bold(fi.Name())
	duMsg := fmt.Sprintf(lang.DUOutputConstructor,
		size,
		sizeKb,
		sizeMb,
		sizeGb,
		sizeTb,
		fileName,
	)
	t.Print(duMsg)
}

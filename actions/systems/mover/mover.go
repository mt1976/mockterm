package mover

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	gomv "github.com/draxil/gomv"
	f "github.com/mt1976/crt/filechooser"
	page "github.com/mt1976/crt/page"
	symb "github.com/mt1976/crt/strings/symbols"
	term "github.com/mt1976/crt/terminal"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	mode "github.com/mt1976/mockterm/support/modes"
)

func Run(t *term.ViewPort, m mode.Modality) error {

	p := page.NewPage(t, lang.FileMigratorTitle.Text())
	p.AddBlankRow()
	if m.IsLive() {
		p.AddFieldValuePair(lang.FileMigratorMode, lang.TxtLiveMode)
	} else {
		p.AddFieldValuePair(lang.FileMigratorMode, lang.TxtDebugMode)
	}
	//p.AddParagraph([]string{"This is a test", "This is a test"})
	//p.AddBlankRow()
	proceed, err := p.Display_Confirmation(lang.FileMigratorModeCheckPrompt.Text())
	if err != nil {
		return err
	}
	if !proceed {
		p.Info(lang.Quitting.Text())
		return nil
	}
	fileName, isDir, err := f.FileChooser(".", f.FilesAll)
	if err != nil {
		return err
	}
	if isDir {
		p.Error(errs.ErrFileMigratorDirectory, fileName)
		return errs.ErrFileMigratorDirectory
	}
	p.AddFieldValuePair(lang.FileMigratorFile, fileName)
	p.Display_Confirmation(lang.FileMigratorModeCheckPrompt.Text() + fileName)

	// load the file from the os
	// open the file and read the contents

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		p.Error(err, fileName)
		//fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close() // Make sure to close the file when the function returns

	// read the file into a string
	data, err := io.ReadAll(file)
	if err != nil {
		p.Error(err, fileName)
		return err
	}

	lines := strings.Split(string(data), symb.Newline.Symbol())
	p.AddFieldValuePair(lang.FileMigratorNoFilesToProcess, strconv.Itoa(len(lines)))

	move, err := p.Display_Confirmation(lang.FileMigratorModeCheckPrompt.Text())
	if err != nil {
		return err
	}
	if !move {
		p.Info(lang.Quitting.Text())
		return nil
	}

	toFolder, isDir, err := f.FileChooser(".", f.DirectoriesOnly)
	if err != nil {
		return err
	}
	if !isDir {
		p.Error(errs.ErrFileMigratorFile, toFolder)
		return errs.ErrFileMigratorFile
	}
	p.AddFieldValuePair(lang.FileMigratorDestination, toFolder)
	p.Display_Confirmation(lang.FileMigratorModeCheckPrompt.Text())
	p.AddBlankRow()
	p.Add(lang.FileMigratorResults.Text(), "", "")
	p.AddBreakRow()
	// loop through each line in the data and read the line
	for z, from := range lines {
		err := moveFile(p, m, from, toFolder, z, len(lines))
		if err != nil {
			p.Error(err, from)
			return err
		}
	}
	p.Display_Confirmation(lang.FileMigratorDonePrompt.Text())

	// write the string to the console
	//page.Dump(string(data))

	return nil

}

func lastNChars(s string, n int) string {
	// Ensure we do not exceed the string's length
	if n > utf8.RuneCountInString(s) {
		return s // Return the entire string if n is greater than the string length
	}
	// Count backwards n runes
	start := len(s)
	for n > 0 {
		_, size := utf8.DecodeLastRuneInString(s[:start])
		start -= size
		n--
	}
	return s[start:]
}

func moveFile(page *page.Page, m mode.Modality, from, to string, pageNo, ofPages int) error {
	from20Chars := lastNChars(from, 20)
	to20Chars := lastNChars(to, 20)

	sep := string(os.PathSeparator)
	//split from by sep

	fromParts := strings.Split(from, sep)
	lastPartPos := len(fromParts)
	lastPart := fromParts[lastPartPos-1]
	destination := to + sep + lastPart

	progressmsg := fmt.Sprintf(lang.FileMigratorMoving.Text(), pageNo+1, ofPages, from20Chars, to20Chars)

	page.Info(progressmsg)
	time.Sleep(2 * time.Second)

	msg := from + lang.FileMigratorMovingArrow.Text() + destination
	page.Add(msg, "", "")

	if m.IsTrialMode() {
		//page.Info("Would have moved", dquote(from20Chars), dquote(destination))
		// fmt.Println("Would have moved", dquote(from), dquote(destination))
		// fmt.Println("from", from)
		// fmt.Println("to", to)
		// fmt.Println("fromParts", fromParts)
		// fmt.Println("lastPartPos", lastPartPos)
		// fmt.Println("lastPart", lastPart)
		// fmt.Println("destination", destination)
		// fmt.Println("sep", sep)
		//os.Exit(0)

		return nil
	}
	err := gomv.MoveFile(from, destination)
	if err != nil {
		return err
	}

	return nil
}

// func dquote(s string) string {
// 	return "\"" + s + "\""
// }

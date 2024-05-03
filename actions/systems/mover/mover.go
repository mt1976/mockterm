package mover

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	gomv "github.com/draxil/gomv"
	term "github.com/mt1976/crt"
	f "github.com/mt1976/crt/filechooser"
)

var mode bool
var TRIAL_MODE = true
var LIVE_MODE = false

func Run(t *term.ViewPort, inMode bool) error {

	mode = !inMode
	p := t.NewPage("Move Files")
	p.AddBlankRow()
	if mode == LIVE_MODE {
		p.AddFieldValuePair("Mode", "LIVE")
	} else {
		p.AddFieldValuePair("Mode", "Trial")
	}
	//p.AddParagraph([]string{"This is a test", "This is a test"})
	//p.AddBlankRow()
	proceed, err := p.Display_Confirmation("Are you sure you want to proceed?")
	if err != nil {
		return err
	}
	if !proceed {
		p.Info("Quitting")
		return nil
	}
	fileName, isDir, err := f.FileChooser(".", f.FilesOnly)
	if err != nil {
		return err
	}
	if isDir {
		p.Error(errors.New("this is a directory"), fileName)
	} else {
		p.Info("This is a file")
	}
	p.AddFieldValuePair("File", fileName)
	p.Display_Confirmation("Are you sure you want to process " + fileName)

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

	lines := strings.Split(string(data), "\n")
	p.AddFieldValuePair("No Items", strconv.Itoa(len(lines)))

	move, err := p.Display_Confirmation("Are you sure you want to proceed")
	if err != nil {
		return err
	}
	if !move {
		p.Info("Quitting")
		return nil
	}

	toFolder, isDir, err := f.FileChooser(".", f.DirectoriesOnly)
	if err != nil {
		return err
	}
	if !isDir {
		p.Error(errors.New("this is a file"), toFolder)
	} else {
		p.Info("This is a directory", toFolder)
	}
	p.AddFieldValuePair("Destination", toFolder)
	p.Display_Confirmation("Are you sure you want to continue processing")
	p.Add("RESULTS", "", "")
	p.AddBreakRow()
	// loop through each line in the data and read the line
	for _, line := range lines {
		err := moveFile(p, line, toFolder)
		if err != nil {
			p.Error(err, line)
			return err
		}
	}
	p.Display_Confirmation("Processing Complete")

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

func moveFile(page *term.Page, from string, to string) error {
	mode = true
	from20Chars := lastNChars(from, 20)
	to20Chars := lastNChars(to, 20)

	sep := "\\"
	//split from by sep

	fromParts := strings.Split(from, sep)
	lastPartPos := len(fromParts)
	lastPart := fromParts[lastPartPos-1]
	destination := to + sep + lastPart

	page.Info("Moving", dquote(from20Chars), dquote(to20Chars))
	time.Sleep(2 * time.Second)

	msg := from + " -> " + destination
	page.Add(msg, "", "")

	if mode == TRIAL_MODE {
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

func dquote(s string) string {
	return "\"" + s + "\""
}

package mover

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	term "github.com/mt1976/crt"
	f "github.com/mt1976/crt/filechooser"
)

func Run(t *term.ViewPort, liveMode bool) error {
	p := t.NewPage("Move Files")
	p.AddBlankRow()
	if liveMode {
		p.AddFieldValuePair("Mode", "LIVE")
	} else {
		p.AddFieldValuePair("Mode", "Trial")
	}
	p.AddParagraph([]string{
		"This is a test", "This is a test"})
	p.AddBlankRow()
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

	// loop through each line in the data and read the line
	for _, line := range lines {
		err := moveFile(p, line, toFolder)
		if err != nil {
			p.Error(err, line)
			return err
		}
	}

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
	from20Chars := lastNChars(from, 20)
	to20Chars := lastNChars(to, 20)

	page.Info("Moving", dquote(from20Chars), dquote(to20Chars))
	time.Sleep(2 * time.Second)
	return nil
}

func dquote(s string) string {
	return "\"" + s + "\""
}

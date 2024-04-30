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
	files "github.com/mt1976/crt/filechooser"
)

func Run(t *term.ViewPort, liveMode bool) error {
	page := t.NewPage("Move Files")
	page.AddBlankRow()
	if liveMode {
		page.AddFieldValuePair("Mode", "LIVE")
	} else {
		page.AddFieldValuePair("Mode", "Trial")
	}
	page.AddParagraph([]string{
		"This is a test", "This is a test"})
	page.AddBlankRow()
	proceed, err := page.Display_Confirmation("Are you sure you want to proceed?")
	if err != nil {
		return err
	}
	if !proceed {
		page.Info("Quitting")
		return nil
	}
	fileName, isDir, err := files.FileChooser(".", files.FilesOnly)
	if err != nil {
		return err
	}
	if isDir {
		page.Error(errors.New("this is a directory"), fileName)
	} else {
		page.Info("This is a file")
	}
	page.AddFieldValuePair("File", fileName)
	page.Display_Confirmation("Are you sure you want to process " + fileName)

	// load the file from the os
	// open the file and read the contents

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		page.Error(err, fileName)
		//fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close() // Make sure to close the file when the function returns

	// read the file into a string
	data, err := io.ReadAll(file)
	if err != nil {
		page.Error(err, fileName)
		return err
	}

	lines := strings.Split(string(data), "\n")
	page.AddFieldValuePair("No Items", strconv.Itoa(len(lines)))

	move, err := page.Display_Confirmation("Are you sure you want to proceed")
	if err != nil {
		return err
	}
	if !move {
		page.Info("Quitting")
		return nil
	}

	toFolder, isDir, err := files.FileChooser(".", files.DirectoriesOnly)
	if err != nil {
		return err
	}
	if !isDir {
		page.Error(errors.New("this is a file"), toFolder)
	} else {
		page.Info("This is a directory", toFolder)
	}
	page.AddFieldValuePair("Destination", toFolder)
	page.Display_Confirmation("Are you sure you want to continue processing")

	// loop through each line in the data and read the line
	for _, line := range lines {
		err := moveFile(page, line, toFolder)
		if err != nil {
			page.Error(err, line)
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

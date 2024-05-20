package language

import lang "github.com/mt1976/crt/language"

// File Migrator
var (
	Title            *lang.Text = lang.New("File Migration")
	Mode             *lang.Text = lang.New("Mode")
	ModeCheckPrompt  *lang.Text = lang.New("Are you sure you want to proceed, select (Y) to continue...")
	File             *lang.Text = lang.New("File")
	NoFilesToProcess *lang.Text = lang.New("No Files")
	Destination      *lang.Text = lang.New("Destination")
	Results          *lang.Text = lang.New("RESULTS")
	DonePrompt       *lang.Text = lang.New("Processing Complete, select (Y) to continue...")
	Moving           *lang.Text = lang.New("Moving [%d/%d] [%20v][%20v]")
	Arrow            *lang.Text = lang.New(" -> ")
)

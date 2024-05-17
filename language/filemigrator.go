package language

import lang "github.com/mt1976/crt/language"

// File Migrator
var (
	FileMigratorTitle            *lang.Text = lang.New("File Migration")
	FileMigratorMode             *lang.Text = lang.New("Mode")
	FileMigratorModeCheckPrompt  *lang.Text = lang.New("Are you sure you want to proceed, select (Y) to continue...")
	FileMigratorFile             *lang.Text = lang.New("File")
	FileMigratorNoFilesToProcess *lang.Text = lang.New("No Files")
	FileMigratorDestination      *lang.Text = lang.New("Destination")
	FileMigratorResults          *lang.Text = lang.New("RESULTS")
	FileMigratorDonePrompt       *lang.Text = lang.New("Processing Complete, select (Y) to continue...")
	FileMigratorMoving           *lang.Text = lang.New("Moving [%d/%d] [%20v][%20v]")
	FileMigratorMovingArrow      *lang.Text = lang.New(" -> ")
)

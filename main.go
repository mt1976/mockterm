package main

import (
	"time"

	term "github.com/mt1976/crt"
	text "github.com/mt1976/crt/language"
	menu "github.com/mt1976/mockterm/actions/mainmenu"
	cnfg "github.com/mt1976/mockterm/config"
	strt "github.com/mt1976/mockterm/start"
)

// config is used to store configuration settings for the program, including terminal
// width and height.
//

// Main is the entry point for the program.
func main() {

	C := cnfg.Configuration

	// create a new instance of the Crt
	crt := term.NewWithSize(C.TerminalWidth, C.TerminalHeight)
	// set the terminal size
	//crt.SetTerminalSize(config.term_width, config.term_height)
	// start a timer
	start := time.Now()

	// run the startup sequence
	crt.SetDelayInSec(C.Delay)
	strt.Run(&crt)
	crt.ResetDelay()
	//godump.Dump(crt)
	//os.Exit(0)
	// run the main menu
	menu.Run(&crt)

	// stop the timer
	elapsed := time.Since(start)
	// output the elapsed time
	crt.Shout(crt.Helpers.Bold(text.TxtDone) + text.Space + elapsed.String())

}

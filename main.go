package main

import (
	"time"

	t "github.com/mt1976/crt/language"
	terminal "github.com/mt1976/crt/support"
	config "github.com/mt1976/crt/support/config"
	mainmenu "github.com/mt1976/mockterm/actions/mainmenu"
	startup "github.com/mt1976/mockterm/start"
)

// config is used to store configuration settings for the program, including terminal
// width and height.
//

// Main is the entry point for the program.
func main() {

	C := config.Configuration

	// create a new instance of the Crt
	crt := terminal.NewWithSize(C.TerminalWidth, C.TerminalHeight)
	// set the terminal size
	//crt.SetTerminalSize(config.term_width, config.term_height)

	// start a timer
	start := time.Now()

	// run the startup sequence
	crt.SetDelayInSec(C.Delay)
	startup.Run(&crt)
	crt.ResetDelay()
	//godump.Dump(crt)
	//os.Exit(0)
	// run the main menu
	mainmenu.Run(&crt)

	// stop the timer
	elapsed := time.Since(start)
	// output the elapsed time
	crt.Shout(crt.Bold(t.TxtDone) + t.Space + elapsed.String())

}

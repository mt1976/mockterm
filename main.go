package main

import (
	"time"

	"github.com/mt1976/crt/terminal"
	menu "github.com/mt1976/mockterm/actions/mainmenu"
	cfg "github.com/mt1976/mockterm/config"
	l "github.com/mt1976/mockterm/language"
	startup "github.com/mt1976/mockterm/start"
)

// config is used to store configuration settings for the program, including terminal
// width and height.
//

// Main is the entry point for the program.
func main() {

	C := cfg.Configuration

	// create a new instance of the ViewPort

	t := terminal.NewWithSize(C.TerminalWidth, C.TerminalHeight)

	// start a timer
	start := time.Now()

	// run the startup sequence

	t.SetDelayInSec(C.Delay)

	startup.Run(&t)

	t.ResetDelay()

	// run the main menu
	menu.Run(&t)

	// stop the timer
	elapsed := time.Since(start)

	// output the elapsed time
	t.Shout(t.Formatters.Bold(l.Done.Text()) + l.Space.Text() + elapsed.String())

}

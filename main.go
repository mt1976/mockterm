package main

import (
	"time"

	viewPort "github.com/mt1976/crt"
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

	// create a new instance of the ViewPort

	vp := viewPort.NewWithSize(C.TerminalWidth, C.TerminalHeight)

	// start a timer
	start := time.Now()

	// run the startup sequence

	vp.SetDelayInSec(C.Delay)

	strt.Run(&vp)

	vp.ResetDelay()

	// run the main menu
	menu.Run(&vp)

	// stop the timer
	elapsed := time.Since(start)

	// output the elapsed time
	vp.Shout(vp.Formatters.Bold(text.TxtDone) + text.Space + elapsed.String())

}

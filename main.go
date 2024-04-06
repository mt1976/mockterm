package main

import (
	"log"
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

	// create a new instance of the ViewPort
	log.Println("Creating new terminal...")
	t := term.NewWithSize(C.TerminalWidth, C.TerminalHeight)
	log.Println("Created new terminal...")
	// set the terminal size
	//t.SetTerminalSize(config.term_width, config.term_height)
	// start a timer
	start := time.Now()

	// run the startup sequence
	log.Println("Setting up...")
	t.SetDelayInSec(C.Delay)
	log.Println("Starting up...")
	strt.Run(&t)
	log.Printf("Startup complete in %v \n", time.Since(start))
	t.ResetDelay()

	//os.Exit(0)
	// run the main menu
	menu.Run(&t)

	// stop the timer
	elapsed := time.Since(start)
	// output the elapsed time
	t.Shout(t.Formatters.Bold(text.TxtDone) + text.Space + elapsed.String())

}

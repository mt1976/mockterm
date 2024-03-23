package start

import (
	"fmt"

	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
	conf "github.com/mt1976/mockterm/config"
)

// Run initializes the terminal and runs the main loop.

var C = conf.Configuration

// Run initializes the terminal and runs the main loop.
func Run(crt *term.ViewPort) {
	// Clear the terminal screen.
	crt.Clear()

	// Display the banner.
	crt.Banner(lang.TxtStarting)

	// Print a message.
	crt.Print(lang.TxtStartingTerminal + lang.SymNewline)

	// Print a message.
	crt.Print(lang.TxtSelfTesting + lang.SymNewline)
	oldDelay := crt.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	crt.SetDelayInSec(0.25)
	crt.Print(lang.TxtSelfTesting + lang.TxtComplete + lang.SymNewline)
	// Print the current date and time.
	crt.SetDelayInMs(oldDelay)
	crt.Print(lang.TxtCurrentDate + crt.Formatters.DateString() + lang.SymNewline)
	crt.Print(lang.TxtCurrentTime + crt.Formatters.TimeString() + lang.SymNewline)

	// Print a message.
	crt.Print(lang.TxtPleaseWait + lang.SymNewline)

	// Check if the terminal has a baud rate set.
	if !crt.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(lang.TxtBaudRate, crt.Baud())
		crt.Print(msg + lang.SymNewline)
	}

	// Print a message.
	crt.Print(lang.TxtConnecting + lang.SymNewline)

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(lang.TxtDialing, crt.Helpers.RandomIP(), crt.Helpers.RandomPort())
	crt.Print(msg + lang.SymNewline)
	if !C.Debug {
		crt.SetDelayInSec(crt.Helpers.RandomFloat(1, 5))
	}
	if crt.Helpers.CoinToss() && !C.Debug {
		crt.Print(lang.ErrorMessageConnectionFailed + lang.SymNewline)
		// Print a message with the IP address and port number.
		crt.ResetDelay()
		msg := fmt.Sprintf(lang.TxtDialing, crt.Helpers.RandomIP(), crt.Helpers.RandomPort())
		crt.Print(msg + lang.SymNewline)
		crt.SetDelayInSec(crt.Helpers.RandomFloat(1, 5))
	}

	// Print a message.
	crt.Print(lang.TxtConnected + lang.SymNewline)
	crt.SetDelayInMs(oldDelay)
}

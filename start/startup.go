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
func Run(t *term.ViewPort) {
	// Clear the terminal screen.
	t.Clear()

	// Display the banner.
	t.Banner(lang.TxtStarting)

	// Print a message.
	t.Print(lang.TxtStartingTerminal + lang.SymNewline)

	// Print a message.
	t.Print(lang.TxtSelfTesting + lang.SymNewline)
	oldDelay := t.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	t.SetDelayInSec(0.25)
	t.Print(lang.TxtSelfTesting + lang.TxtComplete + lang.SymNewline)
	// Print the current date and time.
	t.SetDelayInMs(oldDelay)
	t.Print(lang.TxtCurrentDate + t.Formatters.DateString() + lang.SymNewline)
	t.Print(lang.TxtCurrentTime + t.Formatters.TimeString() + lang.SymNewline)

	// Print a message.
	t.Print(lang.TxtPleaseWait + lang.SymNewline)

	// Check if the terminal has a baud rate set.
	if !t.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(lang.TxtBaudRate, t.Baud())
		t.Print(msg + lang.SymNewline)
	}

	// Print a message.
	t.Print(lang.TxtConnecting + lang.SymNewline)

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(lang.TxtDialing, t.Helpers.RandomIP(), t.Helpers.RandomPort())
	t.Print(msg + lang.SymNewline)
	if !C.Debug {
		t.SetDelayInSec(t.Helpers.RandomFloat(1, 5))
	}
	if t.Helpers.CoinToss() && !C.Debug {
		t.Print(lang.ErrorMessageConnectionFailed + lang.SymNewline)
		// Print a message with the IP address and port number.
		t.ResetDelay()
		msg := fmt.Sprintf(lang.TxtDialing, t.Helpers.RandomIP(), t.Helpers.RandomPort())
		t.Print(msg + lang.SymNewline)
		t.SetDelayInSec(t.Helpers.RandomFloat(1, 5))
	}

	// Print a message.
	t.Print(lang.TxtConnected + lang.SymNewline)
	t.SetDelayInMs(oldDelay)
}

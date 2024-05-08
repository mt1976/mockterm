package start

import (
	"fmt"

	term "github.com/mt1976/crt/terminal"
	conf "github.com/mt1976/mockterm/config"
	lang "github.com/mt1976/mockterm/language"
)

// Run initializes the terminal and runs the main loop.

var cfg = conf.Configuration

// Run initializes the terminal and runs the main loop.
func Run(t *term.ViewPort) {
	// Clear the terminal screen.
	t.Clear()

	// Display the banner.
	t.Banner("Initialise...")
	t.Wait()

	// Print a message.
	fmt.Println(lang.TxtStartingTerminal)
	t.Wait()

	// Print a message.
	fmt.Println(lang.TxtSelfTesting)
	t.Wait()
	oldDelay := t.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	t.SetDelayInSec(0.25)
	fmt.Println(lang.TxtSelfTesting + lang.TxtComplete)
	t.Wait()
	// Print the current date and time.
	t.SetDelayInMs(oldDelay)
	fmt.Println(lang.TxtCurrentDate + t.Formatters.DateString())
	t.Wait()
	fmt.Println(lang.TxtCurrentTime + t.Formatters.TimeString())
	t.Wait()

	// Print a message.
	fmt.Println(lang.TxtPleaseWait)
	t.Wait()

	// Check if the terminal has a baud rate set.
	if !t.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(lang.TxtBaudRate, t.Baud())
		fmt.Println(msg)
		t.Wait()
	}

	// Print a message.
	fmt.Println(lang.TxtConnecting)
	t.Wait()

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(lang.TxtDialing, t.Helpers.RandomIP(), t.Helpers.RandomPort())
	fmt.Println(msg)
	t.Wait()
	if !cfg.Debug {
		t.SetDelayInSec(t.Helpers.RandomFloat(1, 5))
	}
	if t.Helpers.CoinToss() && !cfg.Debug {
		fmt.Println(lang.ErrorMessageConnectionFailed)
		t.Wait()
		// Print a message with the IP address and port number.
		t.ResetDelay()
		msg := fmt.Sprintf(lang.TxtDialing, t.Helpers.RandomIP(), t.Helpers.RandomPort())
		fmt.Println(msg)
		t.Wait()
		t.SetDelayInSec(t.Helpers.RandomFloat(1, 5))
	}

	// Print a message
	fmt.Println(lang.TxtConnected)
	t.Wait()
	t.SetDelayInMs(oldDelay)
}

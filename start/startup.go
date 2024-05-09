package start

import (
	"fmt"

	"github.com/mt1976/crt/terminal"
	cfg "github.com/mt1976/mockterm/config"
	l "github.com/mt1976/mockterm/language"
)

// Run initializes the terminal and runs the main loop.

var c = cfg.Configuration

// Run initializes the terminal and runs the main loop.
func Run(t *terminal.ViewPort) {
	// Clear the terminal screen.
	t.Clear()

	// Display the banner.
	t.Banner("Initialise...")
	t.Wait()

	// Print a message.
	fmt.Println(l.TxtStartingTerminal)
	t.Wait()

	// Print a message.
	fmt.Println(l.TxtSelfTesting)
	t.Wait()
	oldDelay := t.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	t.SetDelayInSec(0.25)
	fmt.Println(l.TxtSelfTesting + l.TxtComplete)
	t.Wait()
	// Print the current date and time.
	t.SetDelayInMs(oldDelay)
	fmt.Println(l.TxtCurrentDate + t.Formatters.DateString())
	t.Wait()
	fmt.Println(l.TxtCurrentTime + t.Formatters.TimeString())
	t.Wait()

	// Print a message.
	fmt.Println(l.TxtPleaseWait)
	t.Wait()

	// Check if the terminal has a baud rate set.
	if !t.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(l.TxtBaudRate, t.Baud())
		fmt.Println(msg)
		t.Wait()
	}

	// Print a message.
	fmt.Println(l.TxtConnecting)
	t.Wait()

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(l.TxtDialing, t.Helpers.RandomIP(), t.Helpers.RandomPort())
	fmt.Println(msg)
	t.Wait()
	if !c.Debug {
		t.SetDelayInSec(t.Helpers.RandomFloat(1, 5))
	}
	if t.Helpers.CoinToss() && !c.Debug {
		fmt.Println(l.ErrorMessageConnectionFailed)
		t.Wait()
		// Print a message with the IP address and port number.
		t.ResetDelay()
		msg := fmt.Sprintf(l.TxtDialing, t.Helpers.RandomIP(), t.Helpers.RandomPort())
		fmt.Println(msg)
		t.Wait()
		t.SetDelayInSec(t.Helpers.RandomFloat(1, 5))
	}

	// Print a message
	fmt.Println(l.TxtConnected)
	t.Wait()
	t.SetDelayInMs(oldDelay)
}

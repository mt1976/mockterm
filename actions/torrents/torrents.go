package skynews

import (
	support "github.com/mt1976/crt"

	config "github.com/mt1976/mockterm/config"
	t "github.com/mt1976/mockterm/language"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(crt *support.Crt) {

	C := config.Configuration

	crt.Clear()
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := support.NewPageWithName(t.TxtTorrentsMenuTitle)
	c := 0
	c++
	m.AddOption(c, t.TxtTransmission, C.TransmissionURI, "")
	c++
	m.AddOption(c, t.TxtQTorrent, C.QTorrentURI, "")
	c++

	m.AddAction(t.SymActionQuit)

	action, nextLevel := m.Display(crt)

	if action == t.SymActionQuit {
		return
	}

	if support.IsInt(action) {
		switch action {
		case "1":
			Trans(crt, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		case "2":
			//QTor(crt, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		default:
			crt.InputError(support.ErrInvalidAction, action)
			action = ""
		}
	}
}

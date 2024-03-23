package skynews

import (
	term "github.com/mt1976/crt"
	conf "github.com/mt1976/mockterm/config"
	lang "github.com/mt1976/mockterm/language"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(crt *term.Crt) {

	C := conf.Configuration

	crt.Clear()
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := crt.NewTitledPage(lang.TxtTorrentsMenuTitle)
	c := 0
	c++
	m.AddOption(c, lang.TxtTransmission, C.TransmissionURI, "")
	c++
	m.AddOption(c, lang.TxtQTorrent, C.QTorrentURI, "")
	c++

	m.AddAction(lang.SymActionQuit)

	action, nextLevel := m.Display(crt)

	if action == lang.SymActionQuit {
		return
	}

	if crt.Helpers.IsInt(action) {
		switch action {
		case "1":
			Trans(crt, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		case "2":
			//QTor(crt, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		default:
			m.PageError(crt, term.ErrInvalidAction, action)
			action = ""
		}
	}
}

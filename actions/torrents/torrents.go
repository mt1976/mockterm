package skynews

import (
	term "github.com/mt1976/crt"
	conf "github.com/mt1976/mockterm/config"
	lang "github.com/mt1976/mockterm/language"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(t *term.ViewPort) {

	C := conf.Configuration

	t.Clear()

	m := t.NewTitledPage(lang.TxtTorrentsMenuTitle)
	c := 0
	c++
	m.AddOption(c, lang.TxtTransmission, C.TransmissionURI, "")
	c++
	m.AddOption(c, lang.TxtQTorrent, C.QTorrentURI, "")
	c++

	m.AddAction(lang.SymActionQuit)

	action, nextLevel := m.DisplayWithActions()

	if action == lang.SymActionQuit {
		return
	}

	if t.Helpers.IsInt(action) {
		switch action {
		case "1":
			Trans(t, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		case "2":
			//QTor(t, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		default:
			m.Error(term.ErrInvalidAction, action)
			action = ""
		}
	}
}

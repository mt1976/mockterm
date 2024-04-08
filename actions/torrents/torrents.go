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

	p := t.NewPage(lang.TxtTorrentsMenuTitle)
	c := 0
	c++
	p.AddMenuOption(c, lang.TxtTransmission, C.TransmissionURI, "")
	c++
	p.AddMenuOption(c, lang.TxtQTorrent, C.QTorrentURI, "")
	c++

	p.AddAction(lang.SymActionQuit)

	action, nextLevel := p.Display_Actions()

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
			p.Error(term.ErrInvalidAction, action)
			action = ""
		}
	}
}

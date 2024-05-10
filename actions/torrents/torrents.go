package skynews

import (
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	conf "github.com/mt1976/mockterm/config"
	lang "github.com/mt1976/mockterm/language"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(t *term.ViewPort) {

	C := conf.Configuration

	t.Clear()

	p := page.NewPage(t, lang.Torrents.Text())
	c := 0
	c++
	p.AddMenuOption(c, lang.TxtTransmission, C.TransmissionURI, "")
	c++
	p.AddMenuOption(c, lang.TxtQTorrent, C.QTorrentURI, "")
	c++

	p.AddAction(acts.Quit)

	action := p.Display_Actions()

	if action == acts.Quit {
		return
	}

	if action.IsInt() {
		switch action.Action() {
		case "1":
			//	Trans(t, nextLevel.AlternateID, nextLevel.Title)
			//action = ""
			return
		case "2":
			//QTor(t, nextLevel.AlternateID, nextLevel.Title)
			//action = ""
			return
		default:
			p.Error(terr.ErrInvalidAction, action.Action())
			return
		}
	}
}

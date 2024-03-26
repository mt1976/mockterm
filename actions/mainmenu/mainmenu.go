package mainmenu

import (
	term "github.com/mt1976/crt"
	dash "github.com/mt1976/mockterm/actions/dashboard"
	plex "github.com/mt1976/mockterm/actions/plexmediaserver"
	news "github.com/mt1976/mockterm/actions/skynews"
	syst "github.com/mt1976/mockterm/actions/systemsmenu"
	trts "github.com/mt1976/mockterm/actions/torrents"
	wthr "github.com/mt1976/mockterm/actions/weather"
	lang "github.com/mt1976/mockterm/language"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(t *term.ViewPort) {
	//log.Println("Starting Main Menu")
	p := t.NewPage(lang.TxtMainMenuTitle)
	p.AddMenuOption(1, lang.TxtDashboardTitle, "", "")
	p.AddMenuOption(2, lang.TxtSkyNewsMenuTitle, "", "")
	p.AddMenuOption(3, lang.TxtBBCNewsMenuTitle, "", "")
	p.AddMenuOption(4, lang.TxtWeatherMenuTitle, "", "")
	p.AddMenuOption(5, lang.TxtTorrentsMenuTitle, "", "")
	p.AddMenuOption(6, lang.TxtPlexMediaServersMenuTitle, "", "")
	p.AddMenuOption(7, lang.TxtRemoteSystemsAccessMenuTitle, "", "")
	p.AddMenuOption(8, lang.TxtSystemsMaintenanceMenuTitle, "", "")
	p.AddMenuOption(9, lang.SymBlank, "", "")
	p.AddMenuOption(10, lang.SymBlank, "", "")
	p.AddAction(lang.SymActionQuit)

	ok := false
	for !ok {

		p.Clear()

		action, _ := p.DisplayWithActions()
		switch action {
		case lang.SymActionQuit:
			p.Info(lang.TxtQuittingMessage)
			ok = true
			continue
		case "1":
			dash.Run(t)
		case "2":
			news.Run(t)
		case "4":
			wthr.Run(t)
		case "5":
			trts.Run(t)
		case "6":
			plex.Run(t)
		case "8":
			syst.Run(t)
		default:
			p.Error(term.ErrInvalidAction, action)
		}
	}
}

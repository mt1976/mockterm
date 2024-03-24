package mainmenu

import (
	"log"

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
	log.Println("Starting Main Menu")
	p := t.NewTitledPage(lang.TxtMainMenuTitle)
	//	spew.Dump(p)
	//	log.Println("Adding Menu Options")
	//for i := range 11 {
	//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	//}

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
	//spew.Dump(p)
	//os.Exit(0)
	// loop while ok
	ok := false
	for !ok {

		t.Clear()

		action, _ := p.DisplayWithActions()
		switch action {
		case lang.SymActionQuit:
			t.InfoMessage(lang.TxtQuittingMessage + lang.SymNewline)
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

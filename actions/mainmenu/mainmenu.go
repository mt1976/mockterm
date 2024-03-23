package mainmenu

import (
	support "github.com/mt1976/crt"
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
func Run(crt *support.Crt) {

	m := crt.NewTitledPage(lang.TxtMainMenuTitle)
	//for i := range 11 {
	//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	//}

	m.AddOption(1, lang.TxtDashboardTitle, "", "")
	m.AddOption(2, lang.TxtSkyNewsMenuTitle, "", "")
	m.AddOption(3, lang.TxtBBCNewsMenuTitle, "", "")
	m.AddOption(4, lang.TxtWeatherMenuTitle, "", "")
	m.AddOption(5, lang.TxtTorrentsMenuTitle, "", "")
	m.AddOption(6, lang.TxtPlexMediaServersMenuTitle, "", "")
	m.AddOption(7, lang.TxtRemoteSystemsAccessMenuTitle, "", "")
	m.AddOption(8, lang.TxtSystemsMaintenanceMenuTitle, "", "")
	m.AddOption(9, lang.SymBlank, "", "")
	m.AddOption(10, lang.SymBlank, "", "")
	m.AddAction(lang.SymActionQuit)

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		//crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")

		action, _ := m.DisplayWithActions(crt)
		switch action {
		case lang.SymActionQuit:
			crt.InfoMessage(lang.TxtQuittingMessage + lang.SymNewline)
			ok = true
			continue
		case "1":
			dash.Run(crt)
			//action, _ = y.Display(crt)
		case "2":
			news.Run(crt)
		case "4":
			wthr.Run(crt)
		case "5":
			trts.Run(crt)
		case "6":
			plex.Run(crt)
		case "8":
			syst.Run(crt)
		default:
			m.Error(crt, support.ErrInvalidAction, action)
		}
	}
}

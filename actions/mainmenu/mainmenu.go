package mainmenu

import (
	support "github.com/mt1976/crt"
	dash "github.com/mt1976/mockterm/actions/dashboard"
	pxms "github.com/mt1976/mockterm/actions/plexmediaserver"
	news "github.com/mt1976/mockterm/actions/skynews"
	trts "github.com/mt1976/mockterm/actions/torrents"
	wthr "github.com/mt1976/mockterm/actions/weather"
	text "github.com/mt1976/mockterm/language"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(crt *support.Crt) {

	m := support.NewPageWithName(text.TxtMainMenuTitle)
	//for i := range 11 {
	//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	//}

	m.AddOption(1, text.TxtDashboardTitle, "", "")
	m.AddOption(2, text.TxtSkyNewsMenuTitle, "", "")
	m.AddOption(3, text.TxtBBCNewsMenuTitle, "", "")
	m.AddOption(4, text.TxtWeatherMenuTitle, "", "")
	m.AddOption(5, text.TxtTorrentsMenuTitle, "", "")
	m.AddOption(6, text.TxtPlexMediaServersMenuTitle, "", "")
	m.AddOption(7, text.TxtRemoteSystemsAccessMenuTitle, "", "")
	m.AddOption(8, text.TxtSystemsMaintenanceMenuTitle, "", "")
	m.AddOption(9, text.SymBlank, "", "")
	m.AddOption(10, text.SymBlank, "", "")
	m.AddAction(text.SymActionQuit)

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		//crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")

		action, _ := m.Display(crt)
		switch action {
		case text.SymActionQuit:
			crt.InfoMessage(text.TxtQuittingMessage + text.SymNewline)
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
			pxms.Run(crt)
		default:
			crt.InputError(support.ErrInvalidAction, action)
		}
	}
}

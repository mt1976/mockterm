package mainmenu

import (
	"os"

	e "github.com/mt1976/crt/errors"
	t "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/page"
	"github.com/mt1976/mockterm/actions/dashboard"
	plexmediaserver "github.com/mt1976/mockterm/actions/plexms"
	"github.com/mt1976/mockterm/actions/skynews"
	torrents "github.com/mt1976/mockterm/actions/torrents"
	"github.com/mt1976/mockterm/actions/weather"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(crt *support.Crt) {

	m := page.New(t.TxtMainMenuTitle)
	//for i := range 11 {
	//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	//}

	m.AddOption(1, t.TxtDashboardTitle, "", "")
	m.AddOption(2, t.TxtSkyNewsMenuTitle, "", "")
	m.AddOption(3, t.TxtBBCNewsMenuTitle, "", "")
	m.AddOption(4, t.TxtWeatherMenuTitle, "", "")
	m.AddOption(5, t.TxtTorrentsMenuTitle, "", "")
	m.AddOption(6, t.TxtPlexMediaServersMenuTitle, "", "")
	m.AddOption(7, t.TxtRemoteSystemsAccessMenuTitle, "", "")
	m.AddOption(8, t.TxtSystemsMaintenanceMenuTitle, "", "")
	m.AddOption(9, t.SymBlank, "", "")
	m.AddOption(10, t.SymBlank, "", "")
	m.AddAction(t.SymActionQuit)

	os.Exit(0)

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		//crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")

		action, _ := m.Display(crt)
		switch action {
		case t.SymActionQuit:
			crt.InfoMessage(t.TxtQuittingMessage + t.SymNewline)
			ok = true
			continue
		case "1":
			dashboard.Run(crt)
			//action, _ = y.Display(crt)
		case "2":
			skynews.Run(crt)
		case "4":
			weather.Run(crt)
		case "5":
			torrents.Run(crt)
		case "6":
			plexmediaserver.Run(crt)
		default:
			crt.InputError(e.ErrInvalidAction + support.SQuote(action))
		}
	}
}

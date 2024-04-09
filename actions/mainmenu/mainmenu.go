package mainmenu

import (
	"errors"
	"fmt"

	term "github.com/mt1976/crt"
	file "github.com/mt1976/crt/filechooser"
	"github.com/mt1976/mockterm/actions/bbcnews"
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
	p.AddBlankRow()
	p.AddMenuOption(1, lang.TxtDashboardTitle, "", "")
	p.AddMenuOption(2, lang.TxtSkyNewsMenuTitle, "", "")
	p.AddMenuOption(3, lang.TxtBBCNewsMenuTitle, "", "")
	p.AddMenuOption(4, lang.TxtWeatherMenuTitle, "", "")
	p.AddMenuOption(5, lang.TxtTorrentsMenuTitle, "", "")
	p.AddMenuOption(6, lang.TxtPlexMediaServersMenuTitle, "", "")
	p.AddBlankRow()
	p.AddMenuOption(7, lang.TxtRemoteSystemsAccessMenuTitle, "", "")
	p.AddMenuOption(8, lang.TxtSystemsMaintenanceMenuTitle, "", "")
	//p.AddMenuOption(9, lang.SymBlank, "", "")
	p.AddAction(lang.SymActionQuit)

	ok := false
	for !ok {

		action := p.Display_Actions()
		switch action {
		case lang.SymActionQuit:
			p.Info(lang.TxtQuittingMessage)
			ok = true
			continue
		case "1":
			dash.Run(t)
		case "2":
			news.Run(t)
		case "3":
			err := bbcnews.Run(t)
			if err != nil {
				t.Error(err, "")
				return
			}
		case "4":
			wthr.Run(t)
		case "5":
			trts.Run(t)
		case "6":
			plex.Run(t)
		case "8":
			syst.Run(t)
		case "10":
			userHome, err := file.UserHome()
			if err != nil {
				t.Error(err)
			}
			selected, isDir, err := file.FileChooser(userHome, file.All)
			if err != nil {
				p.Error(err, "file chooser error")
			}
			if isDir {
				_, _, err := file.FileChooser(selected, file.All)
				if err != nil {
					p.Error(err, "file chooser error")
				}
			}
			prn := fmt.Sprintf("Selected = %v, isDir = %v", selected, isDir)
			p.Error(errors.New("testing"), prn)

		default:
			p.Error(term.ErrInvalidAction, action)
		}
	}
}

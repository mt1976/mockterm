package mainmenu

import (
	"errors"
	"fmt"

	errs "github.com/mt1976/crt/errors"
	file "github.com/mt1976/crt/filechooser"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	dash "github.com/mt1976/mockterm/actions/dashboard"
	plex "github.com/mt1976/mockterm/actions/plexmediaserver"
	srss "github.com/mt1976/mockterm/actions/showsrss"
	news "github.com/mt1976/mockterm/actions/skynews"
	syst "github.com/mt1976/mockterm/actions/systems"
	tfl "github.com/mt1976/mockterm/actions/tfl"
	wthr "github.com/mt1976/mockterm/actions/weather"
	lang "github.com/mt1976/mockterm/language"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(terminal *term.ViewPort) {
	//log.Println("Starting Main Menu")
	p := page.NewPage(terminal, lang.TxtMainMenuTitle)
	p.AddBlankRow()
	p.AddMenuOption(1, lang.TxtDashboardTitle, "", "")
	p.AddMenuOption(2, lang.TxtSkyNewsMenuTitle, "", "")
	p.AddMenuOption(3, "Shows RSS", "", "")
	p.AddMenuOption(4, lang.TxtWeatherMenuTitle, "", "")
	p.AddMenuOption(5, "TFL", "", "")
	p.AddMenuOption(6, lang.TxtPlexMediaServersMenuTitle, "", "")
	p.AddBlankRow()
	p.AddMenuOption(7, lang.TxtRemoteSystemsAccessMenuTitle, "", "")
	p.AddMenuOption(8, lang.TxtSystemsMaintenanceMenuTitle, "", "")
	//p.AddMenuOption(9, lang.SymBlank, "", "")
	p.AddAction(acts.Quit)

	ok := false
	for !ok {

		action := p.Display_Actions()
		switch {
		case action.Is(acts.Quit):
			p.Info(lang.TxtQuittingMessage + " - " + lang.TxtThankYouForUsing + " " + lang.TxtApplicationName)
			ok = true
			continue
		case action.IsInt() && action.Int() == 1:
			dash.Run(terminal)
		case action.IsInt() && action.Int() == 2:
			news.Run(terminal)
		case action.IsInt() && action.Int() == 3:
			err := srss.Run(terminal)
			if err != nil {
				terminal.Error(err, "error running showsrss")
				return
			}
		case action.IsInt() && action.Int() == 4:
			wthr.Run(terminal)
		case action.IsInt() && action.Int() == 5:
			tfl.Run(terminal)
		case action.IsInt() && action.Int() == 6:
			plex.Run(terminal)
		case action.IsInt() && action.Int() == 8:
			syst.Run(terminal)
		case action.IsInt() && action.Int() == 10:
			userHome, err := file.UserHome()
			if err != nil {
				terminal.Error(err)
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
			p.Error(errs.ErrInvalidAction, action.Action())
		}
	}
}

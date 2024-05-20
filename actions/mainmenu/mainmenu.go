package mainmenu

import (
	"errors"
	"fmt"

	errs "github.com/mt1976/crt/errors"
	f "github.com/mt1976/crt/filechooser"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	"github.com/mt1976/crt/terminal"
	dash "github.com/mt1976/mockterm/actions/dashboard"
	dlng "github.com/mt1976/mockterm/actions/dashboard/lang"
	plex "github.com/mt1976/mockterm/actions/plexmediaserver"
	srss "github.com/mt1976/mockterm/actions/showsrss"
	news "github.com/mt1976/mockterm/actions/skynews"
	nlng "github.com/mt1976/mockterm/actions/skynews/lang"
	syst "github.com/mt1976/mockterm/actions/systems"
	tfl "github.com/mt1976/mockterm/actions/tfl"
	wthr "github.com/mt1976/mockterm/actions/weather"
	l "github.com/mt1976/mockterm/language"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(terminal *terminal.ViewPort) {
	//log.Println("Starting Main Menu")
	p := page.NewPage(terminal, l.MainMenuTitle)
	p.AddBlankRow()
	p.AddMenuOption(1, dlng.Title.Text(), "", "")
	p.AddMenuOption(2, nlng.Title.Text(), "", "")
	p.AddMenuOption(3, l.ShowsRSS.Text(), "", "")
	p.AddMenuOption(4, l.Weather.Text(), "", "")
	p.AddMenuOption(5, l.TFL.Text(), "", "")
	p.AddMenuOption(6, l.PlexMediaServer.Text(), "", "")
	p.AddBlankRow()
	p.AddMenuOption(7, l.RemoteSystemsAccess.Text(), "", "")
	p.AddMenuOption(8, l.SystemsMaintenance.Text(), "", "")
	//p.AddMenuOption(9, lang.SymBlank, "", "")
	p.AddAction(acts.Quit)

	ok := false
	for !ok {

		action := p.Display_Actions()
		switch {
		case action.Is(acts.Quit):
			p.Info(l.Quitting, fmt.Sprintf(l.ThankYouForUsing.Text(), l.ApplicationName.Text()))
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
			userHome, err := f.UserHome()
			if err != nil {
				terminal.Error(err)
			}
			selected, isDir, err := f.FileChooser(userHome, f.All)
			if err != nil {
				p.Error(err, "file chooser error")
			}
			if isDir {
				_, _, err := f.FileChooser(selected, f.All)
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

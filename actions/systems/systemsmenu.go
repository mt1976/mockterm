package systemsmenu

import (
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	catalog "github.com/mt1976/mockterm/actions/systems/catalog"
	ctlg "github.com/mt1976/mockterm/actions/systems/catalog/lang"
	clean "github.com/mt1976/mockterm/actions/systems/cleanfilenames"
	slng "github.com/mt1976/mockterm/actions/systems/lang"
	move "github.com/mt1976/mockterm/actions/systems/mover"
	flng "github.com/mt1976/mockterm/actions/systems/mover/lang"
	push "github.com/mt1976/mockterm/actions/systems/pushover"
	plng "github.com/mt1976/mockterm/actions/systems/pushover/lang"
	tidy "github.com/mt1976/mockterm/actions/systems/tidymediafolders"
	tlng "github.com/mt1976/mockterm/actions/systems/tidymediafolders/lang"
	lang "github.com/mt1976/mockterm/language"
	mode "github.com/mt1976/mockterm/support/modes"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(t *term.ViewPort) {

	p := page.NewPage(t, lang.SystemsMaintenance)
	p.AddParagraph(slng.ServiceMenuDescription.String())
	p.AddBlankRow()
	p.AddMenuOption(1, plng.TxtPushoverTitle.Text(), "", "")
	p.AddBlankRow()
	p.AddMenuOption(2, tlng.Title.Text()+" (Trial Mode)", "", "")
	p.AddMenuOption(3, tlng.Title.Text()+" (LIVE)", "", "")
	p.AddBlankRow()
	p.AddMenuOption(4, tlng.CleanFileNames.Text()+" (Trial Mode)", "", "")
	p.AddMenuOption(5, tlng.CleanFileNames.Text()+" (LIVE)", "", "")
	p.AddBlankRow()
	p.AddMenuOption(6, ctlg.Title.Text(), "", "")
	p.AddBlankRow()
	p.AddMenuOption(7, flng.Title.Text()+" (Trial Mode)", "", "")
	p.AddMenuOption(8, flng.Title.Text()+" (LIVE)", "", "")

	p.AddAction(acts.Quit)

	// loop while ok
	ok := false
	for !ok {

		t.Clear()

		action := p.Display_Actions()
		switch {
		case action.Is(acts.Quit):
			p.Info(lang.Quitting, "", "")
			ok = true
			continue
		case action.Equals("1"):
			push.Run(t)
		case action.Equals("2"):
			tidy.Run(t, mode.TRIAL, "")
		case action.Equals("3"):
			tidy.Run(t, mode.LIVE, "")
		case action.Equals("4"):
			clean.Run(t, mode.TRIAL, "")
		case action.Equals("5"):
			clean.Run(t, mode.LIVE, "")
		case action.Equals("6"):
			catalog.Run(t)
		case action.Equals("7"):
			move.Run(t, mode.TRIAL)
		case action.Equals("8"):
			move.Run(t, mode.LIVE)

		default:
			p.Error(terr.ErrInvalidAction, action.Action())
		}
	}
}

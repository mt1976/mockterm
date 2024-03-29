package systemsmenu

import (
	term "github.com/mt1976/crt"
	push "github.com/mt1976/mockterm/actions/pushover"
	lang "github.com/mt1976/mockterm/language"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(t *term.ViewPort) {

	p := t.NewPage(lang.TxtSystemsMaintenanceMenuTitle)
	p.AddParagraph(lang.TxtServiceMenuDescription)
	p.AddBlankRow()
	p.AddMenuOption(1, lang.TxtPushoverTitle, "", "")
	p.AddAction(lang.SymActionQuit)

	// loop while ok
	ok := false
	for !ok {

		t.Clear()

		action, _ := p.DisplayWithActions()
		switch action {
		case lang.SymActionQuit:
			p.Info(lang.TxtQuittingMessage)
			ok = true
			continue
		case "1":
			push.Run(t)
		default:
			p.Error(term.ErrInvalidAction, action)
		}
	}
}

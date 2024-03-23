package systemsmenu

import (
	support "github.com/mt1976/crt"
	push "github.com/mt1976/mockterm/actions/pushover"
	lang "github.com/mt1976/mockterm/language"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(crt *support.ViewPort) {

	m := crt.NewTitledPage(lang.TxtSystemsMaintenanceMenuTitle)
	m.Paragraph(lang.TxtServiceMenuDescription)
	m.BlankRow()
	m.AddOption(1, lang.TxtPushoverTitle, "", "")
	m.AddAction(lang.SymActionQuit)

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		//crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")

		action, _ := m.DisplayWithActions()
		switch action {
		case lang.SymActionQuit:
			crt.InfoMessage(lang.TxtQuittingMessage + lang.SymNewline)
			ok = true
			continue
		case "1":
			push.Run(crt)
		default:
			m.Error(support.ErrInvalidAction, action)
		}
	}
}

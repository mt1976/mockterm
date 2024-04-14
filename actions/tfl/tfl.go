package tfl

import (
	terminal "github.com/mt1976/crt"
	tfl "github.com/mt1976/mockterm/actions/tfl/tfler"
	lang "github.com/mt1976/mockterm/language"
)

func Run(terminal *terminal.ViewPort) {
	terminal.Print("TFL API")

	page := terminal.NewPage("Transport for London")

	tubeLines, _ := tfl.GetTubeLines()
	lineTitle := terminal.Styles.Bold("Line")
	statusTitle := terminal.Styles.Bold("Status")
	page.AddColumnsTitle(lineTitle, statusTitle)
	for _, s := range tubeLines {
		//page.AddFieldValuePair(s.Name, s.Status)
		page.AddColumns(s.Name, s.Status)
	}
	for {
		action := page.Display_Actions()
		if terminal.Formatters.Upcase(action) == lang.SymActionQuit {
			return
		}
	}

	//tfl.Client.Do()
}

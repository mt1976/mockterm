package tfl

import (
	terminal "github.com/mt1976/crt"
	tfl "github.com/mt1976/mockterm/actions/tfl/tfler"
	lang "github.com/mt1976/mockterm/language"
)

func Run(terminal *terminal.ViewPort) {
	terminal.Print("TFL API")

	page := terminal.NewPage("Transport for London")

	page.AddColumnsTitle("One")
	page.AddColumnsTitle("One", "Two")
	page.AddColumnsTitle("One", "Two", "Three")
	page.AddColumnsTitle("One", "Two", "Three", "Four")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five", "Six")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five", "Six", "Seven")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven")
	page.AddColumnsTitle("One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve")
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

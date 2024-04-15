package tfl

import (
	"fmt"

	terminal "github.com/mt1976/crt"
	tfl "github.com/mt1976/mockterm/actions/tfl/tfler"
	lang "github.com/mt1976/mockterm/language"
)

func Run(terminal *terminal.ViewPort) {
	terminal.Print("TFL API")

	page := terminal.NewPage("Transport for London - Line Status")

	tubeLines, _ := tfl.GetTubeLines()
	// lineTitle := "Line"
	// statusTitle := "Status"

	//page.AddColumnsTitle(lineTitle, statusTitle)
	//	title := "TFL Tube Line Status"
	//	page.AddParagraphString(title)
	//	page.AddParagraphString(strings.Repeat("=", len(title)))
	page.AddBlankRow()
	for i, s := range tubeLines {
		//page.AddFieldValuePair(s.Name, s.Status)
		row := fmt.Sprintf("%-20v %v", s.Name, s.Status)
		page.AddMenuOption(i+1, row, s.Code, "")
	}
	for {
		action := page.Display_Actions()
		if terminal.Formatters.Upcase(action) == lang.SymActionQuit {
			return
		}
	}
}

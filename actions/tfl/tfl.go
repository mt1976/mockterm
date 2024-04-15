package tfl

import (
	"fmt"

	terminal "github.com/mt1976/crt"
	tfl "github.com/mt1976/mockterm/actions/tfl/tfler"
	lang "github.com/mt1976/mockterm/language"
)

func Run(terminal *terminal.ViewPort) {

	// var isNumeric func(string) bool
	isNumeric := terminal.Helpers.IsInt
	toInt := terminal.Helpers.ToInt

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
		switch {
		case terminal.Formatters.Upcase(action) == lang.SymActionQuit:
			return
		case isNumeric(action):
			num := toInt(action)
			if num >= 0 && num <= len(tubeLines) {
				Detail(terminal, tubeLines[num-1])
			}
		}
	}
}

func Detail(terminal *terminal.ViewPort, line tfl.Line) {
	page := terminal.NewPage("Transport for London - Line Status")

	lineDetail, err := tfl.GetTubeLineDetails(line.Code)
	if err != nil {
		page.AddParagraphString("Error: " + err.Error())
		return
	}

	page.AddParagraphString("Line: " + lineDetail.Name)
	page.AddParagraphString("Code: " + lineDetail.Code)
	page.AddBlankRow()
	page.AddColumnsTitle("Station", "Status")
	for _, s := range lineDetail.Stations {
		page.AddColumns(s.Name, s.Status)
	}

	page.Display_Confirmation("Press Y to continue...")
}

package tfl

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	terminal "github.com/mt1976/crt"
	tfl "github.com/mt1976/mockterm/actions/tfl/tfler"
	lang "github.com/mt1976/mockterm/language"
)

var isNumeric func(string) bool
var toInt func(string) int

func Run(terminal *terminal.ViewPort) {

	// var isNumeric func(string) bool
	isNumeric = terminal.Helpers.IsInt
	toInt = terminal.Helpers.ToInt

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
	spew.Dump(tubeLines)
	for {
		action := page.Display_Actions()
		switch {
		case terminal.Formatters.Upcase(action) == lang.SymActionQuit:
			return
		case isNumeric(action):
			num := toInt(action)
			if num >= 0 && num <= len(tubeLines) {
				LineDetail(terminal, tubeLines[num-1])
			}
		}
	}
}

func LineDetail(terminal *terminal.ViewPort, line tfl.Line) {
	page := terminal.NewPage("Transport for London - Line Details")

	lineDetail, err := tfl.GetTubeLineDetails(line.Code)
	if err != nil {
		terminal.Error(err, "Error: "+err.Error())
		os.Exit(1)
		return
	}

	spew.Dump(lineDetail)
	page.AddFieldValuePair("Line", lineDetail.Name)
	page.AddFieldValuePair("Code", lineDetail.Code)
	page.AddFieldValuePair("Status", line.Status)
	page.AddBreakRow()
	//page.AddColumnsTitle("Station", "Status")
	for i, s := range lineDetail.Stations {
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
			if num >= 0 && num <= len(lineDetail.Stations) {
				StationDetail(terminal, lineDetail.Stations[num-1])
			}
		}
	}
}

func StationDetail(terminal *terminal.ViewPort, station tfl.Station) {
	page := terminal.NewPage("Transport for London - Station Status")
	stationDetail, err := tfl.GetStationDetails(station.Code)
	if err != nil {
		terminal.Error(err, "Error getting station details")
		return
	}
	spew.Dump(stationDetail)

	page.AddFieldValuePair("Station", stationDetail.Name)
	page.AddFieldValuePair("Code", stationDetail.Code)
	page.AddFieldValuePair("Status", stationDetail.Status)
	page.AddBreakRow()

	for {
		action := page.Display_Actions()
		switch {
		case terminal.Formatters.Upcase(action) == lang.SymActionQuit:
			return
		}
	}
}

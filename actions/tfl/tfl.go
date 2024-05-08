package tfl

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	tfl "github.com/mt1976/mockterm/actions/tfl/tfler"
)

var isNumeric func(string) bool
var toInt func(string) int

func Run(t *term.ViewPort) {

	// var isNumeric func(string) bool
	isNumeric = t.Helpers.IsInt
	toInt = t.Helpers.ToInt

	t.Print("TFL API")

	p := page.NewPage(t, "Transport for London - Line Status")

	tubeLines, _ := tfl.GetTubeLines()
	// lineTitle := "Line"
	// statusTitle := "Status"

	//page.AddColumnsTitle(lineTitle, statusTitle)
	//	title := "TFL Tube Line Status"
	//	page.AddParagraphString(title)
	//	page.AddParagraphString(strings.Repeat("=", len(title)))
	p.AddBlankRow()
	for i, s := range tubeLines {
		//page.AddFieldValuePair(s.Name, s.Status)
		row := fmt.Sprintf("%-20v %v", s.Name, s.Status)
		p.AddMenuOption(i+1, row, s.Code, "")
	}
	spew.Dump(tubeLines)
	for {
		action := p.Display_Actions()
		switch {
		case action.Is(acts.Quit):
			//terminal.Formatters.Upcase(action) == lang.SymActionQuit:
			return
		case action.IsInt():
			num := toInt(action.Action())
			if num >= 0 && num <= len(tubeLines) {
				LineDetail(t, tubeLines[num-1])
			}
		}
	}
}

func LineDetail(t *term.ViewPort, line tfl.Line) {
	page := page.NewPage(t, "Transport for London - Line Details")

	lineDetail, err := tfl.GetTubeLineDetails(line.Code)
	if err != nil {
		t.Error(err, "Error: "+err.Error())
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
		case action.Is(acts.Quit):
			return
		case action.IsInt():
			num := toInt(action.Action())
			if num >= 0 && num <= len(lineDetail.Stations) {
				StationDetail(t, lineDetail.Stations[num-1])
			}
		}
	}
}

func StationDetail(t *term.ViewPort, station tfl.Station) {
	p := page.NewPage(t, "Transport for London - Station Status")
	stationDetail, err := tfl.GetStationDetails(station.Code)
	if err != nil {
		t.Error(err, "Error getting station details")
		return
	}
	spew.Dump(stationDetail)

	p.AddFieldValuePair("Station", stationDetail.Name)
	p.AddFieldValuePair("Code", stationDetail.Code)
	p.AddFieldValuePair("Status", stationDetail.Status)
	p.AddBreakRow()

	for {
		action := p.Display_Actions()
		switch {
		case action.Is(acts.Quit):
			//t.Formatters.Upcase(action) == lang.SymActionQuit:
			return
		}
	}
}

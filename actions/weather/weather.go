package weather

import (
	"fmt"
	"os"
	"strings"

	owm "github.com/briandowns/openweathermap"
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

var C = conf.Configuration

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(t *term.ViewPort) {

	//t.Clear()
	p := page.NewPage(t, lang.TxtWeatherTitle)
	w, err := owm.NewCurrent(C.OpenWeatherMapApiUnits, C.OpenWeatherMapApiLang, C.OpenWeatherMapApiKey)
	if err != nil {
		p.Error(errs.ErrOpenWeather, err.Error())
		os.Exit(1)
		return
	}

	err = w.CurrentByCoordinates(
		&owm.Coordinates{Latitude: C.LocationLatitude, Longitude: C.LocationLogitude})
	if err != nil {
		p.Error(errs.ErrOpenWeather, err.Error())
		os.Exit(1)
		return
	}

	c := 0
	c++
	p.Add(fmt.Sprintf(lang.SymWeatherFormat2, lang.TxtLocationLabel, t.Formatters.Bold(w.Name)), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat2, lang.TxtConditionsLabel, t.Formatters.Bold(w.Weather[0].Main)), "", "")
	p.Add(hr(t), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtTemperatureLabel, boldFloat(t, w.Main.Temp)+lang.SymDegree, lang.TxtFeelsLikeLabel, boldFloat(t, w.Main.FeelsLike)+lang.SymDegree), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtMinLabel, boldFloat(t, w.Main.TempMin)+lang.SymDegree, lang.TxtMaxLabel, boldFloat(t, w.Main.TempMax)+lang.SymDegree), "", "")
	//p.Add(hr())
	p.Add(hr(t), "", "")
	// p.Add(fmt.Sprintf("Feels Like : %v", w.Main.FeelsLike))
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtWindSpeedLabel, boldFloat(t, w.Wind.Speed), lang.TxtWindDirectionLabel, boldFloat(t, w.Wind.Deg)), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat1, lang.TxtCloudCoverLabel, boldInt(t, w.Clouds.All)), "", "")
	p.Add(hr(t), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtRain1Hr, boldFloat(t, w.Rain.OneH), lang.TxtRain3Hr, boldFloat(t, w.Rain.ThreeH)), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtSnow1Hr, boldFloat(t, w.Snow.OneH), lang.TxtSnow3Hr, boldFloat(t, w.Snow.ThreeH)), "", "")
	p.Add(hr(t), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtSunriseLabel, t.Formatters.Bold(outdate(t, w.Sys.Sunrise)), lang.TxtSunsetLabel, t.Formatters.Bold(outdate(t, w.Sys.Sunset))), "", "")
	p.Add(hr(t), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat2, lang.TxtSourceLabel, t.Formatters.Bold(lang.TxtSourceService)), "", "")
	// INSERT CONTENT ABOVE
	p.AddAction(acts.Quit)
	p.AddAction(acts.Forward)
	p.AddAction(acts.Back)
	p.Dump()
	ok := false
	for !ok {

		nextAction := p.Display_Actions()
		switch nextAction {
		case acts.Forward:
			p.Forward()
		case acts.Back:
			p.Back()
		case acts.Quit:
			ok = true
			return
		default:
			p.Error(terr.ErrInvalidAction, nextAction.Action())
		}
	}

}

func outdate(t *term.ViewPort, unixDateTime int) string {
	// int to int64
	// unix date to human date
	return t.Formatters.HumanFromUnixDate(int64(unixDateTime))
}

// The `hr` function returns a string consisting of a line of dashes.
func hr(t *term.ViewPort) string {
	screenwidth := t.Width() - 5
	//fmt.Printf("screenwidth: %v\n", screenwidth)
	return strings.Repeat("-", screenwidth)
}

func boldFloat(t *term.ViewPort, in float64) string {
	return t.Formatters.Bold(fmt.Sprintf("%v", in))
}

func boldInt(t *term.ViewPort, in int) string {
	return t.Formatters.Bold(fmt.Sprintf("%v", in))
}

package weather

import (
	"fmt"
	"os"
	"strings"

	owm "github.com/briandowns/openweathermap"
	term "github.com/mt1976/crt"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

var C conf.Config

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(crt *term.Crt) {

	crt.Clear()
	p := crt.NewTitledPage(lang.TxtWeatherTitle + lang.Space + lang.TxtSourceService)

	w, err := owm.NewCurrent(C.OpenWeatherMapApiUnits, C.OpenWeatherMapApiLang, C.OpenWeatherMapApiKey)
	if err != nil {
		crt.Error(errs.ErrOpenWeather, err.Error())
		os.Exit(1)
		return
	}

	err = w.CurrentByCoordinates(
		&owm.Coordinates{Latitude: C.LocationLatitude, Longitude: C.LocationLogitude})
	if err != nil {
		crt.Error(errs.ErrOpenWeather, err.Error())
		os.Exit(1)
		return
	}

	c := 0
	c++
	p.Add(fmt.Sprintf(lang.SymWeatherFormat2, lang.TxtLocationLabel, crt.Formatters.Bold(w.Name)), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat2, lang.TxtConditionsLabel, crt.Formatters.Bold(w.Weather[0].Main)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtTemperatureLabel, boldFloat(crt, w.Main.Temp)+lang.SymDegree, lang.TxtFeelsLikeLabel, boldFloat(crt, w.Main.FeelsLike)+lang.SymDegree), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtMinLabel, boldFloat(crt, w.Main.TempMin)+lang.SymDegree, lang.TxtMaxLabel, boldFloat(crt, w.Main.TempMax)+lang.SymDegree), "", "")
	//p.Add(hr())
	p.Add(hr(crt), "", "")
	// p.Add(fmt.Sprintf("Feels Like : %v", w.Main.FeelsLike))
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtWindSpeedLabel, boldFloat(crt, w.Wind.Speed), lang.TxtWindDirectionLabel, boldFloat(crt, w.Wind.Deg)), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat1, lang.TxtCloudCoverLabel, boldInt(crt, w.Clouds.All)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtRain1Hr, boldFloat(crt, w.Rain.OneH), lang.TxtRain3Hr, boldFloat(crt, w.Rain.ThreeH)), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtSnow1Hr, boldFloat(crt, w.Snow.OneH), lang.TxtSnow3Hr, boldFloat(crt, w.Snow.ThreeH)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat4, lang.TxtSunriseLabel, crt.Formatters.Bold(outdate(crt, w.Sys.Sunrise)), lang.TxtSunsetLabel, crt.Formatters.Bold(outdate(crt, w.Sys.Sunset))), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(lang.SymWeatherFormat2, lang.TxtSourceLabel, crt.Formatters.Bold(lang.TxtSourceService)), "", "")
	// INSERT CONTENT ABOVE
	p.AddAction(lang.SymActionQuit)
	p.AddAction(lang.SymActionForward)
	p.AddAction(lang.SymActionBack)
	ok := false
	for !ok {

		nextAction, _ := p.DisplayWithActions()
		switch nextAction {
		case lang.SymActionForward:
			p.NextPage()
		case lang.SymActionBack:
			p.PreviousPage()
		case lang.SymActionQuit:
			ok = true
			return
		default:
			p.Error(term.ErrInvalidAction, nextAction)
		}
	}

}

func outdate(crt *term.Crt, t int) string {
	// int to int64
	// unix date to human date
	return crt.Formatters.HumanFromUnixDate(int64(t))
}

// The `hr` function returns a string consisting of a line of dashes.
func hr(crt *term.Crt) string {
	return strings.Repeat(lang.SymBreak, 3)
}

func boldFloat(crt *term.Crt, in float64) string {
	return crt.Formatters.Bold(fmt.Sprintf("%v", in))
}

func boldInt(crt *term.Crt, in int) string {
	return crt.Formatters.Bold(fmt.Sprintf("%v", in))
}

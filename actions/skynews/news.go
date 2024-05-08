package skynews

import (
	lang "github.com/mt1976/crt/language"
	term "github.com/mt1976/crt/terminal"
	conf "github.com/mt1976/mockterm/config"
)

var C = conf.Configuration

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(t *term.ViewPort) {

	// Home
	// UK
	// World
	// US
	// Business
	// Politics
	// Technology
	// Entertainment
	// Strange News

	t.Clear()
	p := t.NewPage(lang.TxtMenuTitle)
	p.AddBlankRow()
	c := 0
	c++
	p.AddMenuOption(c, lang.TxtTopicHome, C.URISkyNews+C.URISkyNewsHome, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicUK, C.URISkyNews+C.URISkyNewsUK, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicWorld, C.URISkyNews+C.URISkyNewsWorld, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicUS, C.URISkyNews+C.URISkyNewsUS, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicBusiness, C.URISkyNews+C.URISkyNewsBusiness, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicPolitics, C.URISkyNews+C.URISkyNewsPolitics, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicTechnology, C.URISkyNews+C.URISkyNewsTechnology, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicEntertainment, C.URISkyNews+C.URISkyNewsEntertainment, "")
	c++
	p.AddMenuOption(c, lang.TxtTopicStrange, C.URISkyNews+C.URISkyNewsStrange, "")
	p.AddAction(lang.SymActionQuit)

	for {
		action := p.Display_Actions()

		if p.ViewPort().Formatters.Upcase(action) == lang.SymActionQuit {
			break
		}
		if t.Helpers.IsInt(action) {
			switch action {
			case "1":
				Topic(p, C.URISkyNews+C.URISkyNewsHome, lang.TxtTopicHome)
			case "2":
				Topic(p, C.URISkyNews+C.URISkyNewsUK, lang.TxtTopicUK)
			case "3":
				Topic(p, C.URISkyNews+C.URISkyNewsWorld, lang.TxtTopicWorld)
			case "4":
				Topic(p, C.URISkyNews+C.URISkyNewsUS, lang.TxtTopicUS)
			case "5":
				Topic(p, C.URISkyNews+C.URISkyNewsBusiness, lang.TxtTopicBusiness)
			case "6":
				Topic(p, C.URISkyNews+C.URISkyNewsPolitics, lang.TxtTopicPolitics)
			case "7":
				Topic(p, C.URISkyNews+C.URISkyNewsTechnology, lang.TxtTopicTechnology)
			case "8":
				Topic(p, C.URISkyNews+C.URISkyNewsEntertainment, lang.TxtTopicEntertainment)
			case "9":
				Topic(p, C.URISkyNews+C.URISkyNewsStrange, lang.TxtTopicStrange)
			default:
				p.Error(term.ErrInvalidAction, action)
			}
		}
	}
}

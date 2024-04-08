package skynews

import (
	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
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

	action, nextLevel := p.Display_Actions()

	if action == lang.SymActionQuit {
		return
	}
	if t.Helpers.IsInt(action) {
		Topic(p, nextLevel.AlternateID, nextLevel.Title)
		action = ""
	}
}

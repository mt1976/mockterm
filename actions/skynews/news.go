package skynews

import (
	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
	conf "github.com/mt1976/mockterm/config"
)

var C = conf.Configuration

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(crt *term.Crt) {

	// Home
	// UK
	// World
	// US
	// Business
	// Politics
	// Technology
	// Entertainment
	// Strange News

	crt.Clear()
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := crt.NewTitledPage(lang.TxtMenuTitle)
	c := 0
	c++
	m.AddOption(c, lang.TxtTopicHome, C.URISkyNews+C.URISkyNewsHome, "")
	c++
	m.AddOption(c, lang.TxtTopicUK, C.URISkyNews+C.URISkyNewsUK, "")
	c++
	m.AddOption(c, lang.TxtTopicWorld, C.URISkyNews+C.URISkyNewsWorld, "")
	c++
	m.AddOption(c, lang.TxtTopicUS, C.URISkyNews+C.URISkyNewsUS, "")
	c++
	m.AddOption(c, lang.TxtTopicBusiness, C.URISkyNews+C.URISkyNewsBusiness, "")
	c++
	m.AddOption(c, lang.TxtTopicPolitics, C.URISkyNews+C.URISkyNewsPolitics, "")
	c++
	m.AddOption(c, lang.TxtTopicTechnology, C.URISkyNews+C.URISkyNewsTechnology, "")
	c++
	m.AddOption(c, lang.TxtTopicEntertainment, C.URISkyNews+C.URISkyNewsEntertainment, "")
	c++
	m.AddOption(c, lang.TxtTopicStrange, C.URISkyNews+C.URISkyNewsStrange, "")
	m.AddAction(lang.SymActionQuit)

	action, nextLevel := m.Display(crt)

	if action == lang.SymActionQuit {
		return
	}
	if crt.Helpers.IsInt(action) {
		Topic(crt, nextLevel.AlternateID, nextLevel.Title)
		action = ""
	}

}

package skynews

import (
	support "github.com/mt1976/crt"
	t "github.com/mt1976/crt/language"
	"github.com/mt1976/mockterm/config"
)

var C = config.Configuration

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(crt *support.Crt) {

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
	m := support.NewPageWithName(t.TxtMenuTitle)
	c := 0
	c++
	m.AddOption(c, t.TxtTopicHome, C.URISkyNews+C.URISkyNewsHome, "")
	c++
	m.AddOption(c, t.TxtTopicUK, C.URISkyNews+C.URISkyNewsUK, "")
	c++
	m.AddOption(c, t.TxtTopicWorld, C.URISkyNews+C.URISkyNewsWorld, "")
	c++
	m.AddOption(c, t.TxtTopicUS, C.URISkyNews+C.URISkyNewsUS, "")
	c++
	m.AddOption(c, t.TxtTopicBusiness, C.URISkyNews+C.URISkyNewsBusiness, "")
	c++
	m.AddOption(c, t.TxtTopicPolitics, C.URISkyNews+C.URISkyNewsPolitics, "")
	c++
	m.AddOption(c, t.TxtTopicTechnology, C.URISkyNews+C.URISkyNewsTechnology, "")
	c++
	m.AddOption(c, t.TxtTopicEntertainment, C.URISkyNews+C.URISkyNewsEntertainment, "")
	c++
	m.AddOption(c, t.TxtTopicStrange, C.URISkyNews+C.URISkyNewsStrange, "")
	m.AddAction(t.SymActionQuit)

	action, nextLevel := m.Display(crt)

	if action == t.SymActionQuit {
		return
	}
	if support.IsInt(action) {
		Topic(crt, nextLevel.AlternateID, nextLevel.Title)
		action = ""
	}

}

package skynews

import (
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	conf "github.com/mt1976/mockterm/config"
	lang "github.com/mt1976/mockterm/language"
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
	p := page.NewPage(t, lang.SkyNewsTitle.Text())
	p.AddBlankRow()
	c := 0
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicHome.Text(), C.URISkyNews+C.URISkyNewsHome, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicUK.Text(), C.URISkyNews+C.URISkyNewsUK, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicWorld.Text(), C.URISkyNews+C.URISkyNewsWorld, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicUS.Text(), C.URISkyNews+C.URISkyNewsUS, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicBusiness.Text(), C.URISkyNews+C.URISkyNewsBusiness, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicPolitics.Text(), C.URISkyNews+C.URISkyNewsPolitics, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicTechnology.Text(), C.URISkyNews+C.URISkyNewsTechnology, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicEntertainment.Text(), C.URISkyNews+C.URISkyNewsEntertainment, "")
	c++
	p.AddMenuOption(c, lang.SkyNewsTopicStrange.Text(), C.URISkyNews+C.URISkyNewsStrange, "")
	p.AddAction(acts.Quit)

	for {
		action := p.Display_Actions()

		if action.Is(acts.Quit) {
			break
		}
		if action.IsInt() {
			switch action.Action() {
			case "1":
				Topic(p, C.URISkyNews+C.URISkyNewsHome, lang.SkyNewsTopicHome.Text())
			case "2":
				Topic(p, C.URISkyNews+C.URISkyNewsUK, lang.SkyNewsTopicUK.Text())
			case "3":
				Topic(p, C.URISkyNews+C.URISkyNewsWorld, lang.SkyNewsTopicWorld.Text())
			case "4":
				Topic(p, C.URISkyNews+C.URISkyNewsUS, lang.SkyNewsTopicUS.Text())
			case "5":
				Topic(p, C.URISkyNews+C.URISkyNewsBusiness, lang.SkyNewsTopicBusiness.Text())
			case "6":
				Topic(p, C.URISkyNews+C.URISkyNewsPolitics, lang.SkyNewsTopicPolitics.Text())
			case "7":
				Topic(p, C.URISkyNews+C.URISkyNewsTechnology, lang.SkyNewsTopicTechnology.Text())
			case "8":
				Topic(p, C.URISkyNews+C.URISkyNewsEntertainment, lang.SkyNewsTopicEntertainment.Text())
			case "9":
				Topic(p, C.URISkyNews+C.URISkyNewsStrange, lang.SkyNewsTopicStrange.Text())
			default:
				p.Error(terr.ErrInvalidAction, action.Action())
			}
		}
	}
}

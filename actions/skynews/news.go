package skynews

import (
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	action "github.com/mt1976/crt/page/actions"
	"github.com/mt1976/crt/terminal"
	"github.com/mt1976/mockterm/config"
	l "github.com/mt1976/mockterm/language"
)

var CFG = config.Configuration

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(t *terminal.ViewPort) {

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
	p := page.NewPage(t, l.SkyNewsTitle)
	p.AddBlankRow()
	c := 0
	c++
	p.AddMenuOption(c, l.SkyNewsTopicHome.Text(), CFG.URISkyNews+CFG.URISkyNewsHome, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicUK.Text(), CFG.URISkyNews+CFG.URISkyNewsUK, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicWorld.Text(), CFG.URISkyNews+CFG.URISkyNewsWorld, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicUS.Text(), CFG.URISkyNews+CFG.URISkyNewsUS, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicBusiness.Text(), CFG.URISkyNews+CFG.URISkyNewsBusiness, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicPolitics.Text(), CFG.URISkyNews+CFG.URISkyNewsPolitics, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicTechnology.Text(), CFG.URISkyNews+CFG.URISkyNewsTechnology, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicEntertainment.Text(), CFG.URISkyNews+CFG.URISkyNewsEntertainment, "")
	c++
	p.AddMenuOption(c, l.SkyNewsTopicStrange.Text(), CFG.URISkyNews+CFG.URISkyNewsStrange, "")
	p.AddAction(action.Quit)

	for {
		next := p.Display_Actions()

		if next.Is(action.Quit) {
			break
		}
		if next.IsInt() {
			switch next.Action() {
			case "1":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsHome, l.SkyNewsTopicHome.Text())
			case "2":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsUK, l.SkyNewsTopicUK.Text())
			case "3":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsWorld, l.SkyNewsTopicWorld.Text())
			case "4":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsUS, l.SkyNewsTopicUS.Text())
			case "5":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsBusiness, l.SkyNewsTopicBusiness.Text())
			case "6":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsPolitics, l.SkyNewsTopicPolitics.Text())
			case "7":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsTechnology, l.SkyNewsTopicTechnology.Text())
			case "8":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsEntertainment, l.SkyNewsTopicEntertainment.Text())
			case "9":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsStrange, l.SkyNewsTopicStrange.Text())
			default:
				p.Error(terr.ErrInvalidAction, next.Action())
			}
		}
	}
}

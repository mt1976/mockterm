package skynews

import (
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	action "github.com/mt1976/crt/page/actions"
	"github.com/mt1976/crt/terminal"
	l "github.com/mt1976/mockterm/actions/skynews/lang"
	"github.com/mt1976/mockterm/config"
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
	p := page.NewPage(t, l.Title)
	p.AddBlankRow()
	c := 0
	c++
	p.AddMenuOption(c, l.Home.Text(), CFG.URISkyNews+CFG.URISkyNewsHome, "")
	c++
	p.AddMenuOption(c, l.UK.Text(), CFG.URISkyNews+CFG.URISkyNewsUK, "")
	c++
	p.AddMenuOption(c, l.World.Text(), CFG.URISkyNews+CFG.URISkyNewsWorld, "")
	c++
	p.AddMenuOption(c, l.US.Text(), CFG.URISkyNews+CFG.URISkyNewsUS, "")
	c++
	p.AddMenuOption(c, l.Business.Text(), CFG.URISkyNews+CFG.URISkyNewsBusiness, "")
	c++
	p.AddMenuOption(c, l.Politics.Text(), CFG.URISkyNews+CFG.URISkyNewsPolitics, "")
	c++
	p.AddMenuOption(c, l.Technology.Text(), CFG.URISkyNews+CFG.URISkyNewsTechnology, "")
	c++
	p.AddMenuOption(c, l.Entertainment.Text(), CFG.URISkyNews+CFG.URISkyNewsEntertainment, "")
	c++
	p.AddMenuOption(c, l.Strange.Text(), CFG.URISkyNews+CFG.URISkyNewsStrange, "")
	p.AddAction(action.Quit)

	for {
		next := p.Display_Actions()

		if next.Is(action.Quit) {
			break
		}
		if next.IsInt() {
			switch next.Action() {
			case "1":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsHome, l.Home.Text())
			case "2":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsUK, l.UK.Text())
			case "3":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsWorld, l.World.Text())
			case "4":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsUS, l.US.Text())
			case "5":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsBusiness, l.Business.Text())
			case "6":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsPolitics, l.Politics.Text())
			case "7":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsTechnology, l.Technology.Text())
			case "8":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsEntertainment, l.Entertainment.Text())
			case "9":
				Topic(p, CFG.URISkyNews+CFG.URISkyNewsStrange, l.Strange.Text())
			default:
				p.Error(terr.ErrInvalidAction, next.Action())
			}
		}
	}
}

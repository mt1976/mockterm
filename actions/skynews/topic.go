package skynews

import (
	"github.com/mmcdole/gofeed"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	lang "github.com/mt1976/mockterm/language"
)

// The function "Topic" takes in a ViewPort object, a topic, and a title as parameters, and then retrieves
// news items for that topic from an RSS feed, displays them in a menu, and allows the user to select a
// news item to view.
func Topic(t *page.Page, topic, title string) {

	vp := t.ViewPort()
	// Get the news for the topic
	t.Info(lang.TxtLoadingTopic + title)
	// get the news for the topic from an rss feed
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(topic)
	t.Clear()

	p := page.NewPage(&vp, feed.Title)

	p.AddBlankRow()
	noNewsItems := len(feed.Items)
	if noNewsItems > C.MaxContentRows {
		noNewsItems = C.MaxContentRows
	}

	for i := range noNewsItems {
		//log.Println("Adding: ", feed.Items[i].Title, i)
		dt := vp.Formatters.TimeAgo(feed.Items[i].Published)
		p.AddMenuOption(i+1, feed.Items[i].Title, feed.Items[i].Link, dt)
		i++
	}

	for {
		action := p.Display_Actions()
		if action == acts.Quit {
			break
		}
		if action.IsInt() {
			i := (vp.Helpers.ToInt(action.Action()) - 1)
			p.Dump(string(i), string(vp.Helpers.ToInt(action.Action())), action.Action())
			//os.Exit(1)
			Story(p, feed.Items[i].Link, feed.Items[i].Title)
			//action = ""
		}
	}
}

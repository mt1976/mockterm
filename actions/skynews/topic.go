package skynews

import (
	"github.com/mmcdole/gofeed"
	support "github.com/mt1976/crt"
	t "github.com/mt1976/crt/language"
)

// The function "Topic" takes in a CRT object, a topic, and a title as parameters, and then retrieves
// news items for that topic from an RSS feed, displays them in a menu, and allows the user to select a
// news item to view.
func Topic(crt *support.Crt, topic, title string) {

	// Get the news for the topic
	crt.InfoMessage(t.TxtLoadingTopic + crt.Bold(title))
	// get the news for the topic from an rss feed
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(topic)
	crt.Clear()

	p := support.NewPageWithName(feed.Title)
	noNewsItems := len(feed.Items)
	if noNewsItems > C.MaxContentRows {
		noNewsItems = C.MaxContentRows
	}

	for i := range noNewsItems {
		//log.Println("Adding: ", feed.Items[i].Title, i)
		dt := support.TimeAgo(feed.Items[i].Published)
		p.AddOption(i+1, feed.Items[i].Title, feed.Items[i].Link, dt)
		i++
	}

	action, mi := p.Display(crt)

	if action == t.SymActionQuit {
		//crt.Println("Quitting")
		return
	}
	if support.IsInt(action) {
		Story(crt, mi.AlternateID)
		action = ""
	}
}

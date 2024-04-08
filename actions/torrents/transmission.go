package skynews

import (
	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
)

// The function "Trans" takes in a ViewPort object, a topic, and a title as parameters, and then retrieves
// news items for that topic from an RSS feed, displays them in a menu, and allows the user to select a
// news item to view.
func Trans(t *term.ViewPort, uri, title string) {
	// Get the news for the topic
	t.InfoMessage(lang.TxtLoadingTorrentsTransmission)
	p := t.NewPage(lang.TxtTransmission)
	// get the news for the topic from an rss feed
	// endpoint, err := url.Parse(uri)
	// if err != nil {
	// 	panic(err)
	// }
	//tbt, err := transmissionrpc.New(endpoint.Host, "admin", "admin", nil)
	//if err != nil {
	//	panic(err)
	//}
	//spew.Dump(tbt)
	//spew.Dump(tbt.RPCVersion())
	//spew.Dump(tbt.TorrentGetAll())
	//os.Exit(0)
	// torrents, _ := tbt.TorrentGetAll()
	// noTorrents := len(torrents)
	// if noTorrents > page.MaxMenuItems {
	// 	noTorrents = page.MaxMenuItems
	// }

	// for i := range noTorrents {
	// 	//log.Println("Adding: ", feed.Items[i].Title, i)
	// 	p.Add(i+1, torrents[i].Name, "", "")
	// }
	ok := false
	for !ok {
		action, _ := p.Display_Actions()

		if action == lang.SymActionQuit {
			ok = true
			continue
		}
		if t.Helpers.IsInt(action) {
			ok = false
			action = ""
		}

		//log.Println("Action: ", action)
		//log.Println("Next Level: ", mi)

		//spew.Dump(nextLevel)
	}
}

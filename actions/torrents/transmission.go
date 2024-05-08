package skynews

import (
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
)

// The function "Trans" takes in a ViewPort object, a topic, and a title as parameters, and then retrieves
// news items for that topic from an RSS feed, displays them in a menu, and allows the user to select a
// news item to view.
func Trans(t *term.ViewPort, uri, title string) {
	// Get the news for the topic
	t.InfoMessage("TxtLoadingTorrentsTransmission")
	p := page.NewPage(t, "TxtTransmission")
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
		action := p.Display_Actions()

		if action == acts.Quit {
			ok = true
			continue
		}
		if action.IsInt() {
			ok = false
			return
		}

		//log.Println("Action: ", action)
		//log.Println("Next Level: ", mi)

		//spew.Dump(nextLevel)
	}
}

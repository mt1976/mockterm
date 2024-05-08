package showsrss

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Detail(t *term.ViewPort, item RssItem) error {

	t.Clear()
	p := page.NewPage(t, item.Title)
	p.AddBlankRow()
	//c := 0

	layout := "Mon, 2 Jan 2006 15:04:05 +0000"
	vp := p.ViewPort()

	// for _, item := range rss.Channel.Item {
	// 	c++

	dt, err := time.Parse(layout, item.PubDate)
	if err != nil {
		p.Error(err, "Error parsing date")
	}

	d := vp.Formatters.HumanFromUnixDate(dt.Unix())
	//dts := vp.Formatters.FormatDate(dt)
	// 	desc := item.Title

	// 	sw := vp.Width() - (len(d) + 10)

	// 	xy := "%-" + vp.Helpers.ToString(sw) + "s"
	// 	desc = fmt.Sprintf(xy, desc)

	// 	desc = desc + d

	// 	p.AddMenuOption(c, desc, item.Link, d)
	// }
	p.AddFieldValuePair("ShowName", item.ShowName)
	p.AddFieldValuePair("Title", item.Title)
	p.AddFieldValuePair("Date", d)
	p.AddBreakRow()
	p.AddFieldValuePair("Show ID", item.ShowID)
	p.AddFieldValuePair("Episode ID", item.EpisodeID)
	p.AddFieldValuePair("RawTitle", item.RawTitle)
	p.AddFieldValuePair("Description", item.Description)
	p.AddFieldValuePair("Published", item.PubDate)
	p.AddBreakRow()

	p.AddFieldValuePair("Link", item.Link)
	//p.AddFieldValuePair("Guid", item.Guid.Text)

	// p.AddFieldValuePair("ExternalID", item.ExternalID)
	// p.AddFieldValuePair("InfoHash", item.InfoHash)
	// p.AddBreakRow()
	// p.AddFieldValuePair("Enclosure", item.Enclosure.Text)
	// p.AddFieldValuePair("Enclosure URL", item.Enclosure.URL)
	// p.AddFieldValuePair("Enclosure Length", item.Enclosure.Length)
	// p.AddFieldValuePair("Enclosure Type", item.Enclosure.Type)
	// p.AddBreakRow()
	// p.AddFieldValuePair("Guid.IsPermaLink", item.Guid.IsPermaLink)
	// p.AddFieldValuePair("Guid.Text", item.Guid.Text)
	// p.AddBreakRow()
	//c++
	//p.AddMenuOption(c, lang.TxtTopicHome, C.URISkyNews+C.URISkyNewsHome, "")

	p.AddAction(acts.Quit)
	addTorrentAction := acts.New("A")
	p.AddAction(acts.New("A"))
	p.SetPrompt("Choose (A)dd Torrent, (F)orward, (B)ack or (Q)uit")

	for {
		action := p.Display_Actions()

		if action.Is(acts.Quit) {
			break
		}
		switch action {
		case acts.Quit:
			return nil
		case addTorrentAction:
			cmd := exec.Command("open", item.Link)
			// cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
			stderr, err := cmd.StderrPipe()
			log.SetOutput(os.Stderr)

			if err != nil {
				log.Fatal(err)
			}
			if err := cmd.Start(); err != nil {
				log.Fatal(err)
			}

			slurp, _ := io.ReadAll(stderr)
			fmt.Printf("%s\n", slurp)

			if err := cmd.Wait(); err != nil {
				log.Fatal(err)
			}
		default:
			p.Error(terr.ErrInvalidAction, action.Action())
		}
	}
	return nil
}

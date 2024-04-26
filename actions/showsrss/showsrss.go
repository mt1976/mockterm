package showsrss

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
	conf "github.com/mt1976/mockterm/config"
)

var C = conf.Configuration

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(t *term.ViewPort) error {

	url := "https://showrss.info/user/282183.rss?magnets=true&namespaces=true&name=null&quality=null&re=null"
	rss, err := fetchXML(url)
	if err != nil {
		fmt.Println("Error fetching XML:", err)
		return err
	}

	t.Clear()
	p := t.NewPage("Shows RSS Feed")
	p.AddBlankRow()
	c := 0

	layout := "Mon, 2 Jan 2006 15:04:05 +0000"
	vp := p.ViewPort()

	for _, item := range rss.Channel.Item {
		c++

		dt, err := time.Parse(layout, item.PubDate)
		if err != nil {
			p.Error(err, "Error parsing date")
		}

		d := vp.Formatters.HumanFromUnixDate(dt.Unix())

		desc := item.Title

		sw := vp.Width() - (len(d) + 10)

		xy := "%-" + vp.Helpers.ToString(sw) + "s"
		desc = fmt.Sprintf(xy, desc)

		desc = desc + d

		p.AddMenuOption(c, desc, item.Link, d)
	}

	//c++
	//p.AddMenuOption(c, lang.TxtTopicHome, C.URISkyNews+C.URISkyNewsHome, "")

	p.AddAction(lang.SymActionQuit)

	p.Dump("Shows RSS Feed", spew.Sdump(rss))

	for {
		action := p.Display_Actions()

		if p.ViewPort().Formatters.Upcase(action) == lang.SymActionQuit {
			break
		}
		if t.Helpers.IsInt(action) {
			switch {
			case t.Helpers.IsInt(action):
				i := t.Helpers.ToInt(action)
				item := rss.Channel.Item[i-1]
				Detail(t, item)

			default:
				p.Error(term.ErrInvalidAction, action)
			}
		}
	}
	return nil
}

func fetchXML(url string) (*RssBody, error) {
	// Fetch the XML document
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal XML data
	var rss RssBody
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}

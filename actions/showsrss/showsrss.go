package showsrss

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/actions/showsrss/lang"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
)

var C = conf.Configuration

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(t *term.ViewPort) error {

	url := "https://showrss.info/user/282183.rss?magnets=true&namespaces=true&name=null&quality=null&re=null"
	rss, err := fetchXML(url)
	if err != nil {
		fmt.Println(errs.FetchingXML, err)
		return err
	}

	t.Clear()
	p := page.NewPage(t, lang.ShowsRssPageTitle)
	p.AddBlankRow()
	c := 0

	layout := "Mon, 2 Jan 2006 15:04:05 +0000"
	vp := p.ViewPort()

	for _, item := range rss.Channel.Item {
		c++

		dt, err := time.Parse(layout, item.PubDate)
		if err != nil {
			p.Error(errs.ParsingDate, err.Error())
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

	p.AddAction(acts.Quit)

	p.Dump("Shows RSS Feed", spew.Sdump(rss))

	for {
		action := p.Display_Actions()

		if action.Is(acts.Quit) {
			break
		}
		if action.IsInt() {
			switch {
			case action.IsInt():
				i := t.Helpers.ToInt(action.Action())
				item := rss.Channel.Item[i-1]
				Detail(t, item)

			default:
				p.Error(terr.ErrInvalidAction, action.Action())
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

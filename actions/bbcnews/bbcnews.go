package bbcnews

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	page "github.com/mt1976/crt/page"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/actions/bbcnews/lang"
)

func Run(t *term.ViewPort) (err error) {

	page := page.NewPage(t, lang.Title)

	uri := "https://feeds.bbci.co.uk/news/rss.xml"

	xmlContent, err := getXML(uri)
	if err != nil {
		t.Error(err, "")
		return err
	}

	var rss BBCNewsRss
	err = xml.Unmarshal([]byte(xmlContent), &rss)
	if err != nil {
		t.Error(err, "")
		return err
	}

	yy := rss.Channel

	count := 0
	for _, item := range yy.Item {
		count++
		page.AddMenuOption(count, item.Title, item.Link, item.PubDate)
	}

	na := page.Display_Actions()
	if na.IsInt() {
		op := yy.Item[t.Helpers.ToInt(na.Action())-1]
		spew.Dump(op)
		os.Exit(0)
	}

	return nil
}

func getXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("bad status: %v", resp.StatusCode))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

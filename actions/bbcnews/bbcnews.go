package bbcnews

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	term "github.com/mt1976/crt"
)

func Run(t *term.ViewPort) (err error) {

	page := t.NewPage("BBC News")

	uri := "https://feeds.bbci.co.uk/news/rss.xml"

	xmlContent, err := getXML(uri)
	if err != nil {
		t.Error(err, "")
		return err
	}
	//fmt.Printf("xml: %v\n", xmlContent)
	var rss BBCNewsRss
	err = xml.Unmarshal([]byte(xmlContent), &rss)
	if err != nil {
		t.Error(err, "")
		return err
	}
	//fmt.Printf("rss: %v\n", rss)
	//page.Add("content", "", "")

	yy := rss.Channel
	//fmt.Printf("yy: %v\n", yy)

	count := 0
	for _, item := range yy.Item {
		//	fmt.Printf("item: %v\n", item.Title)
		//id := ""
		//	yylink := item.Link
		//fmt.Printf("yylink: %+v\n", yylink)
		count++
		page.AddMenuOption(count, item.Title, item.Link, item.PubDate)
	}

	na := page.Display_Actions()
	if t.Helpers.IsInt(na) {
		op := yy.Item[t.Helpers.ToInt(na)-1]
		spew.Dump(op)
		os.Exit(0)
	}

	//page.Info(fmt.Sprintf("na: %+v\n", na))
	//page.Info(fmt.Sprintf("selected: %+v\n", selected))
	//fmt.Printf("na: %+v\n", na)
	//fmt.Printf("selected: %+v\n", selected)

	return nil
}

func getXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

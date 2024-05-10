package skynews

import (
	"github.com/gocolly/colly"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	lang "github.com/mt1976/mockterm/language"
)

// The function "Story" displays a story link and allows the user to interact with a menu until they
// choose to quit.
func Story(p *page.Page, storyLink, title string) {

	t := p.ViewPort()
	np := page.NewPage(&t, "")
	np.Clear()
	np.AddFieldValuePair(lang.Title.Text(), title)
	np.AddFieldValuePair(lang.SkyNewsStory.Text(), storyLink)
	np.AddBreakRow()
	//np.AddBlankRow()
	np = buildPage(np, storyLink)
	np.SetTitle(title)
	np.Info(lang.SkyNewsLoadingStory.Text())
	np.ActivePageIndex = 0

	for {
		x := np.Display_Actions()

		if x.Is(acts.Quit) {
			return
		}
	}
}

// buildPage creates a new page with the given title and adds a link to the given story to the page.
// It uses the colly library to fetch the story content and extract the title.
func buildPage(p *page.Page, storyLink string) *page.Page {
	// Get html from storyLink
	// Parse html for story
	// Create page with story
	// Return page

	// Create a new collector
	c := colly.NewCollector()

	// Store the page title
	var pageTitle string

	// Find and visit all links
	c.OnHTML(lang.SkyNewsHTMLTagTitle, func(e *colly.HTMLElement) {
		pageTitle = e.Text
	})

	// Store the story content
	var storyContent []string

	// Parse the story content
	c.OnHTML(lang.SkyNewsHTMLTagP, func(e *colly.HTMLElement) {
		storyContent = append(storyContent, e.Text)
	})

	// Visit the story link
	c.Visit(storyLink)

	// Create a new page with the title
	p.SetTitle(pageTitle)
	//p.AddBlankRow()

	// Add the story content to the page
	for _, content := range storyContent {
		p.Add(content, "", "")
	}

	return p
}

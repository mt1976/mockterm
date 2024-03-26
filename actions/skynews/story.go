package skynews

import (
	"github.com/gocolly/colly"
	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
)

// The function "Story" displays a story link and allows the user to interact with a menu until they
// choose to quit.
func Story(t *term.ViewPort, storyLink string) {

	p := buildPage(t, storyLink)
	p.Info(lang.TxtLoadingStory)
	p.ActivePageIndex = 0

	x, _ := p.DisplayWithActions()

	if x == lang.SymActionQuit {
		return
	}

}

// buildPage creates a new page with the given title and adds a link to the given story to the page.
// It uses the colly library to fetch the story content and extract the title.
func buildPage(t *term.ViewPort, storyLink string) *term.Page {
	// Get html from storyLink
	// Parse html for story
	// Create page with story
	// Return page

	// Create a new collector
	c := colly.NewCollector()

	// Store the page title
	var pageTitle string

	// Find and visit all links
	c.OnHTML(lang.HTMLTagTitle, func(e *colly.HTMLElement) {
		pageTitle = e.Text
	})

	// Store the story content
	var storyContent []string

	// Parse the story content
	c.OnHTML(lang.HTMLTagTagP, func(e *colly.HTMLElement) {
		storyContent = append(storyContent, e.Text)
	})

	// Visit the story link
	c.Visit(storyLink)

	// Create a new page with the title
	p := t.NewPage(pageTitle)

	// Add the story content to the page
	for _, content := range storyContent {
		p.Add(content, "", "")
	}

	return p
}

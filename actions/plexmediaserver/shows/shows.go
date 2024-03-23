package shows

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	term "github.com/mt1976/crt"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	pmss "github.com/mt1976/mockterm/plexsupport"
)

var C = conf.Configuration

func Run(t *term.ViewPort, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		t.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := t.NewTitledPage(res.MediaContainer.LibrarySectionTitle + lang.Space + t.Formatters.PQuote(noItems))
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.AddOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	nextAction, _ := m.DisplayWithActions()
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		if t.Helpers.IsInt(nextAction) {
			Detail(t, res.MediaContainer.Metadata[t.Helpers.ToInt(nextAction)-1], mediaVault)
		} else {
			m.Error(term.ErrInvalidAction, t.Formatters.SQuote(nextAction))
		}
	}
}

func Detail(t *term.ViewPort, info plex.Metadata, mediaVault *plex.Plex) {
	p := t.NewTitledPage(info.Title)

	p.AddFieldValuePair(lang.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(lang.TxtYear, t.Helpers.ToString(info.Year))
	p.AddFieldValuePair(lang.TxtPlexContentRatingLabel, info.ContentRating)
	p.AddFieldValuePair(lang.TxtPlexReleasedLabel, pmss.FormatPlexDate(info.OriginallyAvailableAt))
	p.BlankRow()
	p.AddFieldValuePair(lang.TxtPlexSummaryLabel, info.Summary)

	p.AddAction(lang.SymActionSeasons) //Drilldown to episodes
	p.SetPrompt(lang.TxtPlexSeasonsPrompt)

	nextAction, _ := p.DisplayWithActions()
	switch nextAction {
	case lang.SymActionQuit:
		return
	case lang.SymActionSeasons:
		SeasonDetails(t, mediaVault, info)
	}
}

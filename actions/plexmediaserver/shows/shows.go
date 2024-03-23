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

func Run(crt *term.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := crt.NewTitledPage(res.MediaContainer.LibrarySectionTitle + lang.Space + crt.Formatters.PQuote(noItems))
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.AddOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	nextAction, _ := m.DisplayWithActions(crt)
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		if crt.Helpers.IsInt(nextAction) {
			Detail(crt, res.MediaContainer.Metadata[crt.Helpers.ToInt(nextAction)-1], mediaVault)
		} else {
			m.Error(crt, term.ErrInvalidAction, crt.Formatters.SQuote(nextAction))
		}
	}
}

func Detail(crt *term.Crt, info plex.Metadata, mediaVault *plex.Plex) {
	p := crt.NewTitledPage(info.Title)

	p.AddFieldValuePair(crt, lang.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(crt, lang.TxtYear, crt.Helpers.ToString(info.Year))
	p.AddFieldValuePair(crt, lang.TxtPlexContentRatingLabel, info.ContentRating)
	p.AddFieldValuePair(crt, lang.TxtPlexReleasedLabel, pmss.FormatPlexDate(info.OriginallyAvailableAt))
	p.BlankRow()
	p.AddFieldValuePair(crt, lang.TxtPlexSummaryLabel, info.Summary)

	p.AddAction(lang.SymActionSeasons) //Drilldown to episodes
	p.SetPrompt(lang.TxtPlexSeasonsPrompt)

	nextAction, _ := p.DisplayWithActions(crt)
	switch nextAction {
	case lang.SymActionQuit:
		return
	case lang.SymActionSeasons:
		SeasonDetails(crt, mediaVault, info)
	}
}

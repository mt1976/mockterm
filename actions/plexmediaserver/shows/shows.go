package shows

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	support "github.com/mt1976/crt"
	"github.com/mt1976/mockterm/config"
	e "github.com/mt1976/mockterm/errors"
	t "github.com/mt1976/mockterm/language"
	plexSupport "github.com/mt1976/mockterm/plexsupport"
)

var C = config.Configuration

func Run(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(e.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := support.NewPageWithName(res.MediaContainer.LibrarySectionTitle + t.Space + support.PQuote(noItems))
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.AddOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	nextAction, _ := m.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		if support.IsInt(nextAction) {
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1], mediaVault)
		} else {
			crt.InputError(support.ErrInvalidAction, support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata, mediaVault *plex.Plex) {
	p := support.NewPageWithName(info.Title)

	p.AddFieldValuePair(crt, t.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(crt, t.TxtYear, support.ToString(info.Year))
	p.AddFieldValuePair(crt, t.TxtPlexContentRatingLabel, info.ContentRating)
	p.AddFieldValuePair(crt, t.TxtPlexReleasedLabel, plexSupport.FormatPlexDate(info.OriginallyAvailableAt))
	p.BlankRow()
	p.AddFieldValuePair(crt, t.TxtPlexSummaryLabel, info.Summary)

	p.AddAction(t.SymActionSeasons) //Drilldown to episodes
	p.SetPrompt(t.TxtPlexSeasonsPrompt)

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	case t.SymActionSeasons:
		SeasonDetails(crt, mediaVault, info)
	}
}

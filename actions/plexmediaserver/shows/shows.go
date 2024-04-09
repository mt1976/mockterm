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

	p := t.NewPage(res.MediaContainer.LibrarySectionTitle + lang.Space + t.Formatters.PQuote(noItems))
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		p.AddMenuOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	for {
		nextAction, _ := p.Display_Actions()
		switch {
		case t.Formatters.Upcase(nextAction) == lang.SymActionQuit:
			return
		case t.Helpers.IsInt(nextAction):
			Detail(t, res.MediaContainer.Metadata[t.Helpers.ToInt(nextAction)-1], mediaVault)
		default:
			p.Error(term.ErrInvalidAction, t.Formatters.SQuote(nextAction))
		}
	}
}

func Detail(t *term.ViewPort, info plex.Metadata, mediaVault *plex.Plex) {
	p := t.NewPage(info.Title)

	p.AddFieldValuePair(lang.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(lang.TxtYear, t.Helpers.ToString(info.Year))
	p.AddFieldValuePair(lang.TxtPlexContentRatingLabel, info.ContentRating)
	p.AddFieldValuePair(lang.TxtPlexReleasedLabel, pmss.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddBlankRow()
	p.AddFieldValuePair(lang.TxtPlexSummaryLabel, info.Summary)

	p.AddAction(lang.SymActionSeasons) //Drilldown to episodes
	p.SetPrompt(lang.TxtPlexSeasonsPrompt)
	p.AddBlankRow()
	yy, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		t.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}
	//p := t.NewPage(lang.TxtPlexSeasons + info.Title)
	noResps := len(yy.MediaContainer.Metadata)
	for i := 0; i < noResps; i++ {
		season := yy.MediaContainer.Metadata[i]
		p.AddMenuOption(i+1, season.Title, "", "")
	}

	for {
		nextAction, _ := p.Display_Actions()
		switch {
		case t.Formatters.Upcase(nextAction) == lang.SymActionQuit:
			return
		case t.Helpers.IsInt(nextAction):
			Episodes(t, mediaVault, info.Title, yy.MediaContainer.Metadata[t.Helpers.ToInt(nextAction)-1])
			//	case t.Formatters.Upcase(nextAction) == lang.SymActionSeasons:
			//		SeasonDetails(t, mediaVault, info)
			//	}
		default:
			p.Error(term.ErrInvalidAction, t.Formatters.SQuote(nextAction))
		}
	}
}

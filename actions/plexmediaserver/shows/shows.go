package shows

import (
	"fmt"
	"os"

	plexms "github.com/jrudio/go-plex-client"
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/actions/plexmediaserver/lang"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	clng "github.com/mt1976/mockterm/language"
	plex "github.com/mt1976/mockterm/plexsupport"
)

var C = conf.Configuration

func Run(t *term.ViewPort, mediaVault *plexms.Plex, wi *plexms.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		t.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	title := clng.New(res.MediaContainer.LibrarySectionTitle + clng.Space.Text() + t.Formatters.PQuote(noItems))
	p := page.NewPage(t, title)
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		p.AddMenuOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	for {
		nextAction := p.Display_Actions()
		switch {
		case nextAction.Is(acts.Quit):
			return
		case nextAction.IsInt():
			Detail(t, res.MediaContainer.Metadata[t.Helpers.ToInt(nextAction.Action())-1], mediaVault)
		default:
			p.Error(terr.ErrInvalidAction, t.Formatters.SQuote(nextAction.Action()))
		}
	}
}

func Detail(t *term.ViewPort, info plexms.Metadata, mediaVault *plexms.Plex) {
	p := page.NewPage(t, clng.New(info.Title))

	p.AddFieldValuePair(lang.Title.Text(), info.Title)
	p.AddFieldValuePair(lang.Year.Text(), t.Helpers.ToString(info.Year))
	p.AddFieldValuePair(lang.ContentRating.Text(), info.ContentRating)
	p.AddFieldValuePair(lang.Released.Text(), plex.FormatDate(info.OriginallyAvailableAt))
	p.AddBlankRow()
	p.AddFieldValuePair(lang.Summary, "")
	p.AddParagraphString(info.Summary)

	//p.AddAction(acts.Seasons) //Drilldown to episodes
	p.SetPrompt(lang.SeasonsPrompt)
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
		nextAction := p.Display_Actions()
		switch {
		case nextAction.Is(acts.Quit):
			return
		case nextAction.IsInt():
			Episodes(t, mediaVault, info.Title, yy.MediaContainer.Metadata[t.Helpers.ToInt(nextAction.Action())-1])
			//	case t.Formatters.Upcase(nextAction) == acts.Seasons:
			//		SeasonDetails(t, mediaVault, info)
			//	}
		default:
			p.Error(terr.ErrInvalidAction, t.Formatters.SQuote(nextAction.Action()))
		}
	}
}

package shows

import (
	"os"

	"github.com/jrudio/go-plex-client"
	term "github.com/mt1976/crt"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

func SeasonDetails(crt *term.ViewPort, mediaVault *plex.Plex, info plex.Metadata) {

	yy, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}
	p := crt.NewTitledPage(lang.TxtPlexSeasons + info.Title)
	noResps := len(yy.MediaContainer.Metadata)
	for i := 0; i < noResps; i++ {
		season := yy.MediaContainer.Metadata[i]
		p.AddOption(i+1, season.Title, "", "")
	}

	na, _ := p.DisplayWithActions()
	switch na {
	case lang.SymActionQuit:
		return
	default:
		if crt.Helpers.IsInt(na) {
			Episodes(crt, mediaVault, info.Title, yy.MediaContainer.Metadata[crt.Helpers.ToInt(na)-1])
		} else {
			p.Error(term.ErrInvalidAction, crt.Formatters.SQuote(na))
		}
	}
}

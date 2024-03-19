package shows

import (
	"os"

	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"
	page "github.com/mt1976/crt/support/page"
	e "github.com/mt1976/mockterm/errors"
	notations "github.com/mt1976/mockterm/language"
	t "github.com/mt1976/mockterm/language"
)

func SeasonDetails(crt *support.Crt, mediaVault *plex.Plex, info plex.Metadata) {

	yy, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error(e.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}
	p := page.New(notations.TxtPlexSeasons + info.Title)
	noResps := len(yy.MediaContainer.Metadata)
	for i := 0; i < noResps; i++ {
		season := yy.MediaContainer.Metadata[i]
		p.AddOption(i+1, season.Title, "", "")
	}

	na, _ := p.Display(crt)
	switch na {
	case t.SymActionQuit:
		return
	default:
		if support.IsInt(na) {
			Episodes(crt, mediaVault, info.Title, yy.MediaContainer.Metadata[support.ToInt(na)-1])
		} else {
			crt.InputError(support.ErrInvalidAction, support.SQuote(na))
		}
	}
}

package movies

import (
	"fmt"
	"os"

	plexms "github.com/jrudio/go-plex-client"
	terr "github.com/mt1976/crt/errors"
	term "github.com/mt1976/crt/terminal"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
)

func Run(t *term.ViewPort, mediaVault *plexms.Plex, wi *plexms.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		t.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	p := t.NewPage(res.MediaContainer.LibrarySectionTitle + lang.Space + t.Formatters.PQuote(noItems))
	p.Clear()

	count := 0

	for range res.MediaContainer.Metadata {
		count++
		p.AddMenuOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	for {
		nextAction := p.Display_Actions()
		switch {
		case t.Formatters.Upcase(nextAction) == lang.SymActionQuit:
			return
		case t.Helpers.IsInt(nextAction):
			Detail(t, res.MediaContainer.Metadata[t.Helpers.ToInt(nextAction)-1])
		default:
			p.Error(terr.ErrInvalidAction, t.Formatters.SQuote(nextAction))
		}
	}
}

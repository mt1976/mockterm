package music

import (
	"fmt"
	"os"

	plexms "github.com/jrudio/go-plex-client"

	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
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

	p := page.NewPage(t, res.MediaContainer.LibrarySectionTitle+lang.Space+t.Formatters.PQuote(noItems))
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
			Detail(t, res.MediaContainer.Metadata[t.Helpers.ToInt(nextAction.Action())-1])
		default:
			p.Error(terr.ErrInvalidAction, t.Formatters.SQuote(nextAction.Action()))
		}
	}
}

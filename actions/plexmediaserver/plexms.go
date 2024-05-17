package plexmediaserver

import (
	"os"
	"strconv"

	"github.com/jrudio/go-plex-client"
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	symb "github.com/mt1976/crt/strings/symbols"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/actions/plexmediaserver/lang"
	movies "github.com/mt1976/mockterm/actions/plexmediaserver/movies"
	music "github.com/mt1976/mockterm/actions/plexmediaserver/music"
	shows "github.com/mt1976/mockterm/actions/plexmediaserver/shows"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
)

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(t *term.ViewPort) {

	t.Clear()

	//spew.Dump(cfg.Configuration)
	//os.Exit(1)

	plexConnection, err := plex.New(conf.Configuration.PlexURI+":"+conf.Configuration.PlexPort, conf.Configuration.PlexToken)
	if err != nil {
		t.Error(errs.ErrPlexInit, err.Error())
		os.Exit(1)
	}

	// Test your connection to your Plex server
	result, err := plexConnection.Test()
	if err != nil || !result {
		t.Error(errs.ErrPlexConnectionTest, err.Error())
		os.Exit(1)
	}

	devices, err := plexConnection.GetServers()
	if err != nil {
		t.Error(errs.ErrPlexInit, err.Error())
		os.Exit(1)
	}
	//spew.Dump(devices)

	mediaV := 0
	for i := 0; i < len(devices); i++ {
		if devices[i].ClientIdentifier == conf.Configuration.PlexClientID {
			mediaV = i
		}
	}

	mediaVaultProperties := devices[mediaV]
	//spew.Dump(mediaVaultProperties)

	mediaVault, err := plex.New(mediaVaultProperties.Connection[0].URI, conf.Configuration.PlexToken)
	if err != nil {
		t.Error(errs.ErrPlexConnect, mediaVaultProperties.Name)
		os.Exit(1)
	}

	mvLibraries, err := mediaVault.GetLibraries()
	if err != nil {
		t.Error(errs.ErrLibraryResponse, mediaVaultProperties.Name)
		os.Exit(1)
	}

	p := page.NewPage(t, lang.PlexTitle.Text()+symb.TextDelimiter.Symbol()+mediaVaultProperties.Name)
	count := 0
	for mvLibrary := range mvLibraries.MediaContainer.Directory {
		xx := mvLibraries.MediaContainer.Directory[mvLibrary]
		count++
		p.AddMenuOption(count, xx.Title, "", "")
	}

	p.AddAction(acts.Quit)
	p.AddAction(acts.Forward)
	p.AddAction(acts.Back)

	for {
		nextAction := p.Display_Actions()
		switch {
		case nextAction.Is(acts.Quit):
			return

		case nextAction.IsInt():
			naInt, _ := strconv.Atoi(nextAction.Action())
			wi := mvLibraries.MediaContainer.Directory[naInt-1]
			Action(t, mediaVault, &wi)

		default:
			p.Error(terr.ErrInvalidAction, t.Formatters.SQuote(nextAction.Action()))
		}
	}
}

func Action(t *term.ViewPort, mediaVault *plex.Plex, wi *plex.Directory) {

	switch wi.Type {
	case "movie":
		t.Shout(wi.Title)
		movies.Run(t, mediaVault, wi)
	case "show":
		t.Shout(wi.Title)
		shows.Run(t, mediaVault, wi)
	case "artist":
		t.Shout(wi.Title)
		music.Run(t, mediaVault, wi)
	default:
		t.Shout(wi.Title)
	}
}

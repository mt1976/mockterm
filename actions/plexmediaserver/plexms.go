package plexmediaserver

import (
	"os"
	"strconv"

	"github.com/jrudio/go-plex-client"
	support "github.com/mt1976/crt"
	movies "github.com/mt1976/mockterm/actions/plexmediaserver/movies"
	music "github.com/mt1976/mockterm/actions/plexmediaserver/music"
	shows "github.com/mt1976/mockterm/actions/plexmediaserver/shows"
	cfg "github.com/mt1976/mockterm/config"
	e "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	t "github.com/mt1976/mockterm/language"
)

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(crt *support.Crt) {

	crt.Clear()

	//spew.Dump(cfg.Configuration)
	//os.Exit(1)

	plexConnection, err := plex.New(cfg.Configuration.PlexURI+":"+cfg.Configuration.PlexPort, cfg.Configuration.PlexToken)
	if err != nil {
		crt.Error(e.ErrPlexInit, err.Error())
		os.Exit(1)
	}

	// Test your connection to your Plex server
	result, err := plexConnection.Test()
	if err != nil || !result {
		crt.Error(e.ErrPlexConnectionTest, err.Error())
		os.Exit(1)
	}

	devices, err := plexConnection.GetServers()
	if err != nil {
		crt.Error(e.ErrPlexInit, err.Error())
		os.Exit(1)
	}
	//spew.Dump(devices)

	mediaV := 0
	for i := 0; i < len(devices); i++ {
		if devices[i].ClientIdentifier == cfg.Configuration.PlexClientID {
			mediaV = i
		}
	}

	mediaVaultProperties := devices[mediaV]
	//spew.Dump(mediaVaultProperties)

	mediaVault, err := plex.New(mediaVaultProperties.Connection[0].URI, cfg.Configuration.PlexToken)
	if err != nil {
		crt.Error(e.ErrPlexConnect, mediaVaultProperties.Name)
		os.Exit(1)
	}

	mvLibraries, err := mediaVault.GetLibraries()
	if err != nil {
		crt.Error(e.ErrLibraryResponse, mediaVaultProperties.Name)
		os.Exit(1)
	}

	p := support.NewPageWithName(lang.TxtPlexTitle + lang.SymDelimiter + mediaVaultProperties.Name)
	count := 0
	for mvLibrary := range mvLibraries.MediaContainer.Directory {
		xx := mvLibraries.MediaContainer.Directory[mvLibrary]
		count++
		p.AddOption(count, xx.Title, "", "")
	}

	p.AddAction(t.SymActionQuit)
	p.AddAction(t.SymActionForward)
	p.AddAction(t.SymActionBack)

	nextAction, _ := p.Display(crt)
	switch {
	case nextAction == t.SymActionQuit:
		return
	case support.IsInt(nextAction):
		//crt.Error(lang.TxtYouSelected+nextAction, nil)
		naInt, _ := strconv.Atoi(nextAction)
		wi := mvLibraries.MediaContainer.Directory[naInt-1]
		Action(crt, mediaVault, &wi)

	default:
		crt.InputError(support.ErrInvalidAction, support.SQuote(nextAction))
	}
	//}

}

func Action(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	switch wi.Type {
	case "movie":
		crt.Shout(wi.Title)
		movies.Run(crt, mediaVault, wi)
	case "show":
		crt.Shout(wi.Title)
		shows.Run(crt, mediaVault, wi)
	case "artist":
		crt.Shout(wi.Title)
		music.Run(crt, mediaVault, wi)
	default:
		crt.Shout(wi.Title)
	}
}

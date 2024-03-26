package music

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"

	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
	errs "github.com/mt1976/mockterm/errors"
)

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

	nextAction, _ := p.DisplayWithActions()
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		if t.Helpers.IsInt(nextAction) {
			Detail(t, res.MediaContainer.Metadata[t.Helpers.ToInt(nextAction)-1])
		} else {
			p.Error(term.ErrInvalidAction, t.Formatters.SQuote(nextAction))
		}
	}
}

func Detail(t *term.ViewPort, info plex.Metadata) {

	p := t.NewPage(info.Title)

	p.AddFieldValuePair(lang.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(lang.TxtPlexSummaryLabel, info.Summary)

	count := 0
	p.AddBlankRow()
	p.AddColumnsTitle(lang.TxtPlexContainerLabel, lang.TxtPlexResolutionLabel, lang.TxtPlexCodecLabel, lang.TxtPlexAspectRatioLabel, lang.TxtPlexFrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	nextAction, _ := p.DisplayWithActions()
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		p.Error(term.ErrInvalidAction, t.Formatters.SQuote(nextAction))
	}
}

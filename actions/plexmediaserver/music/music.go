package music

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"

	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
	errs "github.com/mt1976/mockterm/errors"
)

func Run(crt *term.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := crt.NewTitledPage(res.MediaContainer.LibrarySectionTitle + lang.Space + crt.Formatters.PQuote(noItems))
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.AddOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	nextAction, _ := m.Display(crt)
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		if crt.Helpers.IsInt(nextAction) {
			//	Action(crt, mediaVault, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
			Detail(crt, res.MediaContainer.Metadata[crt.Helpers.ToInt(nextAction)-1])

		} else {
			m.PageError(crt, term.ErrInvalidAction, crt.Formatters.SQuote(nextAction))
		}
	}
}

func Detail(crt *term.Crt, info plex.Metadata) {

	p := crt.NewTitledPage(info.Title)

	p.AddFieldValuePair(crt, lang.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(crt, lang.TxtPlexSummaryLabel, info.Summary)

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, lang.TxtPlexContainerLabel, lang.TxtPlexResolutionLabel, lang.TxtPlexCodecLabel, lang.TxtPlexAspectRatioLabel, lang.TxtPlexFrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		p.PageError(crt, term.ErrInvalidAction, crt.Formatters.SQuote(nextAction))
	}
}

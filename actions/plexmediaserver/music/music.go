package music

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"

	support "github.com/mt1976/crt"
	t "github.com/mt1976/crt/language"
	text "github.com/mt1976/crt/language"
	e "github.com/mt1976/mockterm/errors"
)

func Run(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(e.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := support.NewPageWithName(res.MediaContainer.LibrarySectionTitle + t.Space + support.PQuote(noItems))
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.AddOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	nextAction, _ := m.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		if support.IsInt(nextAction) {
			//	Action(crt, mediaVault, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])

		} else {
			crt.InputError(support.ErrInvalidAction, support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata) {

	p := support.NewPageWithName(info.Title)

	p.AddFieldValuePair(crt, text.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(crt, text.TxtPlexSummaryLabel, info.Summary)

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, text.TxtPlexContainerLabel, text.TxtPlexResolutionLabel, text.TxtPlexCodecLabel, text.TxtPlexAspectRatioLabel, text.TxtPlexFrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		crt.InputError(support.ErrInvalidAction, support.SQuote(nextAction))
	}
}

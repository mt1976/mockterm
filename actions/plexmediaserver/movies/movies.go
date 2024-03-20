package movies

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	term "github.com/mt1976/crt"
	terr "github.com/mt1976/crt/errors"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	pmms "github.com/mt1976/mockterm/plexsupport"
)

func Run(crt *term.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := crt.NewTitledPage(res.MediaContainer.LibrarySectionTitle + lang.Space + term.PQuote(noItems))
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
		if term.IsInt(nextAction) {
			Detail(crt, res.MediaContainer.Metadata[term.ToInt(nextAction)-1])
		} else {
			crt.InputError(terr.ErrInvalidAction, term.SQuote(nextAction))
		}
	}
}

func Detail(crt *term.Crt, info plex.Metadata) {
	p := crt.NewTitledPage(info.Title)

	p.AddFieldValuePair(crt, lang.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(crt, lang.TxtPlexContentRatingLabel, info.ContentRating)
	dur := pmms.FormatPlexDuration(info.Duration)
	p.AddFieldValuePair(crt, lang.TxtPlexDurationLabel, dur)
	p.AddFieldValuePair(crt, lang.TxtPlexReleasedLabel, pmms.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, lang.TxtPlexSummaryLabel, info.Summary)
	//unix time to hrs mins secs
	p.BlankRow()
	for i := 0; i < len(info.Director); i++ {
		data := info.Director[i]
		lbl := lang.TxtPlexDirectorLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, data.Tag)
	}

	for i := 0; i < len(info.Writer); i++ {
		poobum := info.Writer[i]
		lbl := lang.TxtPlexWriterLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, poobum.Tag)
	}

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, lang.TxtPlexContainerLabel, lang.TxtPlexResolutionLabel, lang.TxtPlexCodecLabel, lang.TxtPlexAspectRatioLabel, lang.TxtPlexFrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	//range trhough parts
	p.BlankRow()
	p.AddColumnsTitle(crt, lang.TxtPlexMediaLabel)
	for _, v := range info.Media {
		p.AddColumns(crt, v.Part[0].File)
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		crt.InputError(terr.ErrInvalidAction, term.SQuote(nextAction))
	}

}

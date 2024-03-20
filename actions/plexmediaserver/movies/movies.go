package movies

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	support "github.com/mt1976/crt"
	x "github.com/mt1976/crt/errors"
	e "github.com/mt1976/mockterm/errors"

	t "github.com/mt1976/mockterm/language"
	plexSupport "github.com/mt1976/mockterm/plexsupport"
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
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
		} else {
			crt.InputError(x.ErrInvalidAction, support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata) {
	p := support.NewPageWithName(info.Title)

	p.AddFieldValuePair(crt, t.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(crt, t.TxtPlexContentRatingLabel, info.ContentRating)
	dur := plexSupport.FormatPlexDuration(info.Duration)
	p.AddFieldValuePair(crt, t.TxtPlexDurationLabel, dur)
	p.AddFieldValuePair(crt, t.TxtPlexReleasedLabel, plexSupport.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, t.TxtPlexSummaryLabel, info.Summary)
	//unix time to hrs mins secs
	p.BlankRow()
	for i := 0; i < len(info.Director); i++ {
		data := info.Director[i]
		lbl := t.TxtPlexDirectorLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, data.Tag)
	}

	for i := 0; i < len(info.Writer); i++ {
		poobum := info.Writer[i]
		lbl := t.TxtPlexWriterLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, poobum.Tag)
	}

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, t.TxtPlexContainerLabel, t.TxtPlexResolutionLabel, t.TxtPlexCodecLabel, t.TxtPlexAspectRatioLabel, t.TxtPlexFrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	//range trhough parts
	p.BlankRow()
	p.AddColumnsTitle(crt, t.TxtPlexMediaLabel)
	for _, v := range info.Media {
		p.AddColumns(crt, v.Part[0].File)
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		crt.InputError(x.ErrInvalidAction, support.SQuote(nextAction))
	}

}

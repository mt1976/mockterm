package shows

import (
	"os"

	"github.com/jrudio/go-plex-client"
	term "github.com/mt1976/crt"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	pmss "github.com/mt1976/mockterm/plexsupport"
)

func Episodes(crt *term.Crt, mediaVault *plex.Plex, seriesTitle string, info plex.Metadata) {
	res, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}
	m := crt.NewTitledPage(seriesTitle + lang.Space + info.Title)

	noEps := len(res.MediaContainer.Metadata)
	for i := 0; i < noEps; i++ {
		ep := res.MediaContainer.Metadata[i]
		m.AddOption(i+1, ep.Title, "", "")
	}

	nextAction, _ := m.DisplayWithActions(crt)
	switch nextAction {
	case lang.SymActionQuit:
		return
	default:
		if crt.Helpers.IsInt(nextAction) {
			EpisodeDetail(crt, res.MediaContainer.Metadata[crt.Helpers.ToInt(nextAction)-1])
		} else {
			m.Error(crt, term.ErrInvalidAction, crt.Formatters.SQuote(nextAction))
		}
	}
}

func EpisodeDetail(crt *term.Crt, info plex.Metadata) {

	title := info.GrandparentTitle + lang.Space + info.ParentTitle + lang.Space + info.Title
	p := crt.NewTitledPage(title)
	p.AddFieldValuePair(crt, lang.TxtPlexShow, info.GrandparentTitle)
	p.AddFieldValuePair(crt, lang.TxtPlexSeason, info.ParentTitle)
	p.AddFieldValuePair(crt, lang.TxtPlexEpisode, info.Title)
	p.AddFieldValuePair(crt, lang.TxtPlexSummaryLabel, info.Summary)
	p.AddFieldValuePair(crt, lang.TxtPlexDurationLabel, pmss.FormatPlexDuration(info.Duration))
	p.AddFieldValuePair(crt, lang.TxtPlexReleasedLabel, pmss.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, lang.TxtPlexContentRatingLabel, info.ContentRating)
	videoCodec := info.Media[0].VideoCodec
	videoFrameRate := info.Media[0].VideoFrameRate
	videoResolution := info.Media[0].VideoResolution
	videoContainer := info.Media[0].Container
	aspectRatio := info.Media[0].AspectRatio

	p.BlankRow()
	p.AddColumnsTitle(crt, lang.TxtPlexCodecLabel, lang.TxtPlexFrameRateLabel, lang.TxtPlexResolutionLabel, lang.TxtPlexContainerLabel, lang.TxtPlexAspectRatioLabel)
	p.AddColumns(crt, videoCodec, videoFrameRate, videoResolution, videoContainer, aspectRatio.String())
	p.BlankRow()
	p.AddColumnsTitle(crt, lang.TxtPlexMediaLabel)
	for _, v := range info.Media {
		p.AddColumns(crt, v.Part[0].File)
	}

	nextAction, _ := p.DisplayWithActions(crt)
	switch nextAction {
	case lang.SymActionQuit:
		return
	}
}

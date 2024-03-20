package shows

import (
	"os"

	"github.com/jrudio/go-plex-client"
	support "github.com/mt1976/crt"
	e "github.com/mt1976/mockterm/errors"
	notations "github.com/mt1976/mockterm/language"
	t "github.com/mt1976/mockterm/language"
	pms "github.com/mt1976/mockterm/plexsupport"
)

func Episodes(crt *support.Crt, mediaVault *plex.Plex, seriesTitle string, info plex.Metadata) {
	res, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error(e.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}
	m := support.NewPageWithName(seriesTitle + t.Space + info.Title)

	noEps := len(res.MediaContainer.Metadata)
	for i := 0; i < noEps; i++ {
		ep := res.MediaContainer.Metadata[i]
		m.AddOption(i+1, ep.Title, "", "")
	}

	nextAction, _ := m.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		if support.IsInt(nextAction) {
			EpisodeDetail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
		} else {
			crt.InputError(support.ErrInvalidAction, support.SQuote(nextAction))
		}
	}
}

func EpisodeDetail(crt *support.Crt, info plex.Metadata) {

	title := info.GrandparentTitle + t.Space + info.ParentTitle + t.Space + info.Title
	p := support.NewPageWithName(title)
	p.AddFieldValuePair(crt, notations.TxtPlexShow, info.GrandparentTitle)
	p.AddFieldValuePair(crt, notations.TxtPlexSeason, info.ParentTitle)
	p.AddFieldValuePair(crt, notations.TxtPlexEpisode, info.Title)
	p.AddFieldValuePair(crt, notations.TxtPlexSummaryLabel, info.Summary)
	p.AddFieldValuePair(crt, notations.TxtPlexDurationLabel, pms.FormatPlexDuration(info.Duration))
	p.AddFieldValuePair(crt, notations.TxtPlexReleasedLabel, pms.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, notations.TxtPlexContentRatingLabel, info.ContentRating)
	videoCodec := info.Media[0].VideoCodec
	videoFrameRate := info.Media[0].VideoFrameRate
	videoResolution := info.Media[0].VideoResolution
	videoContainer := info.Media[0].Container
	aspectRatio := info.Media[0].AspectRatio

	p.BlankRow()
	p.AddColumnsTitle(crt, notations.TxtPlexCodecLabel, notations.TxtPlexFrameRateLabel, notations.TxtPlexResolutionLabel, notations.TxtPlexContainerLabel, notations.TxtPlexAspectRatioLabel)
	p.AddColumns(crt, videoCodec, videoFrameRate, videoResolution, videoContainer, aspectRatio.String())
	p.BlankRow()
	p.AddColumnsTitle(crt, notations.TxtPlexMediaLabel)
	for _, v := range info.Media {
		p.AddColumns(crt, v.Part[0].File)
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	}
}

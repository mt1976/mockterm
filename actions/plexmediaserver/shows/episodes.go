package shows

import (
	"os"

	plexms "github.com/jrudio/go-plex-client"
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	plex "github.com/mt1976/mockterm/plexsupport"
)

func Episodes(t *term.ViewPort, mediaVault *plexms.Plex, seriesTitle string, info plexms.Metadata) {
	res, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		t.Error(errs.ErrLibraryResponse, err.Error())
		os.Exit(1)
	}
	p := page.NewPage(t, seriesTitle+lang.Space+info.Title)

	noEps := len(res.MediaContainer.Metadata)
	for i := 0; i < noEps; i++ {
		ep := res.MediaContainer.Metadata[i]
		p.AddMenuOption(i+1, ep.Title, "", "")
	}

	for {
		nextAction := p.Display_Actions()
		switch {
		case nextAction.Is(acts.Quit):
			return
		case nextAction.IsInt():
			EpisodeDetail(t, res.MediaContainer.Metadata[t.Helpers.ToInt(nextAction.Action())-1])
		default:
			p.Error(terr.ErrInvalidAction, t.Formatters.SQuote(nextAction.Action()))
		}
	}
}

func EpisodeDetail(t *term.ViewPort, info plexms.Metadata) {

	title := info.GrandparentTitle + lang.Space + info.ParentTitle + lang.Space + info.Title
	p := page.NewPage(t, title)
	p.AddFieldValuePair(lang.TxtPlexShow, info.GrandparentTitle)
	p.AddFieldValuePair(lang.TxtPlexSeason, info.ParentTitle)
	p.AddFieldValuePair(lang.TxtPlexEpisode, info.Title)
	p.AddFieldValuePair(lang.TxtPlexSummaryLabel, info.Summary)
	p.AddFieldValuePair(lang.TxtPlexDurationLabel, plex.FormatDuration(info.Duration))
	p.AddFieldValuePair(lang.TxtPlexReleasedLabel, plex.FormatDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(lang.TxtPlexContentRatingLabel, info.ContentRating)
	videoCodec := info.Media[0].VideoCodec
	videoFrameRate := info.Media[0].VideoFrameRate
	videoResolution := info.Media[0].VideoResolution
	videoContainer := info.Media[0].Container
	aspectRatio := info.Media[0].AspectRatio

	p.AddBlankRow()
	p.AddColumnsTitle(lang.TxtPlexCodecLabel, lang.TxtPlexFrameRateLabel, lang.TxtPlexResolutionLabel, lang.TxtPlexContainerLabel, lang.TxtPlexAspectRatioLabel)
	p.AddColumns(videoCodec, videoFrameRate, videoResolution, videoContainer, aspectRatio.String())
	p.AddBlankRow()
	p.AddColumnsTitle(lang.TxtPlexMediaLabel)
	for _, v := range info.Media {
		p.AddColumns(v.Part[0].File)
	}

	nextAction := p.Display_Actions()
	switch nextAction {
	case acts.Quit:
		return
	}
}

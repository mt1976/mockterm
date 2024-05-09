package music

import (
	plexms "github.com/jrudio/go-plex-client"
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/language"
)

func Detail(t *term.ViewPort, info plexms.Metadata) {

	p := page.NewPage(t, info.Title)

	p.AddFieldValuePair(lang.PlexTitle, info.Title)
	p.AddFieldValuePair(lang.PlexSummary, info.Summary)

	count := 0
	p.AddBlankRow()
	p.AddColumnsTitle(lang.PlexContainer.Text(), lang.PlexResolution.Text(), lang.PlexCodec.Text(), lang.PlexAspectRatio.Text(), lang.PlexFrameRate.Text())

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	for {
		nextAction := p.Display_Actions()
		switch {
		case nextAction.Is(acts.Quit):
			return
		default:
			p.Error(terr.ErrInvalidAction, t.Formatters.SQuote(nextAction.Action()))
		}
	}
}

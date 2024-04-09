package music

import (
	"github.com/jrudio/go-plex-client"

	term "github.com/mt1976/crt"
	lang "github.com/mt1976/crt/language"
)

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

	for {
		nextAction := p.Display_Actions()
		switch {
		case t.Formatters.Upcase(nextAction) == lang.SymActionQuit:
			return
		default:
			p.Error(term.ErrInvalidAction, t.Formatters.SQuote(nextAction))
		}
	}
}

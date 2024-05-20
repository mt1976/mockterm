package movies

import (
	plexms "github.com/jrudio/go-plex-client"
	terr "github.com/mt1976/crt/errors"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
	plng "github.com/mt1976/mockterm/actions/plexmediaserver/lang"
	lang "github.com/mt1976/mockterm/language"
	plex "github.com/mt1976/mockterm/plexsupport"
)

func Detail(t *term.ViewPort, info plexms.Metadata) {
	p := page.NewPage(t, lang.New(info.Title))

	p.AddFieldValuePair(plng.Title.Text(), info.Title)
	p.AddFieldValuePair(plng.ContentRating.Text(), info.ContentRating)
	dur := plex.FormatDuration(info.Duration)
	p.AddFieldValuePair(plng.Duration.Text(), dur)
	p.AddFieldValuePair(plng.Released.Text(), plex.FormatDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(plng.Summary, info.Summary)
	//unix time to hrs mins secs
	p.AddBlankRow()
	for i := 0; i < len(info.Director); i++ {
		data := info.Director[i]
		lbl := plng.Director.Text()
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(lbl, data.Tag)
	}

	for i := 0; i < len(info.Writer); i++ {
		poobum := info.Writer[i]
		lbl := plng.Writer.Text()
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(lbl, poobum.Tag)
	}

	count := 0
	p.AddBlankRow()
	p.AddColumnsTitle(plng.Container.Text(), plng.Resolution.Text(), plng.Codec.Text(), plng.AspectRatio.Text(), plng.FrameRate.Text())

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	//range trhough parts
	p.AddBlankRow()
	p.AddColumnsTitle(plng.Media.Text())
	for _, v := range info.Media {
		p.AddColumns(v.Part[0].File)
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

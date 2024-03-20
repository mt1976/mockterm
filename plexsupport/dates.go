package plexsupport

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	term "github.com/mt1976/crt"
	conf "github.com/mt1976/mockterm/config"
	lang "github.com/mt1976/mockterm/language"
)

var c = conf.Configuration

func PlexDateToDate(date string) time.Time {
	t, err := time.Parse(c.PlexDateFormat, date)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func PlexDurationToTime(duration int) time.Duration {
	//int to string
	d := fmt.Sprintf("%d", duration)
	//string to int
	i, err := strconv.Atoi(d)
	if err != nil {
		fmt.Println(err)
	}
	if i != duration {
		fmt.Println(ErrDurationMismatch, i, lang.Space, duration)
	}
	//return time.Duration(i) * time.Second
	t, err := time.ParseDuration(d + lang.TxtMillisecondsShort)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func FormatPlexDate(t string) string {
	return humanize.Time(PlexDateToDate(t)) + lang.Space + term.PQuote(term.FormatDate(PlexDateToDate(t)))
}

func FormatPlexDuration(t int) string {
	return term.FormatDuration(PlexDurationToTime(t))
}

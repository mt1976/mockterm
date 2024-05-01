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
var dummy = term.New()

func PlexDateToDate(date string) time.Time {
	t, err := time.Parse(c.PlexDateFormat, date)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func DurationToTime(duration int) time.Duration {
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

func FormatDate(t string) string {
	return humanize.Time(PlexDateToDate(t)) + lang.Space + dummy.Formatters.DQuote(dummy.Formatters.FormatDate(PlexDateToDate(t)))
}

func FormatDuration(t int) string {
	return dummy.Formatters.FormatDuration(DurationToTime(t))
}

package language

import (
	lang "github.com/mt1976/crt/language"
)

// General
var (
	YouSelected *lang.Text = lang.New("you selected: ")
	Done        *lang.Text = lang.New("DONE")
	LiveRun     *lang.Text = lang.New("This is a live run. PLEASE BE CAREFUL!")
	TrailRun    *lang.Text = lang.New("This is a trial run.")
	Mode        *lang.Text = lang.New("Mode")
	Path        *lang.Text = lang.New("Path")
	Title       *lang.Text = lang.New("Title")
)

// Page - Paging

const TxtPagingPrompt string = "Choose (F)orward, (B)ack or (Q)uit"

// Support
const (
	TxtOneWord      string = "one"
	TxtOneNumeric   string = "1"
	TxtMinutes      string = "minutes"
	TxtMinutesShort string = "mins"
	TxtHour         string = "hour"
	TxtHourShort    string = "hr"
)

var (
	TxtMillisecondsShort  string     = "ms"
	TxtApplicationVersion string     = "StarTerm - Utilities 1.0 %s"
	ApplicationName       *lang.Text = lang.New("StarTerm")
)

var ApplicationHeader []string = []string{
	"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
	"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
	"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
	"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
	"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
}

func New(in string) *lang.Text {
	return lang.New(in)
}

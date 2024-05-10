package language

import lang "github.com/mt1976/crt/language"

// PLEX - Media Server
var (
	PlexTitle         *lang.Text = lang.New("PMS") // TODO: Change to Title
	PlexSummary                  = lang.New("Summary")
	PlexContainer                = lang.New("Container")
	PlexResolution               = lang.New("Resolution")
	PlexCodec                    = lang.New("Codec")
	PlexAspectRatio              = lang.New("Aspect Ratio")
	PlexFrameRate                = lang.New("Frame Rate")
	PlexDuration                 = lang.New("Duration")
	PlexReleased                 = lang.New("Released")
	PlexDirector                 = lang.New("Director")
	PlexWriter                   = lang.New("Writer")
	PlexMedia                    = lang.New("Media")
	PlexContentRating            = lang.New("Content Rating")
	PlexShow                     = lang.New("Show")
	PlexSeason                   = lang.New("Season")
	PlexEpisode                  = lang.New("Episode")
	PlexSeasons                  = lang.New("Seasons")
	PlexEpisodes                 = lang.New("Episodes")
	PlexYear                     = lang.New("Year")
	PlexSeasonsPrompt            = lang.New("Choose (1...n)Season, (F)orward, (B)ack or (Q)uit")
)

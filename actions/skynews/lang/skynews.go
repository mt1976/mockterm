package language

import lang "github.com/mt1976/crt/language"

// SkyNews
var (
	Title         *lang.Text = lang.New("SKY News")
	Home          *lang.Text = lang.New("Home")
	UK            *lang.Text = lang.New("UK")
	World         *lang.Text = lang.New("World")
	US            *lang.Text = lang.New("US")
	Business      *lang.Text = lang.New("Business")
	Politics      *lang.Text = lang.New("Politics")
	Technology    *lang.Text = lang.New("Technology")
	Entertainment *lang.Text = lang.New("Entertainment")
	Strange       *lang.Text = lang.New("Strange News")
	LoadingTopic  *lang.Text = lang.New("Loading news for topic: %v")
	LoadingStory  *lang.Text = lang.New("Loading news for story...")
	HTMLTagTitle  string     = "title"
	HTMLTagP      string     = "p"
	Topic         *lang.Text = lang.New("Topic")
	Story         *lang.Text = lang.New("Story")
)

package pushover

import (
	"errors"
	"runtime"

	"github.com/gregdel/pushover"
	page "github.com/mt1976/crt/page"
	acts "github.com/mt1976/crt/page/actions"
	symb "github.com/mt1976/crt/strings/symbols"
	term "github.com/mt1976/crt/terminal"
	lang "github.com/mt1976/mockterm/language"
)

// Init Objects

// var (
//
//	msgType  string
//	msgTitle string
//	msgBody  string
//
// )
var messages = make(map[string]string)
var titles = make(map[string]string)
var hostName string
var t term.ViewPort
var sendAction = acts.New("S")

//var debugMode bool

type MessageTypes struct {
	Up        string
	Heartbeat string
	Other     string
	Message   string
}

// End Init Objects
func init() {

	var types = MessageTypes{}
	types.Up = "-up"
	types.Heartbeat = "-heartbeat"
	types.Other = "-other"
	types.Message = "-message"
	hostName = "unknown"
	if !(runtime.GOOS == "windows") {
		hostName = term.New().Helpers.GetHostName()
	}
	messages[types.Up] = hostName + " has started - System is available"
	messages[types.Heartbeat] = hostName + " - Heartbeat - OK"
	messages[types.Other] = hostName + " - Unknown Message"
	messages[types.Message] = hostName + " - Message"

	titles[types.Up] = "System"
	titles[types.Heartbeat] = "Heartbeat"
	titles[types.Other] = "Other"
	titles[types.Message] = "This Message"
}

func Run(t *term.ViewPort) {
	optionsScreen := page.NewPage(t, lang.TxtPushoverTitle)
	optionsScreen.AddParagraph(lang.TxtPushoverDescription)
	optionsScreen.AddBlankRow()
	optionsScreen.AddMenuOption(1, lang.TxtPushoverMsgPriorityNormal, "", "")
	optionsScreen.AddMenuOption(2, lang.TxtPushoverMsgPriorityHigh, "", "")
	optionsScreen.AddMenuOption(3, lang.TxtPushoverMsgPriorityLow, "", "")
	optionsScreen.AddMenuOption(4, lang.TxtPushoverMsgPriorityEmergency, "", "")
	optionsScreen.SetPrompt(lang.TxtPushoverPrompt)
	optionsScreen.ShowOptions()
	optionsScreen.AddAction(acts.Quit)
	action := optionsScreen.Display_Actions()
	if action.Is(acts.Quit) {
		return
	}
	if action.IsInt() {

		err := processMessage(t, action)
		if err != nil {
			t.Error(err, "")
			return
		}

	}
}

func processMessage(t *term.ViewPort, action *acts.Action) error {

	var priority int
	switch action.Int() {
	case 1:
		priority = pushover.PriorityNormal
	case 2:
		priority = pushover.PriorityHigh
	case 3:
		priority = pushover.PriorityLow
	case 4:
		priority = pushover.PriorityEmergency
	default:
		priority = pushover.PriorityNormal
	}

	//messageBody := "Message Body"
	messageTitle, na, err := getMessageTitle(t)
	if err != nil {
		t.InputError(err, na)
		return err
	}
	if acts.Quit.Equals(na) {
		return nil
	}

	messageBody, fa, err := getMessageBody(t, messageTitle)
	if err != nil {
		t.InputError(err)
		return err
	}
	if t.Helpers.IsActionIn(fa, acts.Quit) {
		return nil
	}

	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, priority)

	p := page.NewPage(t, lang.TxtPushoverTitle)
	p.AddBlankRow()
	p.AddFieldValuePair("Title", message.Title)
	p.AddFieldValuePair("Message", message.Message)
	p.AddBlankRow()
	p.AddFieldValuePair("Priority", t.Formatters.Upcase(poPriorityToString(message.Priority)))
	p.AddFieldValuePair("Device Name", message.DeviceName)
	ts := t.Formatters.HumanFromUnixDate(message.Timestamp)
	p.AddFieldValuePair("Timestamp", ts)
	p.AddFieldValuePair("Expires at", message.Expire.String())
	p.AddFieldValuePair("Retry every", message.Retry.String())
	p.AddFieldValuePair("URLTitle", message.URLTitle)
	p.AddFieldValuePair("URL", message.URL)
	p.AddFieldValuePair("CallbackURL", message.CallbackURL)
	p.AddFieldValuePair("Sound", message.Sound)
	p.AddAction(sendAction)
	p.AddAction(acts.Quit)
	p.SetPrompt(lang.TxtPushoverConfirmation)
	p.ShowOptions()

	for {
		act := p.Display_Actions()
		if act.Is(sendAction) {
			p.Info(lang.TxtPushoverMessageSending)
			_, err = app.SendMessage(message, recipient)
			if err != nil {
				t.Error(err, "")
				return err
			}
			p.Info(lang.TxtPushoverMessageSent)
		}
		if act.Is(acts.Quit) {
			return nil
		}
	}
	return nil
}

func upcase(p *page.Page, in string) string {
	return p.ViewPort().Formatters.Upcase(in)
}

func getMessageTitle(t *term.ViewPort) (string, string, error) {
	p := page.NewPage(t, lang.TxtPushoverTitle)
	p.AddBlankRow()
	p.AddFieldValuePair("Title", "")
	p.AddBlankRow()
	xx := []string{symb.Blank.Symbol(), "Enter a title of the message to be sent to pushover", symb.Blank.Symbol(), "e.g. :-", "Test Pushover Message", "System Pushover Message", "Heartbeat Message"}
	p.AddParagraph(xx)
	p.AddBlankRow()
	p.SetPrompt(lang.TxtPushoverMessageTitlePrompt)
	for {
		op, _ := p.Display_Input(3, 15)
		//op := t.Input("", "")
		if acts.Quit.Equals(op) {
			return "", acts.Quit.Action(), nil
		}
		if op != "" {
			return op, "", nil
		}
	}

	return "", "", errors.New("Failed to get message title")
}

func getMessageBody(t *term.ViewPort, title string) (string, string, error) {

	p := page.NewPage(t, lang.TxtPushoverTitle)
	p.AddBlankRow()
	p.AddFieldValuePair("Title", title)
	p.AddFieldValuePair("Message", "")
	p.AddBlankRow()
	xx := []string{symb.Blank.Symbol(), "Enter a message to be sent to pushover", symb.Blank.Symbol(), "e.g. :-", "This is a test message", "This is a system message", "This is a heartbeat message"}
	p.AddParagraph(xx)
	p.AddBlankRow()
	p.SetPrompt(lang.TxtPushoverMessageBodyPrompt)

	//p.SetPrompt("Enter the title of the message, or (Q)uit")
	for {
		op, _ := p.Display_Input(3, 20)
		//op := t.Input("", "")
		if acts.Quit.Equals(op) {
			return "", acts.Quit.Action(), nil
		}
		if op != "" {
			return op, "", nil
		}
	}

	return "", "", errors.New("Failed to get message title")
}
func poPriorityToString(in int) string {
	switch in {
	case pushover.PriorityNormal:
		return "Normal"
	case pushover.PriorityHigh:
		return "High"
	case pushover.PriorityLow:
		return "Low"
	case pushover.PriorityEmergency:
		return "Emergency"
	default:
		return "Unknown"
	}
}

package pushover

import (
	"errors"
	flags "flag"
	"time"

	"github.com/gregdel/pushover"
	term "github.com/mt1976/crt"
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

	hostName = term.New().Helpers.GetHostName()

	messages[types.Up] = hostName + " has started - System is available"
	messages[types.Heartbeat] = hostName + " - Heartbeat - OK"
	messages[types.Other] = hostName + " - Unknown Message"
	messages[types.Message] = hostName + " - Message"

	titles[types.Up] = "System"
	titles[types.Heartbeat] = "Heartbeat"
	titles[types.Other] = "Other"
	titles[types.Message] = "This Message"
}

func RunOld(t term.ViewPort, debug bool, messageType, messageTitle, messageBody string) {

	//debugMode = debug

	flags.Parse()

	t.Print("Starting Alive")
	t.Print("Message Type: " + messageType)
	t.Print("Message Title: " + messageTitle)
	t.Print("Message Body: " + messageBody)
	t.Blank()

	//argsWithoutProg := os.Args[1:]
	if messageType != "" {
		//fmt.Println("Args: ", argsWithoutProg)
		//CONFIG.DebugCFG()
		//Get Time
		now := time.Now().Format("2006-01-02 15:04:05")

		//L.WithFields(xlg.Fields{"args": argsWithoutProg, "msgType": msgType}).Info("Arguments")
		switch messageType {
		case "up":
			//xlg.Info("ACTION=UP")
			sendMessage("System Started", hostName+" started at "+now)
		case "heartbeat":
			//xlg.Info("ACTION=HEARTBEAT")
			sendMessage("System Online", hostName+" online at "+now)
		case "other":
			//xlg.Info("ACTION=OTHER")
			sendMessage("Unknown Message", "An unknown message was received at "+now)
		default:
			//xlg.Info("ACTION=RAW")
			sendMessage(messageTitle, messageBody)
		}
	}
}

//TODO create a menu to select the type of notification to send
//TODO create a menu to offer a series of default notifications, or a custom notification
//TODO add input box to enter the message to send
//TODO add a Preview page, with the message type, title, body etc.
//TODO add a confirmation box to confirm the message to send

func Run(t *term.ViewPort) {
	optionsScreen := t.NewPage(lang.TxtPushoverTitle)
	optionsScreen.AddParagraph(lang.TxtPushoverDescription)
	optionsScreen.AddBlankRow()
	optionsScreen.AddMenuOption(1, lang.TxtPushoverMsgPriorityNormal, "", "")
	optionsScreen.AddMenuOption(2, lang.TxtPushoverMsgPriorityHigh, "", "")
	optionsScreen.AddMenuOption(3, lang.TxtPushoverMsgPriorityLow, "", "")
	optionsScreen.AddMenuOption(4, lang.TxtPushoverMsgPriorityEmergancy, "", "")
	optionsScreen.SetPrompt(lang.TxtPushoverPrompt)
	optionsScreen.AddAction(lang.SymActionQuit)
	action, _ := optionsScreen.DisplayWithActions()
	if action == lang.SymActionQuit {
		return
	}
	if t.Helpers.IsInt(action) {

		po, recp, msg, err := processMessage(t, action)
		if err != nil {
			t.Error(err, "")
			return
		}

		_, err = po.SendMessage(msg, recp)
		if err != nil {
			t.Error(err, "")
			return
		}
	}
}

func sendMessage(inMessage, inTitle string) {
	//L.WithFields(xlg.Fields{"Message": inMessage, "Title": inTitle}).Info("Sending Message")
	t.Print("Sending Message")
	t.Print("Message: " + inMessage)
	t.Print("Title: " + inTitle)
	Normal(inMessage, inTitle)
	t.Print("Message Sent")
}

func processMessage(t *term.ViewPort, action string) (*pushover.Pushover, *pushover.Recipient, *pushover.Message, error) {

	var priority int
	switch action {
	case "1":
		priority = pushover.PriorityNormal
	case "2":
		priority = pushover.PriorityHigh
	case "3":
		priority = pushover.PriorityLow
	case "4":
		priority = pushover.PriorityEmergency
	default:
		priority = pushover.PriorityNormal
	}

	//messageBody := "Message Body"
	messageTitle, action, err := getMessageTitle(t)
	if err != nil {
		t.InputError(err)
		return nil, nil, nil, err
	}
	if t.Helpers.IsActionIn(action, lang.SymActionQuit) {
		return nil, nil, nil, nil
	}

	messageBody, action, err := getMessageBody(t, messageTitle)
	if err != nil {
		t.InputError(err)
		return nil, nil, nil, err
	}
	if t.Helpers.IsActionIn(action, lang.SymActionQuit) {
		return nil, nil, nil, nil
	}

	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, priority)

	p := t.NewPage(lang.TxtPushoverTitle)
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
	p.AddAction("S")
	p.AddAction("Q")
	p.SetPrompt(lang.TxtPushoverConfirmation)
	//spew.Dump(p, t)
	p.DisplayWithActions()
	return app, recipient, message, nil
}

func getMessageTitle(t *term.ViewPort) (string, string, error) {
	p := t.NewPage(lang.TxtPushoverTitle)
	p.AddBlankRow()
	p.AddFieldValuePair("Title", "")
	p.AddBlankRow()
	xx := []string{lang.SymBlank, "Enter a title of the message to be sent to pushover", lang.SymBlank, "e.g. :-", "Test Pushover Message", "System Pushover Message", "Heartbeat Message"}
	p.AddParagraph(xx)
	p.AddBlankRow()
	p.SetPrompt(lang.TxtPushoverMessageTitlePrompt)
	for {
		op, _ := p.DisplayAndInput(3, 15)
		//op := t.Input("", "")
		if op == lang.SymActionQuit {
			return "", lang.SymActionQuit, nil
		}
		if op != "" {
			return op, "", nil
		}
	}

	return "", "", errors.New("Failed to get message title")
}

func getMessageBody(t *term.ViewPort, title string) (string, string, error) {

	p := t.NewPage(lang.TxtPushoverTitle)
	p.AddBlankRow()
	p.AddFieldValuePair("Title", title)
	p.AddFieldValuePair("Message", "")
	p.AddBlankRow()
	xx := []string{lang.SymBlank, "Enter a message to be sent to pushover", lang.SymBlank, "e.g. :-", "This is a test message", "This is a system message", "This is a heartbeat message"}
	p.AddParagraph(xx)
	p.AddBlankRow()
	p.SetPrompt(lang.TxtPushoverMessageBodyPrompt)

	//p.SetPrompt("Enter the title of the message, or (Q)uit")
	for {
		op, _ := p.DisplayAndInput(3, 20)
		//op := t.Input("", "")
		if op == lang.SymActionQuit {
			return "", lang.SymActionQuit, nil
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

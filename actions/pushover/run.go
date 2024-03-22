package pushover

import (
	flags "flag"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
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
var crt term.Crt

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

func RunOld(t term.Crt, debug bool, messageType, messageTitle, messageBody string) {

	//crt = t
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

func Run(t *term.Crt) {
	optionsScreen := t.NewTitledPage(lang.TxtPushoverTitle)
	optionsScreen.Paragraph(lang.TxtPushoverDescription)
	optionsScreen.BlankRow()
	optionsScreen.AddOption(1, lang.TxtPushoverMsgPriorityNormal, "", "")
	optionsScreen.AddOption(2, lang.TxtPushoverMsgPriorityHigh, "", "")
	optionsScreen.AddOption(3, lang.TxtPushoverMsgPriorityLow, "", "")
	optionsScreen.AddOption(4, lang.TxtPushoverMsgPriorityEmergancy, "", "")
	optionsScreen.SetPrompt(lang.TxtPushoverPrompt)
	optionsScreen.AddAction(lang.SymActionQuit)
	action, _ := optionsScreen.Display(t)
	if action == lang.SymActionQuit {
		return
	}
	if t.Helpers.IsInt(action) {

		po, recp, msg, err := processMessage(t, action)
		if err != nil {
			t.Error(err, "")
			return
		}
		return
		_, err = po.SendMessage(msg, recp)
		if err != nil {
			t.Error(err, "")
			return
		}
	}
}

func sendMessage(inMessage, inTitle string) {
	//L.WithFields(xlg.Fields{"Message": inMessage, "Title": inTitle}).Info("Sending Message")
	crt.Print("Sending Message")
	crt.Print("Message: " + inMessage)
	crt.Print("Title: " + inTitle)
	Normal(inMessage, inTitle)
	crt.Print("Message Sent")
}

func processMessage(t *term.Crt, action string) (*pushover.Pushover, *pushover.Recipient, *pushover.Message, error) {

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

	messageBody := "Message Body"
	messageTitle := fmt.Sprintf("%v Priority Message Title", poPriorityToString(priority))

	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, priority)

	p := t.NewTitledPage("Pushover Message Preview")
	p.AddFieldValuePair(t, "Title", message.Title)
	p.AddFieldValuePair(t, "Message", message.Message)
	p.BlankRow()
	p.AddFieldValuePair(t, "Priority", t.Formatters.Upcase(poPriorityToString(message.Priority)))
	p.AddFieldValuePair(t, "Device Name", message.DeviceName)
	ts := t.Formatters.HumanFromUnixDate(message.Timestamp)
	p.AddFieldValuePair(t, "Timestamp", ts)
	p.AddFieldValuePair(t, "Expires at", message.Expire.String())
	p.AddFieldValuePair(t, "Retry every", message.Retry.String())
	p.AddFieldValuePair(t, "URLTitle", message.URLTitle)
	p.AddFieldValuePair(t, "URL", message.URL)
	p.AddFieldValuePair(t, "CallbackURL", message.CallbackURL)
	p.AddFieldValuePair(t, "Sound", message.Sound)
	p.AddAction("S")
	p.AddAction("Q")
	p.SetPrompt(lang.TxtPushoverConfirmation)
	spew.Dump(p, t)
	p.Display(t)
	return app, recipient, message, nil
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

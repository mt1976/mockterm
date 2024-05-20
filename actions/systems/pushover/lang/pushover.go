package language

import (
	lang "github.com/mt1976/crt/language"
)

var (
	TxtPushoverTitle                *lang.Text = lang.New("Pushover Messaging Service")
	TxtServiceMenuDescription       []string   = []string{"This menu shows the list of services available for maintenance.", "Select the service you wish to use. PLEASE BE CAREFUL!"}
	TxtPushoverDescription          []string   = []string{"Pushover is a service to receive instant push notifications on your phone or tablet from a variety of sources."}
	TxtPushoverMsgPriorityEmergency string     = "Emergency Message"
	TxtPushoverMsgPriorityNormal    string     = "Normal Priority"
	TxtPushoverMsgPriorityHigh      string     = "High Priority"
	TxtPushoverMsgPriorityLow       string     = "Low Priority"
	TxtPushoverPrompt               *lang.Text = lang.New("Choose a message type to send")
	TxtPushoverConfirmation         *lang.Text = lang.New("Choose (S)end or (Q)uit")
	TxtPushoverMessageTitlePrompt   *lang.Text = lang.New("Enter the title of the message, or (Q)uit")
	TxtPushoverMessageBodyPrompt    *lang.Text = lang.New("Enter the body of the message, or (Q)uit")
	TxtPushoverMessageSending       *lang.Text = lang.New("Sending Pushover Message")
	TxtPushoverMessageSent          *lang.Text = lang.New("Pushover Message Sent")
)

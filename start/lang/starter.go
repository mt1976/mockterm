package language

import lang "github.com/mt1976/crt/language"

// Starter - Terminal Startup Simulator
var (
	Initialise       *lang.Text = lang.New("Initialise...")
	StartingTerminal *lang.Text = lang.New("Starting Terminal...")
	SelfTesting      *lang.Text = lang.New("Self Testing...")
	CurrentDate      *lang.Text = lang.New("Current Date: ")
	CurrentTime      *lang.Text = lang.New("Current Time: ")
	PleaseWait       *lang.Text = lang.New("Please Wait...")
	BaudRate         *lang.Text = lang.New("Baud Rate Set to %v kbps")
	Connecting       *lang.Text = lang.New("Connecting...")
	Dialing          *lang.Text = lang.New("Dialing... %v:%v")
	Connected        *lang.Text = lang.New("Connected.")
	ConnectionFailed *lang.Text = lang.New("Connection failed. Retrying...")
	Complete         *lang.Text = lang.New("Complete")
)

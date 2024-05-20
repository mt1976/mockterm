package language

import lang "github.com/mt1976/crt/language"

// Title - Server Title
var (
	Title                       *lang.Text = lang.New("Server Dashboard")
	TxtStatusOnline                        = lang.New("Online ")
	TxtStatusOffline                       = lang.New("Offline")
	TxtDashboardChecking                   = lang.New("Please wait... Checking Services")
	TxtDashboardCheckingService            = lang.New("Checking %v...")
	TxtNoResponseFromServer                = lang.New("No response from server")
)

package language

import lang "github.com/mt1976/crt/language"

// Dashboard - Server Dashboard
var (
	Dashboard                   *lang.Text = lang.New("Server Dashboard")
	TxtStatusOnline             string     = "Online "
	TxtStatusOffline            string     = "Offline"
	TxtDashboardChecking        string     = "Please wait... Checking Services"
	TxtDashboardCheckingService string     = "Checking %v..."
	TxtNoResponseFromServer     string     = "No response from server"
)

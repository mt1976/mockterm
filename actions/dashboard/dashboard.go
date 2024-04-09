package dashboard

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strconv"

	term "github.com/mt1976/crt"
	terr "github.com/mt1976/crt/errors"
	conf "github.com/mt1976/mockterm/config"
	errs "github.com/mt1976/mockterm/errors"
	lang "github.com/mt1976/mockterm/language"
	probing "github.com/prometheus-community/pro-bing"
)

var C = conf.Configuration
var dummy = term.New()

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(t *term.ViewPort) {

	p := t.NewPage(lang.TxtDashboardTitle)
	p.Info(lang.TxtDashboardChecking)
	c := 0
	c++
	//p.Add("Testing Server/Service Dashboard", "", time.Now().Format("2006-01-02"))
	for i := 0; i < C.DashboardURINoEntries; i++ {
		//p.Add(C.DashboardURIName[i], "", "")
		p.Info(fmt.Sprintf(lang.TxtDashboardCheckingService, C.DashboardURIName[i]))
		result := CheckService(i)
		//p.AddFieldValuePair(C.DashboardURIName[i], dummy.Formatters.Bold(result))
		p.AddFieldValuePair(C.DashboardURIName[i], result)
	}

	p.AddAction(lang.SymActionQuit)
	p.AddAction(lang.SymActionForward)
	p.AddAction(lang.SymActionBack)
	ok := false
	for !ok {

		nextAction := p.Display_Actions()
		switch nextAction {
		case lang.SymActionForward:
			p.Forward()
		case lang.SymActionBack:
			p.Back()
		case lang.SymActionQuit:
			ok = true
			return
		default:
			p.Error(terr.ErrInvalidAction, nextAction)
		}
	}
}

// CheckService checks the status of a service
func CheckService(i int) string {
	// Extract the configuration values for the service
	protocol := C.DashboardURIProtocol[i]
	host := C.DashboardURIHost[i]
	if host == "" {
		host = C.DashboardDefaultHost
	}
	if host == "" {
		panic(errs.ErrDashboardNoHost)
	}
	port := C.DashboardURIPort[i]
	if port == "" {
		port = C.DashboardDefaultPort
	}
	query := C.DashboardURIQuery[i]
	operation := C.DashboardURIOperation[i]
	success := C.DashboardURISuccess[i]

	// Check if the operation is a valid operation
	if !slices.Contains(C.DashboardURIValidActions, dummy.Formatters.Upcase(operation)) {
		return term.ErrInvalidAction.Error()
	}

	// Ping the service
	if dummy.Formatters.Upcase(operation) == "PING" {
		pinger, err := probing.NewPinger(host)
		if err != nil {
			return lang.TxtStatusOffline + lang.Space + dummy.Formatters.PQuote(err.Error())
		}
		pinger.Count = 3
		err = pinger.Run() // Blocks until finished.
		if err != nil {
			return lang.TxtStatusOffline + lang.Space + dummy.Formatters.PQuote(err.Error())
		}
		stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
		avgRtt := stats.AvgRtt

		return lang.TxtStatusOnline + lang.Space + dummy.Formatters.PQuote(fmt.Sprintf("%v", avgRtt))
	}

	// Perform an HTTP request to the service
	if dummy.Formatters.Upcase(operation) == "HTTP" {
		//fmt.Printf("GET %v://%v:%v%v - %v %v\n", protocol, host, port, query, operation, success)
		var u url.URL

		u.Scheme = protocol
		u.Host = host + ":" + port
		u.Path = query

		//fmt.Println(u)

		return StatusCode(u.String(), "", success)
	}

	// Return the operation and success values
	return fmt.Sprintf("%v://%v:%v%v - %v %v", protocol, host, port, query, operation, success)
}

// StatusCode performs an HTTP request to the specified page and returns the status code and status message.
// The status code is compared to the specified success code, and if they match, the status message for
// online is returned. If the status code does not match the success code, the status message for offline
// is returned.
//
// The function takes the following parameters:
//
//	PAGE: The URL of the page to request.
//	AUTH: The authorization header value to use for the request.
//	SUCCESS: The expected status code for a successful response.
//
// The function returns the status message for the specified status code.
func StatusCode(PAGE string, AUTH string, SUCCESS string) (r string) {
	// Setup the request.
	req, err := http.NewRequest("GET", PAGE, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", AUTH)

	// Execute the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return lang.TxtStatusOffline + lang.Space + dummy.Formatters.PQuote(lang.TxtNoResponseFromServer)
	}

	// Close response body as required.
	defer resp.Body.Close()

	//fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode == 200 {
		return lang.TxtStatusOnline + lang.Space + dummy.Formatters.PQuote(resp.Status)
	}
	//resp.StatusCode to string
	scString := strconv.Itoa(resp.StatusCode)
	if scString == SUCCESS {
		return lang.TxtStatusOnline + lang.Space + dummy.Formatters.PQuote(resp.Status)
	}

	return lang.TxtStatusOffline + lang.Space + dummy.Formatters.PQuote(resp.Status)
	// or fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}

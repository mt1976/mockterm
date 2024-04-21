package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gen2brain/beeep"
	errs "github.com/mt1976/crt/errors"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

type Config struct {
	PlexURI                    string  `mapstructure:"PlexURI"`
	PlexPort                   string  `mapstructure:"PlexPort"`
	PlexToken                  string  `mapstructure:"PlexToken"`
	PlexClientID               string  `mapstructure:"PlexClientID"`
	PlexDateFormat             string  `mapstructure:"PlexDateFormat"`
	ApplicationDateFormat      string  `mapstructure:"ApplicationDateFormat"`
	ApplicationDateFormatShort string  `mapstructure:"ApplicationDateFormatShort"`
	ApplicationTimeFormat      string  `mapstructure:"ApplicationTimeFormat"`
	TerminalWidth              int     `mapstructure:"TerminalWidth"`
	TerminalHeight             int     `mapstructure:"TerminalHeight"`
	Delay                      float64 `mapstructure:"Delay"`
	Baud                       int     `mapstructure:"Baud"`
	TransmissionURI            string  `mapstructure:"TransmissionURI"`
	QTorrentURI                string  `mapstructure:"QTorrentURI"`
	MaxContentRows             int     `mapstructure:"MaxContentRows"`
	MaxNoItems                 int     `mapstructure:"MaxNoItems"`
	TitleLength                int     `mapstructure:"TitleLength"`
	Debug                      bool    `mapstructure:"Debug"`

	OpenWeatherMapApiKey   string `mapstructure:"OpenWeatherMapApiKey"`
	OpenWeatherMapApiLang  string `mapstructure:"OpenWeatherMapApiLang"`
	OpenWeatherMapApiUnits string `mapstructure:"OpenWeatherMapApiUnits"`

	LocationLogitude float64 `mapstructure:"LocationLongitude"`
	LocationLatitude float64 `mapstructure:"LocationLatitude"`

	URISkyNews              string `mapstructure:"SkyNewsURI"`
	URISkyNewsHome          string `mapstructure:"SkyNewsHomeURI"`
	URISkyNewsUK            string `mapstructure:"SkyNewsUKURI"`
	URISkyNewsWorld         string `mapstructure:"SkyNewsWorldURI"`
	URISkyNewsUS            string `mapstructure:"SkyNewsUSURI"`
	URISkyNewsBusiness      string `mapstructure:"SkyNewsBusinessURI"`
	URISkyNewsPolitics      string `mapstructure:"SkyNewsPoliticsURI"`
	URISkyNewsTechnology    string `mapstructure:"SkyNewsTechnologyURI"`
	URISkyNewsEntertainment string `mapstructure:"SkyNewsEntertainmentURI"`
	URISkyNewsStrange       string `mapstructure:"SkyNewsStrangeURI"`

	DefaultErrorDelay    float64 `mapstructure:"DefaultErrorDelay"`
	DefaultRandomPortMin int     `mapstructure:"DefaultRandomPortMin"`
	DefaultRandomPortMax int     `mapstructure:"DefaultRandomPortMax"`
	DefaultRandomMACMin  int     `mapstructure:"DefaultRandomMACMin"`
	DefaultRandomMACMax  int     `mapstructure:"DefaultRandomMACMax"`
	DefaultRandomIPMin   int     `mapstructure:"DefaultRandomIPMin"`
	DefaultRandomIPMax   int     `mapstructure:"DefaultRandomIPMax"`
	DefaultBaud          int     `mapstructure:"DefaultBaud"`

	DefaultBeepDuration  int
	DefaultBeepFrequency float64

	ValidBaudRates          []int
	ValidFileNameCharacters []string

	DashboardURINameIN      string `mapstructure:"DashboardURIName"`
	DashboardURIProtocolIN  string `mapstructure:"DashboardURIProtocol"`
	DashboardURIHostIN      string `mapstructure:"DashboardURIHost"`
	DashboardURIPortIN      string `mapstructure:"DashboardURIPort"`
	DashboardURIQueryIN     string `mapstructure:"DashboardURIQuery"`
	DashboardURIOperationIN string `mapstructure:"DashboardURIOperation"`
	DashboardURISuccessIN   string `mapstructure:"DashboardURISuccess"`
	DashboardOrderIN        string `mapstructure:"DashboardOrder"`
	DashboardDefaultHost    string `mapstructure:"DashboardDefaultHost"`
	DashboardDefaultPort    string

	DashboardURIName         []string
	DashboardURIProtocol     []string
	DashboardURIHost         []string
	DashboardURIPort         []string
	DashboardURIQuery        []string
	DashboardURIOperation    []string
	DashboardURISuccess      []string
	DashboardURIValidActions []string
	DashboardURINoEntries    int
	DashboardOrdering        []int
	TimeStampFormat          string
	OnlyFansDateTimeFormat   string
}

var Configuration = Config{}

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.SetConfigName("mockterm")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Configuration)
	if err != nil {
		panic(err)
	}

	Configuration.DefaultBeepDuration = beeep.DefaultDuration
	Configuration.DefaultBeepFrequency = beeep.DefaultFreq
	Configuration.ValidBaudRates = []int{0, 300, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200}
	Configuration.ValidFileNameCharacters = []string{" ", "-", "_", ".", "(", ")", "[", "]", "!", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b",
		"c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	//Configuration.PlexPort = strconv.Itoa(Configuration.plexPortInt)
	// spew.Dump(&Configuration)
	// os.Exit(1)
	Configuration.DashboardURIName = split(Configuration.DashboardURINameIN)
	Configuration.DashboardURIHost = split(Configuration.DashboardURIHostIN)
	Configuration.DashboardURIProtocol = split(Configuration.DashboardURIProtocolIN)
	Configuration.DashboardURIPort = split(Configuration.DashboardURIPortIN)
	Configuration.DashboardURIQuery = split(Configuration.DashboardURIQueryIN)
	Configuration.DashboardURIOperation = split(Configuration.DashboardURIOperationIN)
	Configuration.DashboardURISuccess = split(Configuration.DashboardURISuccessIN)
	NoEntries := len(Configuration.DashboardURIHost)
	Configuration.DashboardURINoEntries = NoEntries

	if NoEntries != len(Configuration.DashboardURIProtocol) {
		panic(fmt.Sprintf(errs.ErrConfigurationColumnMismatch.Error(), len(Configuration.DashboardURIProtocol), NoEntries, "DashboardURIProtocol"))
	}
	if NoEntries != len(Configuration.DashboardURIPort) {
		panic(fmt.Sprintf(errs.ErrConfigurationColumnMismatch.Error(), len(Configuration.DashboardURIPort), NoEntries, "DashboardURIPort"))
	}
	if NoEntries != len(Configuration.DashboardURIQuery) {
		panic(fmt.Sprintf(errs.ErrConfigurationColumnMismatch.Error(), len(Configuration.DashboardURIQuery), NoEntries, "DashboardURIQuery"))
	}
	if NoEntries != len(Configuration.DashboardURIName) {
		panic(fmt.Sprintf(errs.ErrConfigurationColumnMismatch.Error(), len(Configuration.DashboardURIName), NoEntries, "DashboardURIName"))
	}
	if NoEntries != len(Configuration.DashboardURIOperation) {
		panic(fmt.Sprintf(errs.ErrConfigurationColumnMismatch.Error(), len(Configuration.DashboardURIOperation), NoEntries, "DashboardURIOperation"))
	}
	if NoEntries != len(Configuration.DashboardURISuccess) {
		panic(fmt.Sprintf(errs.ErrConfigurationColumnMismatch.Error(), len(Configuration.DashboardURISuccess), NoEntries, "DashboardURISuccess"))
	}

	Configuration.DashboardURIValidActions = []string{"PING", "HTTP"}
	Configuration.DashboardOrdering = buildOrder(Configuration.DashboardOrderIN)
	if NoEntries != len(Configuration.DashboardOrdering) {
		panic(fmt.Sprintf(errs.ErrConfigurationColumnMismatch.Error(), len(Configuration.DashboardOrdering), NoEntries, "DashboardOrdering"))
	}
	Configuration.DashboardDefaultPort = "80"
	Configuration.TimeStampFormat = "20060102"
	Configuration.OnlyFansDateTimeFormat = "060102150405"

	width, height, err := term.GetSize(0)
	if err != nil {
		return
	}

	Configuration.TerminalWidth = width
	Configuration.TerminalHeight = height
	//spew.Dump(Configuration)
}

func split(s string) (r []string) {
	return strings.Split(s, "|")
}

func buildOrder(in string) (r []int) {
	s := strings.Split(in, "|")
	r = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[i], _ = strconv.Atoi(string(s[i]))
	}
	return
}

func (config *Config) Dump() {
	spew.Dump(config)
}

func (config *Config) GetPlexURI() string {
	return config.PlexURI
}

func (config *Config) GetPlexPort() string {
	return config.PlexPort
}

func (config *Config) GetPlexToken() string {
	return config.PlexToken
}

func (config *Config) GetPlexClientID() string {
	return config.PlexClientID
}

func (config *Config) GetPlexDateFormat() string {
	return config.PlexDateFormat
}

func (config *Config) GetApplicationDateFormat() string {
	return config.ApplicationDateFormat
}

func (config *Config) GetApplicationDateFormatShort() string {
	return config.ApplicationDateFormatShort
}

func (config *Config) GetApplicationTimeFormat() string {
	return config.ApplicationTimeFormat
}

func (config *Config) GetTerminalWidth() int {
	return config.TerminalWidth
}

func (config *Config) GetTerminalHeight() int {
	return config.TerminalHeight
}

func (config *Config) GetDelay() float64 {
	return config.Delay
}

func (config *Config) GetBaud() int {
	return config.Baud
}

func (config *Config) GetTransmissionURI() string {
	return config.TransmissionURI
}

func (config *Config) GetQTorrentURI() string {
	return config.QTorrentURI
}

func (config *Config) GetMaxContentRows() int {
	return config.MaxContentRows
}

func (config *Config) GetMaxNoItems() int {
	return config.MaxNoItems
}

func (config *Config) GetTitleLength() int {
	return config.TitleLength
}

func (config *Config) GetDebug() bool {
	return config.Debug
}

func (config *Config) GetOpenWeatherMapApiKey() string {
	return config.OpenWeatherMapApiKey
}

func (config *Config) GetOpenWeatherMapApiLang() string {
	return config.OpenWeatherMapApiLang
}

func (config *Config) GetOpenWeatherMapApiUnits() string {
	return config.OpenWeatherMapApiUnits
}

func (config *Config) GetLocationLogitude() float64 {
	return config.LocationLogitude
}

func (config *Config) GetLocationLatitude() float64 {
	return config.LocationLatitude
}

func (config *Config) GetURISkyNews() string {
	return config.URISkyNews
}

func (config *Config) GetURISkyNewsHome() string {
	return config.URISkyNewsHome
}

func (config *Config) GetURISkyNewsUK() string {
	return config.URISkyNewsUK
}

func (config *Config) GetURISkyNewsWorld() string {
	return config.URISkyNewsWorld
}

func (config *Config) GetURISkyNewsUS() string {
	return config.URISkyNewsUS
}

func (config *Config) GetURISkyNewsBusiness() string {
	return config.URISkyNewsBusiness
}

func (config *Config) GetURISkyNewsPolitics() string {
	return config.URISkyNewsPolitics
}

func (config *Config) GetURISkyNewsTechnology() string {
	return config.URISkyNewsTechnology
}

func (config *Config) GetURISkyNewsEntertainment() string {
	return config.URISkyNewsEntertainment
}

func (config *Config) GetURISkyNewsStrange() string {
	return config.URISkyNewsStrange
}

func (config *Config) GetDefaultErrorDelay() float64 {
	return config.DefaultErrorDelay
}

func (config *Config) GetDefaultRandomPortMin() int {
	return config.DefaultRandomPortMin
}

func (config *Config) GetDefaultRandomPortMax() int {
	return config.DefaultRandomPortMax
}

func (config *Config) GetDefaultRandomMACMin() int {
	return config.DefaultRandomMACMin
}

func (config *Config) GetDefaultRandomMACMax() int {
	return config.DefaultRandomMACMax
}

func (config *Config) GetDefaultRandomIPMin() int {
	return config.DefaultRandomIPMin
}

func (config *Config) GetDefaultRandomIPMax() int {
	return config.DefaultRandomIPMax
}

func (config *Config) GetDefaultBaud() int {
	return config.DefaultBaud
}

func (config *Config) GetDefaultBeepDuration() int {
	return config.DefaultBeepDuration
}

func (config *Config) GetDefaultBeepFrequency() float64 {
	return config.DefaultBeepFrequency
}

func (config *Config) GetValidBaudRates() []int {
	return config.ValidBaudRates
}

func (config *Config) GetValidFileNameCharacters() []string {
	return config.ValidFileNameCharacters
}

func (config *Config) GetDashboardURINameIN() string {
	return config.DashboardURINameIN
}

func (config *Config) GetDashboardURIProtocolIN() string {
	return config.DashboardURIProtocolIN
}

func (config *Config) GetDashboardURIHostIN() string {
	return config.DashboardURIHostIN
}

func (config *Config) GetDashboardURIPortIN() string {
	return config.DashboardURIPortIN
}

func (config *Config) GetDashboardURIQueryIN() string {
	return config.DashboardURIQueryIN
}

func (config *Config) GetDashboardURIOperationIN() string {
	return config.DashboardURIOperationIN
}

func (config *Config) GetDashboardURISuccessIN() string {
	return config.DashboardURISuccessIN
}

func (config *Config) GetDashboardOrderIN() string {
	return config.DashboardOrderIN
}

func (config *Config) GetDashboardDefaultHost() string {
	return config.DashboardDefaultHost
}

func (config *Config) GetDashboardDefaultPort() string {
	return config.DashboardDefaultPort
}

func (config *Config) GetDashboardURIName() []string {
	return config.DashboardURIName
}

func (config *Config) GetDashboardURIProtocol() []string {
	return config.DashboardURIProtocol
}

func (config *Config) GetDashboardURIHost() []string {
	return config.DashboardURIHost
}

func (config *Config) GetDashboardURIPort() []string {
	return config.DashboardURIPort
}

func (config *Config) GetDashboardURIQuery() []string {
	return config.DashboardURIQuery
}

func (config *Config) GetDashboardURIOperation() []string {
	return config.DashboardURIOperation
}

func (config *Config) GetDashboardURISuccess() []string {
	return config.DashboardURISuccess
}

func (config *Config) GetDashboardURIValidActions() []string {
	return config.DashboardURIValidActions
}

func (config *Config) GetDashboardURINoEntries() int {
	return config.DashboardURINoEntries
}

func (config *Config) GetDashboardOrdering() []int {
	return config.DashboardOrdering
}

func (config *Config) GetTimeStampFormat() string {
	return config.TimeStampFormat
}

func (config *Config) GetOnlyFansDateTimeFormat() string {
	return config.OnlyFansDateTimeFormat
}

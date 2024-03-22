package language

// General
// TxtPlexTitleLabel represents the label for the title of a Plex item.
const (
	//	ErrLibraryResponse    string = "library fetch error"
	TxtPlexTitleLabel            string = "Title"   // TODO: Change to Title
	TxtPlexSummaryLabel          string = "Summary" // TODO: Change to Summary
	TxtPlexContainerLabel        string = "Container"
	TxtPlexResolutionLabel       string = "Resolution"
	TxtPlexCodecLabel            string = "Codec"
	TxtPlexAspectRatioLabel      string = "Aspect Ratio"
	TxtPlexFrameRateLabel        string = "Frame Rate"
	TxtPlexDurationLabel         string = "Duration"
	TxtPlexReleasedLabel         string = "Released"
	TxtPlexDirectorLabel         string = "Director"
	TxtPlexWriterLabel           string = "Writer"
	TxtPlexMediaLabel            string = "Media"
	TxtPlexContentRatingLabel    string = "Content Rating"
	TxtPlexTitle                 string = "PMS"
	TxtYouSelected               string = "you selected: "
	TxtYear                      string = "Year"
	TxtPlexSeasons               string = "Seasons - "
	TxtPlexShow                  string = "Show"
	TxtPlexSeason                string = "Season"
	TxtPlexEpisode               string = "Episode"
	SymDelimiter                 string = " - "
	TxtDone                      string = "DONE"
	TxtStarting                  string = "Starting..."
	TxtStartingTerminal          string = "Starting Terminal..."
	TxtSelfTesting               string = "Self Testing..."
	TxtCurrentDate               string = "Current Date: "
	TxtCurrentTime               string = "Current Time: "
	TxtPleaseWait                string = "Please Wait..."
	TxtBaudRate                  string = "Baud Rate Set to %v kbps"
	TxtConnecting                string = "Connecting..."
	TxtDialing                   string = "Dialing... %v:%v"
	TxtConnected                 string = "Connected."
	ErrorMessageConnectionFailed string = "Connection failed. Retrying..."
	TxtComplete                  string = "Complete"
	SymNewline                   string = "\n"
	SymCarridgeReturn            string = "\r"
	SymTab                       string = "\t"
	SymDoubleQuote               string = "\""
	SymSingleQuote               string = "'"
	SymOpenBracket               string = "("
	SymCloseBracket              string = ")"
)

// Mainmenu

const (
	TxtMainMenuTitle                string = "Main Menu"
	TxtSkyNewsMenuTitle             string = "SKY News"
	TxtBBCNewsMenuTitle             string = "BBC News"
	TxtWeatherMenuTitle             string = "Weather"
	TxtRemoteSystemsAccessMenuTitle string = "Remote Systems Access"
	TxtSystemsMaintenanceMenuTitle  string = "Systems Maintenance"
	TxtPlexMediaServersMenuTitle    string = "Plex Media Server"
	TxtQuittingMessage              string = "Quitting"
	TxtSubMenuTitle                 string = "Sub Menu"
	SymBlank                        string = "-"
	TxtTorrentsMenuTitle            string = "Torrents"
)

// SkyNews
const (
	TxtMenuTitle          string = "SKY News"
	TxtTopicHome          string = "Home"
	TxtTopicUK            string = "UK"
	TxtTopicWorld         string = "World"
	TxtTopicUS            string = "US"
	TxtTopicBusiness      string = "Business"
	TxtTopicPolitics      string = "Politics"
	TxtTopicTechnology    string = "Technology"
	TxtTopicEntertainment string = "Entertainment"
	TxtTopicStrange       string = "Strange News"
	TxtLoadingTopic       string = "Loading news for topic: "
	TxtLoadingStory       string = "Loading news for story..."
	HTMLTagTitle          string = "title"
	HTMLTagTagP           string = "p"
)

// Torrents
const (
	TxtTransmission                string = "Transmission"
	TxtQTorrent                    string = "qTorrent"
	TxtLoadingTorrentsTransmission string = "Loading Transmission Torrents..."
	TxtLoadingTorrentsQTor         string = "Loading qTorrent Torrents..."
)

// Weather
const (
	TxtWeatherTitle       string = "Weather"
	TxtWeatherPrompt      string = "Select (Q)uit"
	SymWeatherFormat2     string = "%-25s | %-15v"
	SymWeatherFormat4     string = "%-25s | %-15v  %-15s : %-15v"
	SymWeatherFormat1     string = "%-25s | %v%%"
	TxtLocationLabel      string = "Location"
	TxtConditionsLabel    string = "Conditions"
	TxtTemperatureLabel   string = "Temperature"
	TxtFeelsLikeLabel     string = "Feels Like"
	TxtMinLabel           string = "Min"
	TxtMaxLabel           string = "Max"
	TxtWindSpeedLabel     string = "Wind Speed"
	TxtWindDirectionLabel string = "Wind Direction"
	TxtCloudCoverLabel    string = "Cloud Cover"
	TxtRainLabel          string = "Rain"
	TxtSnowLabel          string = "Snow"
	TxtSunriseLabel       string = "Sunrise"
	TxtSunsetLabel        string = "Sunset"
	TxtSourceLabel        string = "Source"
	TxtSourceService      string = "OpenWeatherMap"
	SymDegree             string = "°"
	SymBreak              string = " ━━ "
	TxtOneHour            string = " (1hr)"
	TxtThreeHour          string = " (3hr)"
	Space                 string = " "
)

var TxtRain1Hr string = TxtRainLabel + TxtOneHour
var TxtRain3Hr string = TxtRainLabel + TxtThreeHour
var TxtSnow1Hr string = TxtSnowLabel + TxtOneHour
var TxtSnow3Hr string = TxtSnowLabel + TxtThreeHour

// Page - Paging

const TxtPagingPrompt string = "Choose (F)orward, (B)ack or (Q)uit"

const (
	SymActionQuit    string = "Q"
	SymActionForward string = "F"
	SymActionBack    string = "B"
	SymActionExit    string = "EX"
	SymTruncate      string = "..."
	SymWildcardBlank string = "{{blank}}"
)

// Plex Seasons

const (
	SymActionSeasons     string = "S"
	TxtPlexSeasonsPrompt string = "Choose (S)easons, (F)orward, (B)ack or (Q)uit"
)

// Support

const (
	TxtOneWord      string = "one"
	TxtOneNumeric   string = "1"
	TxtMinutes      string = "minutes"
	TxtMinutesShort string = "mins"
	TxtHour         string = "hour"
	TxtHourShort    string = "hr"
)

const (
	TxtMillisecondsShort  string = "ms"
	TxtApplicationVersion string = "StarTerm - Utilities 1.0 %s"
	TxtApplicationName    string = "StarTerm"
	SymPromptSymbol       string = "? "
	TxtError              string = "ERROR : "
	TxtInfo               string = "INFO : "
	TxtPaging             string = "Page %v of %v"
)

var ApplicationHeader []string = []string{
	"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
	"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
	"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
	"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
	"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
}

// General
const (
	BoxCharacterNormal      string = "┃"
	BoxCharacterBreak       string = "┣"
	BoxCharacterStart       string = "┏"
	BoxCharacterBar         string = "━"
	BoxCharacterBarBreak    string = "┗"
	TableCharacterUnderline string = "-"
	TextStyleBold           string = "\033[1m"
	TextStyleReset          string = "\033[0m"
	TextStyleUnderline      string = "\033[4m"
	TextColorRed            string = "\033[31m"
	ConsoleClearLine        string = "\033[2K"
	TextLineConstructor     string = "%s%s%s"
	MACAddressConstructor   string = "%v:%v:%v:%v:%v:%v"
	IPAddressConstructor    string = "%v.%v.%v.%v"
)

// Dashboard
const (
	TxtDashboardTitle           string = "Server Dashboard"
	TxtStatusOnline             string = "Online "
	TxtStatusOffline            string = "Offline"
	TxtDashboardChecking        string = "Please wait... Checking Services"
	TxtDashboardCheckingService string = "Checking %v..."
	TxtNoResponseFromServer     string = "No response from server"
)

// Disksize
const (
	TxtDiskSizeTitle     string = "Disk Size"
	TxtDiskSizePrompt    string = "Choose (Q)uit"
	TxtFileSizesOfNFiles string = "File Sizes of %v files"
	DUOutputConstructor  string = "%5d b | %5.2f kb | %5.2f mb | %5.2f gb | %5.2f tb | %s"
)

// CleanFileNames
const (
	TxtStartingCleanFileNames string = "Starting file name cleanse"
	TxtNoFilesFoundInFolder   string = "No files found in folder %s\n"
	TxtProcessingNFilesIn     string = "Processing %d files in %v"
	TxtProcessedNFilesIn      string = "Cleaned %d filenames in %s"
	TxtNoFilesProcessed       string = "No files cleaned in %s"
	TxtOnlyFans               string = "OnlyFans"
	FileExtensionMP4          string = ".mp4"
	TxtOnlyFansFilename       string = TxtOnlyFans + FileExtensionMP4
	TxtRemamedFile            string = "Renamed file [%s -> %s]"
)

// TidyFiles
const (
	TxtTidyFilesTitle                        string = "Media Management - Tidy Folders"
	TxtStorageReportTitle                    string = "STORAGE REPORT"
	TxtTidyFilesBefore                       string = "BEFORE  : %v available out of %v total (%vperc used)"
	TxtTidyFilesAfter                        string = "AFTER   : %v available out of %v total (%vperc used)"
	TxtTidyFilesMachine                      string = "MACHINE : %v"
	TxtTidyFilesHost                         string = "HOST    : %v"
	TxtTidyFilesUser                         string = "USER    : %v"
	TxtTidyFilesMode                         string = "MODE    : %v"
	TxtTidyFilesTypes                        string = "TYPES   : %v"
	TxtTidyFilesEnd                          string = "END     : %v"
	TxtResolvedPath                          string = "Resolved path: %v"
	TxtDebugMode                             string = "DEBUG"
	TxtLiveMode                              string = "LIVE"
	TxtTidyFilesWouldHaveRemoved             string = TxtDebugMode + ": Would have removed empty directories"
	TxtRemovingFilesWithExt                  string = "Removing all files with extension .%v"
	TxtFindingFilesWithExt                   string = "Finding all files with extension .%v"
	TxtOperationComplete                     string = "Operation on .%v files completed in %v"
	TxtTidyFilesStart                        string = "Starting file removal process for %v"
	TxtAreYouSureYouWantToProceed            string = "Are you sure you want to proceed"
	OptAreYouSureYouWantToProceed            string = "y/n"
	TxtChangingDirectory                     string = "Changing directory to %v"
	TxtTidyFilesDeletingDirectories          string = "Deleting empty directories"
	TxtTidyFilesDeletingDirectoriesCompleted string = "Deleting empty directories completed in %v seconds"
	TxtCommandRun                            string = "Command run: %v"
	TxtLiveRun                               string = "This is a live run. Files will be deleted."
	TxtTrailRun                              string = "This is a trial run. Files & Folders will not be deleted."
)

const (
	TxtPushoverTitle string = "Pushover Messaging Service"
)

var (
	TxtServiceMenuDescription       []string = []string{"This menu shows the list of services available for maintenance.", "Select the service you wish to use. PLEASE BE CAREFUL!"}
	TxtPushoverDescription          []string = []string{"Pushover is a service to receive instant push notifications on your phone or tablet from a variety of sources."}
	TxtPushoverMsgPriorityEmergancy string   = "Emergancy Message"
	TxtPushoverMsgPriorityNormal    string   = "Normal Priority"
	TxtPushoverMsgPriorityHigh      string   = "High Priority"
	TxtPushoverMsgPriorityLow       string   = "Low Priority"
	TxtPushoverPrompt               string   = "Choose a message type to send"
	TxtPushoverConfirmation         string   = "Choose (S)end or (Q)uit"
)

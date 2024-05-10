package language

import (
	lang "github.com/mt1976/crt/language"
)

// General
var (
	YouSelected *lang.Text = lang.New("you selected: ")
	Done        *lang.Text = lang.New("DONE")
	LiveRun     *lang.Text = lang.New("This is a live run. PLEASE BE CAREFUL!")
	TrailRun    *lang.Text = lang.New("This is a trial run.")
	Mode        *lang.Text = lang.New("Mode")
	Path        *lang.Text = lang.New("Path")
)

// PLEX represents the label for the title of a Plex item.
var (
	PlexTitle         *lang.Text = lang.New("PMS") // TODO: Change to Title
	PlexSummary       *lang.Text = lang.New("Summary")
	PlexContainer                = lang.New("Container")
	PlexResolution               = lang.New("Resolution")
	PlexCodec                    = lang.New("Codec")
	PlexAspectRatio              = lang.New("Aspect Ratio")
	PlexFrameRate                = lang.New("Frame Rate")
	PlexDuration                 = lang.New("Duration")
	PlexReleased                 = lang.New("Released")
	PlexDirector                 = lang.New("Director")
	PlexWriter                   = lang.New("Writer")
	PlexMedia                    = lang.New("Media")
	PlexContentRating            = lang.New("Content Rating")
	PlexShow                     = lang.New("Show")
	PlexSeason                   = lang.New("Season")
	PlexEpisode                  = lang.New("Episode")
	PlexSeasons                  = lang.New("Seasons")
	PlexEpisodes                 = lang.New("Episodes")
	PlexYear                     = lang.New("Year")
)

var (
	//	PlexSummaryLabel             string = "Summary"
	//TxtPlexContainerLabel        string = "Container"
	//TxtPlexResolutionLabel       string = "Resolution"
	//TxtPlexCodecLabel            string = "Codec"
	// TxtPlexAspectRatioLabel      string = "Aspect Ratio"
	// TxtPlexFrameRateLabel        string = "Frame Rate"
	// TxtPlexDurationLabel         string = "Duration"
	// TxtPlexReleasedLabel         string = "Released"
	// TxtPlexDirectorLabel         string = "Director"
	// TxtPlexWriterLabel           string = "Writer"
	// TxtPlexMediaLabel            string = "Media"
	// TxtPlexContentRatingLabel    string = "Content Rating"
	// TxtPlexTitle   string = "PMS"
	// TxtYouSelected string = "you selected: "
	// TxtYear        string = "Year"
	// TxtPlexSeasons               string = "Seasons - "
	// TxtPlexShow                  string = "Show"
	// TxtPlexSeason                string = "Season"
	// TxtPlexEpisode               string = "Episode"
	//TxtDone                      string = "DONE"
	//TxtStarting                  string     = "Starting..."
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
	TxtThankYouForUsing             string = "Thank you for using"
	TxtSubMenuTitle                 string = "Sub Menu"
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

// BBC
const (
	TxtBBCError string = "error with BBC News"
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

// Plex Seasons
const (
	TxtPlexSeasonsPrompt string = "Choose (1...n)Season, (F)orward, (B)ack or (Q)uit"
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
)

var ApplicationHeader []string = []string{
	"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
	"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
	"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
	"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
	"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
}

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
	TxtDiskSizeResults   string = "Disk Size - Results"
	TxtDiskSizeReport    string = "Disk Size - Report"
	TxtDiskSizePrompt    string = "Choose (Q)uit"
	TxtFileSizesOfNFiles string = "File Sizes of %v files"
	DUOutputConstructor  string = "%5d b | %5.2f kb | %5.2f mb | %5.2f gb | %5.2f tb | %s"
)

// CleanFileNames
var (
	TxtCleanFileNames           string          = "Clean File Names"
	TxtCleanFileNamesResults    string          = "Clean File Names - Results"
	TxtCleanFileNamesReport     string          = "Clean File Names - Report"
	TxtStartingCleanFileNames   string          = "Starting file name cleanse"
	TxtNoFilesFoundInFolder     string          = "No files found in folder %s\n"
	TxtProcessingNFilesIn       string          = "Processing %d files in %v"
	TxtProcessedNFilesIn        string          = "Cleaned %d filenames in %s"
	TxtNoFilesProcessed         string          = "No files cleaned in %s"
	TxtOnlyFans                 string          = "OnlyFans"
	FileExtensionMP4            string          = ".mp4"
	TxtOnlyFansFilename         string          = TxtOnlyFans + FileExtensionMP4
	TxtRemamedFile              string          = "Renamed file [%s -> %s]"
	TxtProcessing               string          = "Processing %v type files"
	TxtRemovingEmptyDirectories string          = "Removing empty directories"
	TxtFindingEmptyDirectories  string          = "Finding empty directories"
	CleanFileNamesDescription   *lang.Paragraph = lang.NewParagraph([]string{"This menu shows the list of files available for maintenance.", "Select the file you wish to use. PLEASE BE CAREFUL!"})
)

// TidyFiles
const (
	TxtTidyFilesTitle                        string = "Clean Media Folders"
	TxtTidyFilesTitleResults                 string = "Clean Media Folders - Results"
	TxtStorageReportTitle                    string = "Clean Media Folders - Report"
	TxtTidyFilesBefore                       string = "BEFORE  : %v available out of %v total (%vperc used)"
	TxtTidyFilesAfter                        string = "AFTER   : %v available out of %v total (%vperc used)"
	TxtTidyFilesMachine                      string = "MACHINE : %v"
	TxtTidyFilesHost                         string = "HOST    : %v"
	TxtTidyFilesUser                         string = "USER    : %v"
	TxtTidyFilesMode                         string = "MODE    : %v"
	TxtTidyFilesTypes                        string = "TYPES   : %v"
	TxtTidyFilesEnd                          string = "END     : %v"
	TxtResolvedPath                          string = "Resolved path: %v"
	TxtDebugMode                             string = "TRIAL"
	TxtLiveMode                              string = "LIVE"
	TxtTidyFilesWouldHaveRemoved             string = TxtDebugMode + ": Would have removed empty directories"
	TxtRemovingFilesWithExt                  string = "Removing all files with extension .%v"
	TxtFindingFilesWithExt                   string = "Finding all files with extension .%v"
	TxtOperationComplete                     string = "Operation on .%v files completed in %v"
	TxtOperationCompleteIncomplete           string = "Operation completed in %v"
	TxtTidyFilesStart                        string = "Starting file removal process for %v"
	TxtFileTypes                             string = "File Types"
	TxtAreYouSureYouWantToProceed            string = "Are you sure you want to proceed"
	OptAreYouSureYouWantToProceed            string = "y/n"
	TxtChangingDirectory                     string = "Changing directory to %v"
	TxtTidyFilesDeletingDirectories          string = "Deleting empty directories"
	TxtTidyFilesDeletingDirectoriesCompleted string = "Deleting empty directories completed in %v seconds"
	TxtCommandRun                            string = "Command run: %v"
	// TxtLiveRun                               string = "This is a live run. Files will be deleted."
	// TxtTrailRun                              string = "This is a trial run. Files & Folders will not be deleted."
	// TxtMode                                  string = "Mode"
	// TxtPath                                  string = "Path"
)

var (
	TxtPushoverTitle                string   = "Pushover Messaging Service"
	TxtServiceMenuDescription       []string = []string{"This menu shows the list of services available for maintenance.", "Select the service you wish to use. PLEASE BE CAREFUL!"}
	TxtPushoverDescription          []string = []string{"Pushover is a service to receive instant push notifications on your phone or tablet from a variety of sources."}
	TxtPushoverMsgPriorityEmergency string   = "Emergency Message"
	TxtPushoverMsgPriorityNormal    string   = "Normal Priority"
	TxtPushoverMsgPriorityHigh      string   = "High Priority"
	TxtPushoverMsgPriorityLow       string   = "Low Priority"
	TxtPushoverPrompt               string   = "Choose a message type to send"
	TxtPushoverConfirmation         string   = "Choose (S)end or (Q)uit"
	TxtPushoverMessageTitlePrompt   string   = "Enter the title of the message, or (Q)uit"
	TxtPushoverMessageBodyPrompt    string   = "Enter the body of the message, or (Q)uit"
	TxtPushoverMessageSending       string   = "Sending Pushover Message"
	TxtPushoverMessageSent          string   = "Pushover Message Sent"
)

// FileChooser
var (
	TxtFileChooserTitle        string   = "File Chooser"
	TxtFileChooserDescription  []string = []string{"This menu shows the list of files available for maintenance.", "Select the file you wish to use. PLEASE BE CAREFUL!"}
	TxtFileChooserPrompt       string   = "Choose a file to use"
	TxtFileChooserConfirmation string   = "Choose (S)end or (Q)uit"
)

// Catalog
var (
	CatalogTitle        *lang.Text      = lang.New("Systems Catalog")
	CatalogPrompt       *lang.Text      = lang.New("Choose a file to use")
	CatalogConfirmation *lang.Text      = lang.New("Choose (S)end or (Q)uit")
	CatalogDescription  *lang.Paragraph = lang.NewParagraph([]string{"This menu shows the list of files available for maintenance.", "Select the file you wish to use. PLEASE BE CAREFUL!"})
)

// File Migrator
var (
	FileMigratorTitle            *lang.Text = lang.New("File Migration")
	FileMigratorMode             *lang.Text = lang.New("Mode")
	FileMigratorModeCheckPrompt  *lang.Text = lang.New("Are you sure you want to proceed, select (Y) to continue...")
	FileMigratorFile             *lang.Text = lang.New("File")
	FileMigratorNoFilesToProcess *lang.Text = lang.New("No Files")
	FileMigratorDestination      *lang.Text = lang.New("Destination")
	FileMigratorResults          *lang.Text = lang.New("RESULTS")
	FileMigratorDonePrompt       *lang.Text = lang.New("Processing Complete, select (Y) to continue...")
	FileMigratorMoving           *lang.Text = lang.New("Moving [%d/%d] [%20v][%20v]")
	FileMigratorMovingArrow      *lang.Text = lang.New(" -> ")
)

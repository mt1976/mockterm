package language

import (
	lang "github.com/mt1976/crt/language"
)

// General
// TxtPlexTitleLabel represents the label for the title of a Plex item.
var (
	PlexTitleLabel *lang.Text = lang.New("Title") // TODO: Change to Title
)

const (
	PlexSummaryLabel             string = "Summary"
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
	TxtLiveRun                               string = "This is a live run. Files will be deleted."
	TxtTrailRun                              string = "This is a trial run. Files & Folders will not be deleted."
	TxtMode                                  string = "Mode"
	TxtPath                                  string = "Path"
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
	TxtCatalogTitle        string   = "Systems Catalog"
	TxtCatalogPrompt       string   = "Choose a file to use"
	TxtCatalogConfirmation string   = "Choose (S)end or (Q)uit"
	TxtCatalogDescription  []string = []string{"This menu shows the list of files available for maintenance.", "Select the file you wish to use. PLEASE BE CAREFUL!"}
)

// File Migrator
const (
	TxtFileMigratorTitle            string = "File Migration"
	TxtFileMigratorMode             string = "Mode"
	TxtFileMigratorModeCheckPrompt  string = "Are you sure you want to proceed, select (Y) to continue..."
	TxtFileMigratorFile             string = "File"
	TxtFileMigratorNoFilesToProcess string = "No Files"
	TxtFileMigratorDestination      string = "Destination"
	TxtFileMigratorResults          string = "RESULTS"
	TxtFileMigratorDonePrompt       string = "Processing Complete, select (Y) to continue..."
	TxtFileMigratorMoving           string = "Moving [%d/%d] [%20v][%20v]"
	TxtFileMigratorMovingArrow      string = " -> "
)

package language

import lang "github.com/mt1976/crt/language"

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

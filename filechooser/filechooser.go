package filechooser

import (
	"fmt"
	"os"
	"strings"

	"github.com/mt1976/crt"
	lang "github.com/mt1976/mockterm/language"
)

type File struct {
	Name     string
	Path     string
	Created  string
	Modified string
	Size     int64
	SizeTxt  string
	IsDir    bool
	Icon     string
	Mode     string
	Seq      int
}

func FileChooser(root string, includeDotFiles, includeDirectories, showFiles bool) (string, bool, error) {

	term := crt.New()
	page := term.NewPage(lang.TxtFileChooserTitle)

	files, err := GetFolderList(root, includeDotFiles, includeDirectories, showFiles)
	if err != nil {
		return "", false, err
	}
	uh, _ := UserHome()
	page.AddFieldValuePair("Home", uh)
	page.AddFieldValuePair("Directory", root)
	page.AddBlankRow()
	//page.AddColumnsTitle("Name", "Mode", "Size", "Modified")
	formatter := "%1v %-30v %10v %12v %15v"
	tformatter := "%6v %1v|%-30v|%-10v|%-12v|%-15v"
	title := fmt.Sprintf(tformatter, "", "T", " Name", " Mode", " Modified", " Size")
	title = strings.ReplaceAll(title, "|", " ")
	page.Add(title, "", "")
	breaker := fmt.Sprintf(tformatter, strings.Repeat(" ", 5), strings.Repeat("-", 1), strings.Repeat("-", 30), strings.Repeat("-", 10), strings.Repeat("-", 12), strings.Repeat("-", 15))
	//breaker = strings.Repeat("-", len(breaker))
	page.Add(breaker, "", "")
	for _, file := range files {
		row := fmt.Sprintf(formatter, file.Icon, file.Name, file.Mode, file.Modified, file.SizeTxt)
		page.AddMenuOption(file.Seq+1, row, "", "")
	}
	na, _ := page.DisplayWithActions()
	if na == lang.SymActionQuit {
		return "", false, nil
	}
	if term.Helpers.IsInt(na) {
		r := files[term.Helpers.ToInt(na)-1]
		return r.Path, r.IsDir, nil
	}
	return "", false, nil
}

func UserHome() (string, error) {
	// Function gets the home directory of the current user, or returns an error if it cant.
	//
	// Returns:
	// The home directory of the current user, or an error if it cant.
	return os.UserHomeDir()
}

func GetFolderList(dir string, includeDotFiles, includeDirectories, showFiles bool) ([]File, error) {
	// Get a list of files in the specified directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Filter the list of files to only include directories
	var directories []File
	include := false
	itemNo := 0
	for _, file := range files {
		if file.IsDir() && includeDirectories {
			include = true
		}
		if file.Name()[0] == '.' && !includeDotFiles {
			include = false
			continue
		}
		if !file.IsDir() && showFiles {
			include = true
		}
		if include {
			var this File
			this.Name = file.Name()
			this.Path = dir + "/" + file.Name()
			inf, _ := file.Info()
			this.Created = "N/A"
			this.Modified = crt.New().Formatters.HumanFromUnixDate(inf.ModTime().Local().Unix())
			this.Size = inf.Size()
			yy := fmt.Sprintf("%v\n", this.Size)
			this.SizeTxt = yy + "b"
			this.Mode = inf.Mode().String()
			this.IsDir = file.IsDir()
			if this.IsDir {
				this.Icon = lang.TxtFolderIcon
			} else {
				this.Icon = lang.TxtFileIcon
			}
			if isSymLink(this.Mode) {
				this.Icon = lang.TxtSymLinkIcon
			}
			this.Icon = this.Icon + " "
			this.Seq = itemNo
			directories = append(directories, this)
			itemNo++
			//directories = append(directories, "📁 "+file.Name())
		}
	}

	return directories, nil
}

func isSymLink(mode string) bool {
	return mode[0] == 'L' || mode[0] == 'l'
}

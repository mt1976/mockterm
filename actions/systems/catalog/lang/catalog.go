package language

import lang "github.com/mt1976/crt/language"

// Catalog
var (
	CatalogTitle        *lang.Text      = lang.New("Systems Catalog")
	CatalogPrompt       *lang.Text      = lang.New("Choose a file to use")
	CatalogConfirmation *lang.Text      = lang.New("Choose (S)end or (Q)uit")
	CatalogDescription  *lang.Paragraph = lang.NewParagraph([]string{"This menu shows the list of files available for maintenance.", "Select the file you wish to use. PLEASE BE CAREFUL!"})
)

package language

import lang "github.com/mt1976/crt/language"

// Catalog
var (
	Title        *lang.Text      = lang.New("Systems Catalog")
	Cataloging   *lang.Text      = lang.New("Cataloging System Resources")
	Prompt       *lang.Text      = lang.New("Choose a file to use")
	Confirmation *lang.Text      = lang.New("Choose (S)end or (Q)uit")
	Description  *lang.Paragraph = lang.NewParagraph([]string{"This menu shows the list of files available for maintenance.", "Select the file you wish to use. PLEASE BE CAREFUL!"})
	Proceed      *lang.Text      = lang.New("Do you want to continue with the cataloging process")
	Quitting     *lang.Text      = lang.New("Quitting")
	Complete     *lang.Text      = lang.New("Cataloging complete")
)

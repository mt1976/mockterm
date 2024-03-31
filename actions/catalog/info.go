package catalog

import (
	"fmt"

	term "github.com/mt1976/crt"
)

type info struct {
	data []string
}

// New returns a new instance of the info struct
func New() info {
	return info{}
}

// breakData adds a new section to the info struct with the given title and adds a line break
func (i *info) breakData(t *term.ViewPort, title string) {
	// add a new section to the info struct
	i.data = append(i.data, t.Format("", ""))
	i.data = append(i.data, t.Format(title, ""))
	i.data = append(i.data, t.Format("", ""))
	// print the given title in bold
	t.Print(t.Formatters.Bold(title))
	// add a line break
	t.Break()
	results = append(results, title)
}

// storeData adds a new key-value pair to the info struct as a new section with the given title and data
func (i *info) storeData(t *term.ViewPort, title string, data string) {
	// Padd title to 15 characters
	title = fmt.Sprintf("%-30s", title)
	// add a new section to the info struct
	i.data = append(i.data, t.Format(title+": "+data, ""))
	//print the given title with data in bold
	t.Print(title + ": " + t.Formatters.Bold(data))
	results = append(results, title+": "+data)
}

package catalog

import (
	"fmt"

	term "github.com/mt1976/crt/terminal"
)

type info struct {
	data []string
}

// New returns a new instance of the info struct
func New() *info {
	return &info{}
}

// breakData adds a new section to the info struct with the given title and adds a line break
func breakData(p *term.Page, i info, title string) (info, term.Page) {
	t := p.ViewPort()
	// add a new section to the info struct
	i.data = append(i.data, t.Format("", ""))
	i.data = append(i.data, t.Format(title, ""))
	i.data = append(i.data, t.Format("", ""))
	// print the given title in bold
	p.Add(t.Formatters.Bold(title), "", "")
	// add a line break
	p.AddBlankRow()
	results = append(results, title)
	return i, *p
}

// storeData adds a new key-value pair to the info struct as a new section with the given title and data
func storeData(p *term.Page, i *info, title string, data string) {
	t := p.ViewPort()
	//p.AddFieldValuePair(title, data)
	// Padd title to 15 characters
	title = fmt.Sprintf("%-30s", title)
	// add a new section to the info struct
	msg := t.Format(title+": "+data, "")
	i.data = append(i.data, msg)
	//print the given title with data in bold
	//t.Print(title + ": " + t.Formatters.Bold(data))
	//results = append(results, msg)
}

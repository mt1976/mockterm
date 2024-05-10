package errors

import (
	"errors"
)

var (
	FetchingXML = errors.New("error fetching XML")
	ParsingXML  = errors.New("error parsing XML")
	ParsingDate = errors.New("error parsing date")
)

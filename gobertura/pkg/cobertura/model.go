package cobertura

import (
	"encoding/xml"
)

// Coverage Cobertura XML structure
type Coverage struct {
	XMLName  xml.Name  `xml:"coverage"`
	Sources  []string  `xml:"sources>source"`
	Packages []Package `xml:"packages>package"`
}

type Package struct {
	Name    string  `xml:"name,attr"`
	Classes []Class `xml:"classes>class"`
}

type Class struct {
	Name     string `xml:"name,attr"`
	Filename string `xml:"filename,attr"`
	Lines    []Line `xml:"lines>line"`
}

type Line struct {
	Number int `xml:"number,attr"`
	Hits   int `xml:"hits,attr"`
}

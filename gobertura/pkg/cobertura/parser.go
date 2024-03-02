package cobertura

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseCoberturaXML(filePath string) (*Coverage, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var coverage Coverage
	if err := xml.Unmarshal(byteValue, &coverage); err != nil {
		return nil, fmt.Errorf("error parsing XML: %v", err)
	}

	return &coverage, nil
}

package report

import (
	"buhnik.com/kuc/pkg/cobertura"
	"fmt"
	"io/ioutil"
	"strings"
)

func PrintSourceWithCoverage(sourceFile string, lines []cobertura.Line) (string, error) {
	fileContent, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return "", fmt.Errorf("error reading source file: %v", err)
	}

	sourceLines := map[int]bool{}
	for _, line := range lines {
		sourceLines[line.Number] = line.Hits > 0
	}

	var report string
	for i, line := range strings.Split(string(fileContent), "\n") {
		lineNumber := i + 1
		covered, exists := sourceLines[lineNumber]
		coverageInfo := "miss"
		if exists && covered {
			coverageInfo = "covered"
		}
		report = report + fmt.Sprintf("%d (%s): %s\n", lineNumber, coverageInfo, line)
	}
	return report, nil
}

package main

import (
	"buhnik.com/kuc/pkg/cobertura"
	"buhnik.com/kuc/pkg/report"
	"fmt"
)

func main() {
	coverage, err := cobertura.ParseCoberturaXML("./cobertura_test.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pkg := range coverage.Packages {
		for _, class := range pkg.Classes {
			fmt.Printf("Processing file: %s\n", class.Filename)
			r, err := report.PrintSourceWithCoverage(class.Filename, class.Lines)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(r)
		}
	}
}

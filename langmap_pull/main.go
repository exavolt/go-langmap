//

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// A tool to pull the data from Mozilla's repo and generate the Go code.
//
// https://raw.githubusercontent.com/mozilla/language-mapping-list/master/language-mapping-list.js

//NOTE: ensure that the resulting code has no syntax errors, all correct, no
// formatting errors, no style errors. Ensure no complaints from gofmt, go vet
// and golint.
//
//TODO: actual pull (from github, or even from npm). probably check github
// release.
//TODO: flags (URL)
//TODO: use buffer for the output
//TODO: optional test when done.
//TODO: whitelisting. limit to certain languages.

func main() {
	srcf, err := os.Open("../js/language-mapping-list.js")
	if err != nil {
		log.Fatal(err)
	}
	defer srcf.Close()

	keyRE := regexp.MustCompile(`^\s*'(.*)':\s*{$`)
	nativeNameRE := regexp.MustCompile(`^\s*nativeName\s*:\s*(.*)\s*,\s*$`)
	englishNameRE := regexp.MustCompile(`^\s*englishName\s*:\s*(.*)\s*$`)
	closeBrRE := regexp.MustCompile(`^\s*},?\s*$`)

	outf := os.Stdout

	fmt.Fprintf(outf,
		`// This file is generated. Do not edit directly.

package langmap

// Names contains the actual data.
var Names = map[string]Name{
`)

	started := false
	scanner := bufio.NewScanner(srcf)
	for scanner.Scan() {
		if !started {
			if scanner.Text() == "}(this, function() {" {
				if !scanner.Scan() {
					break
				}
				//TODO: use strings.Contains?
				if scanner.Text() == "  return {" {
					started = true
					continue
				} else {
					break
				}
			}
			continue
		}
		if scanner.Text() == "  };" {
			break
		}
		txt := scanner.Text()
		groups := keyRE.FindStringSubmatch(txt)
		if len(groups) == 2 {
			fmt.Fprintf(outf, "\t\"%s\": {\n", groups[1])
			continue
		}
		groups = nativeNameRE.FindStringSubmatch(txt)
		if len(groups) == 2 {
			fmt.Fprintf(outf, "\t\t%s,\n", groups[1])
			continue
		}
		groups = englishNameRE.FindStringSubmatch(txt)
		if len(groups) == 2 {
			fmt.Fprintf(outf, "\t\t%s,\n", groups[1])
			continue
		}
		if closeBrRE.MatchString(txt) {
			fmt.Fprintf(outf, "\t},\n")
			continue
		}
		log.Fatal("unexpected text:", txt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(outf, "}\n")
}

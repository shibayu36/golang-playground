package main

import (
	"fmt"
	"regexp"
	"strings"
)

// (?:<plugin_name>|<owner>/<repo>)(?:@<releaseTag>)?
var targetReg = regexp.MustCompile(`^(?:([^@/]+)/([^@/]+)|([^@/]+))(?:@(.+))?$`)

func main() {
	fmt.Println(strings.Split("hoge@", "@"))
	matches := targetReg.FindStringSubmatch("hoge/fuga@v0.0.1")
	fmt.Println(len(matches))
	for _, match := range matches {
		fmt.Println(match)
	}

	matches = targetReg.FindStringSubmatch("hoge/fuga")
	fmt.Println(len(matches))
	for _, match := range matches {
		fmt.Println(match)
	}

	matches = targetReg.FindStringSubmatch("fuga@v0.0.0")
	fmt.Println(len(matches))
	for _, match := range matches {
		fmt.Println(match)
	}

	matches = targetReg.FindStringSubmatch("fuga/fuga/fuga@v0.0.0")
	fmt.Println(len(matches))
	for _, match := range matches {
		fmt.Println(match)
	}
}

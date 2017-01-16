package tests

import "regexp"

var (
	testReplRx = regexp.MustCompile(`\$\{([\w]{1,1000})\}`)
)

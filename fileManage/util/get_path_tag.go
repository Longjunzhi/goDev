package util

import (
	"runtime"
	"strings"
)

func GetPathTag() (pathTag string) {
	var concurrentOsTag string
	if strings.Compare("windows", runtime.GOOS) == 0 {
		concurrentOsTag = "\\"
	} else {
		concurrentOsTag = "/"
	}
	return concurrentOsTag
}

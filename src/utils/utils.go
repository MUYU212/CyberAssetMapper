package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func AppendError(existError, newError error) error {
	if existError == nil {
		return newError
	}
	return fmt.Errorf("%v,%w", existError, newError)
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

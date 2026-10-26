package main

import (
	"errors"
	"os"
	"strings"
)

var errBadChar = errors.New("unsupported character")

func isValidBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}

func makeArt(text, banner string) (string, error) {
	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return "", err
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	if len(lines) < 95*9 {
		return "", errors.New("banner file is broken")
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	for _, letter := range text {
		if letter != '\n' && (letter < 32 || letter > 126) {
			return "", errBadChar
		}
	}

	var result strings.Builder
	for _, word := range strings.Split(text, "\n") {
		if word == "" {
			result.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, letter := range word {
				start := (int(letter)-32)*9 + 1
				result.WriteString(lines[start+row])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

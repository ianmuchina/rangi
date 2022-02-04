package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	var themes = make(map[string]VsCodeTheme)
	var b Base16

	for _, extension := range getExtensions() {
		theme := parseThemes(extension)
		for _, theme := range theme {
			themes[theme.Name] = theme
		}
	}

	for _, theme := range themes {
		// Make scheme camel case
		b = vsodeThemeToBase16(theme)
		scheme := strings.Replace(strings.ToLower(b.Scheme), " ", "_", -1)

		filename := path.Join(b.Type, fmt.Sprint(scheme, ".yaml"))

		out := Base16Out(b)
		fmt.Println(filename)
		err := os.WriteFile(filename, []byte(out), 0644)
		if err != nil {
			fmt.Println("error writing", "to", filename)
		} else {
			fmt.Println(filename)
		}
	}
}

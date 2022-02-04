package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	var themes = make(map[string]VsCodeTheme)

	for _, extension := range getExtensions() {
		theme := LoadThemeJson(extension)
		for _, theme := range theme {
			themes[theme.Name] = theme
		}
	}

	//Iterate over every theme
	for _, theme := range themes {
		var b Base16 = vsodeThemeToBase16(theme)
		// Make scheme camel case
		scheme := strings.Replace(strings.ToLower(b.Scheme), " ", "_", -1)
		if strings.Contains(b.Scheme, "dark") {
			b.Type = "dark"
		} else if strings.Contains(b.Scheme, "light") {
			b.Type = "light"
		} else {
			b.Type = "unknown"
		}

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

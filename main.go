package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {

	//Map of Theme names to vscode themes
	var themeMap = make(map[string]VsCodeTheme)

	//Populate Map
	for _, i := range getExtensions() {
		t := LoadThemeJson(i)
		for _, theme := range t {
			themeMap[theme.Name] = theme
		}
	}

	//Iterate over every theme
	for _, t := range themeMap {
		var b Base16 = vsodeThemeToBase16(t)

		// Make scheme camel case

		b.Scheme = strings.ToLower(b.Scheme)

		scheme := strings.Replace(b.Scheme, " ", "_", -1)

		//detect light mode or dark mode
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
		// err := os.Mkdir("./tmp", 0755)
		// if err != nil {
		// 	fmt.Printf("error creating folder %w", err)
		// }

		err := os.Chdir("/tmp/rangi/")

		if err != nil {
			fmt.Println("error changing to folder")
		}

		// fmt.Println(filename)
		err = os.WriteFile(filename, []byte(out), 0644)
		if err != nil {
			fmt.Println("error writing", "to", filename)
		}
	}

}

package main

import (
	"fmt"
)

func main() {

	//Array of vscode themes
	var themes []VsCodeTheme = getAllThemes()

	//Populate Array
	for _, i := range getExtensions() {
		t := LoadThemeJson(i)
		themes = append(themes, t...)
	}

	//Map of theme names to themes
	var themeMap = make(map[string]VsCodeTheme)
	//Populate Map
	for _, t := range themes {
		themeMap[t.Name] = t
	}

	//Test: List all Themes
	// for _, t := range themes {
	// fmt.Println(t.Name)
	// }

	var themeV VsCodeTheme = themes[1]

	//Theme converted to base16
	var themeB Base16 = vsodeThemeToBase16(themeV)

	//Test: Render Output
	out := Base16Out(themeB)

	fmt.Println(out)
}

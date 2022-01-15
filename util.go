package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

//Extensions that Contain themes
func getExtensions() []string {
	result := []string{}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Path To Extensions Folder on Linux
	extensionsDir := path.Join(home, ".vscode/extensions")

	//Open Extensions Folder
	extensions, err := os.ReadDir(extensionsDir)

	//Check Errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open every extension folder
	for _, extension := range extensions {
		// Files contained in extension
		files, err := os.ReadDir(path.Join(extensionsDir, extension.Name()))

		//Check Errors
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//Check if extension has theme folder
		for _, file := range files {
			if file.Name() == "themes" {
				fullPath := path.Join(extensionsDir, extension.Name())
				result = append(result, fullPath)

			}
		}
	}

	return result

}

//parse Package.json and get a list of all theme json files
func LoadThemeJson(s string) []VsCodeTheme {
	var result []VsCodeTheme

	filename := path.Join(s, "package.json")

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var p PackageJson

	err = json.Unmarshal(f, &p)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for _, theme := range p.Contributes.Themes {
		t, err := ioutil.ReadFile(path.Join(s, theme.Path))

		t = UncommentJson(t)

		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		var v VsCodeTheme
		err = json.Unmarshal(t, &v)

		result = append(result, v)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

	}

	return result
}

//Remove Comments from json
func UncommentJson(buf []byte) []byte {
	re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	return re.ReplaceAll(buf, nil)
}

//Output base16 theme as yaml string
// Todo: attach to struct
func Base16Out(b Base16) string {
	var result string
	result += fmt.Sprintf("scheme: \"%s\" \n", b.Scheme)
	result += fmt.Sprintf("type:   \"%s\" \n", b.Type)
	result += fmt.Sprintf("author: \"%s\" \n", b.Author)
	result += fmt.Sprintf("base00: \"%s\" \n", b.Base00)
	result += fmt.Sprintf("base01: \"%s\" \n", b.Base01)
	result += fmt.Sprintf("base02: \"%s\" \n", b.Base02)
	result += fmt.Sprintf("base03: \"%s\" \n", b.Base03)
	result += fmt.Sprintf("base04: \"%s\" \n", b.Base04)
	result += fmt.Sprintf("base05: \"%s\" \n", b.Base05)
	result += fmt.Sprintf("base06: \"%s\" \n", b.Base06)
	result += fmt.Sprintf("base07: \"%s\" \n", b.Base07)
	result += fmt.Sprintf("base08: \"%s\" \n", b.Base08)
	result += fmt.Sprintf("base09: \"%s\" \n", b.Base09)
	result += fmt.Sprintf("base0A: \"%s\" \n", b.Base0A)
	result += fmt.Sprintf("base0B: \"%s\" \n", b.Base0B)
	result += fmt.Sprintf("base0C: \"%s\" \n", b.Base0C)
	result += fmt.Sprintf("base0D: \"%s\" \n", b.Base0D)
	result += fmt.Sprintf("base0E: \"%s\" \n", b.Base0E)
	result += fmt.Sprintf("base0F: \"%s\" \n", b.Base0F)

	return result
}

//Convert a vscode theme to a base16 theme
func vsodeThemeToBase16(v VsCodeTheme) Base16 {
	var b Base16
	//Apply base16 theme
	b.Scheme = v.Name
	b.Type = v.Type
	b.Author = "Anon"
	//Shades of Background to foreground
	//FIXME: Proper colors
	b.Base00 = v.Colors.EditorBackground
	b.Base01 = v.Colors.EditorBackground
	b.Base02 = v.Colors.TerminalAnsiBlack
	b.Base03 = v.Colors.TerminalAnsiBrightBlack
	b.Base04 = v.Colors.EditorForeground
	b.Base05 = v.Colors.EditorForeground
	b.Base06 = v.Colors.EditorForeground
	b.Base07 = v.Colors.EditorForeground

	// Colors
	b.Base08 = v.Colors.TerminalAnsiBrightRed
	b.Base09 = v.Colors.TerminalAnsiBrightYellow
	b.Base0A = v.Colors.TerminalAnsiBrightYellow
	b.Base0B = v.Colors.TerminalAnsiBrightGreen
	b.Base0C = v.Colors.TerminalAnsiBrightCyan
	b.Base0D = v.Colors.TerminalAnsiBrightBlue
	b.Base0E = v.Colors.TerminalAnsiBrightMagenta
	b.Base0F = v.Colors.ErrorForeground

	return b
}

func getAllThemes() []VsCodeTheme {
	var themes []VsCodeTheme

	for _, i := range getExtensions() {
		t := LoadThemeJson(i)
		themes = append(themes, t...)

	}

	return themes
}

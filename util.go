package main

import (
	"encoding/json"

	"github.com/tailscale/hujson"

	"fmt"
	"os"
	"path"
	"strings"
)

//Extensions that Contain themes
func getExtensions() []string {
	var pj PackageJson
	result := []string{}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Path To Extensions Folder on Linux
	extensionsDir := path.Join(home, ".vscode/extensions")
	// Open Extensions Folder
	extensions, err := os.ReadDir(extensionsDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Open every extension folder
	for _, extension := range extensions {
		// Skip entries that are not folders
		if !extension.IsDir() {
			continue
		}
		// Parse Package json to struct using hujson, a JWCC Extension
		f, err := os.ReadFile(path.Join(extensionsDir, extension.Name(), "package.json"))
		if err != nil {
			fmt.Println("error opening package json", err)
			os.Exit(1)
		}
		ast, err := hujson.Parse(f)
		if err != nil {
			fmt.Println("error unmarshaling package json (hujson)", err)
			os.Exit(1)
		}
		ast.Standardize()
		f = ast.Pack()
		err = json.Unmarshal(f, &pj)

		if err != nil {
			fmt.Println("error unmarshaling package json", err)
			os.Exit(1)
		}
		// append extensions with themes
		if len(pj.Contributes.Themes) > 0 {
			fullPath := path.Join(extensionsDir, extension.Name())
			result = append(result, fullPath)
		}
	}
	return result
}

// Unmarshal json theme to  string
func parseThemes(s string) []VsCodeTheme {
	var v VsCodeTheme
	var p PackageJson
	var result []VsCodeTheme

	filename := path.Join(s, "package.json")
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	err = json.Unmarshal(f, &p)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	for _, theme := range p.Contributes.Themes {
		tpath := path.Join(s, theme.Path)
		t, err := os.ReadFile(tpath)
		if err != nil {
			fmt.Println("error opening theme file", tpath)
			continue
		}
		ast, err := hujson.Parse(t)
		if err != nil {
			fmt.Println("error parsing theme file(hujson)", tpath)
			continue
		}
		ast.Standardize()
		t = ast.Pack()
		err = json.Unmarshal(t, &v)
		if err != nil {
			fmt.Println("error unmarshalling json", tpath)
			continue
		}
		err = json.Unmarshal(t, &v)
		if err != nil {
			fmt.Println("error parsing json file in ", theme.Path, err)
			continue
		} else {
			v.Name = theme.Label
			v.Type = theme.UITheme
			result = append(result, v)
		}
	}
	return result
}

//Output base16 theme as yaml string
// Todo: attach to struct
func unhash(s string) string {
	return strings.Replace(s, "#", "", 1)
}

func Base16Out(b Base16) string {
	var result string
	result += `scheme: "` + unhash(b.Scheme) + `" \n"`
	result += `type:   "` + unhash(b.Type) + `  " \n"`
	result += `author: "` + unhash(b.Author) + `" \n"`
	result += `base00: "` + unhash(b.Base00) + `" \n"`
	result += `base01: "` + unhash(b.Base01) + `" \n"`
	result += `base02: "` + unhash(b.Base02) + `" \n"`
	result += `base03: "` + unhash(b.Base03) + `" \n"`
	result += `base04: "` + unhash(b.Base04) + `" \n"`
	result += `base05: "` + unhash(b.Base05) + `" \n"`
	result += `base06: "` + unhash(b.Base06) + `" \n"`
	result += `base07: "` + unhash(b.Base07) + `" \n"`
	result += `base08: "` + unhash(b.Base08) + `" \n"`
	result += `base09: "` + unhash(b.Base09) + `" \n"`
	result += `base0A: "` + unhash(b.Base0A) + `" \n"`
	result += `base0B: "` + unhash(b.Base0B) + `" \n"`
	result += `base0C: "` + unhash(b.Base0C) + `" \n"`
	result += `base0D: "` + unhash(b.Base0D) + `" \n"`
	result += `base0E: "` + unhash(b.Base0E) + `" \n"`
	result += `base0F: "` + unhash(b.Base0F) + `" \n"`

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
		t := parseThemes(i)
		themes = append(themes, t...)
	}
	return themes
}

func ListThemes() {
	for i, t := range getAllThemes() {
		fmt.Println(i, t.Name)
	}
}

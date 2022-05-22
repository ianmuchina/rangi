package main

import (
	"encoding/json"
	"log"

	"github.com/tailscale/hujson"

	"fmt"
	"os"
	"path"
	"strings"
)

type Config struct {
	base16ThemePath   string
	base16ThemePrefix string
}

func foo() {
	// List all extensions
	// Get extensions with theme
	// Read theme data
	// Unmarshal theme data to struct
}

func main() {
	var themes = make(map[string]VsCodeTheme)
	var b Base16
	var err error

	for _, extension := range getExtensions() {
		theme := parseThemes(extension)
		for _, theme := range theme {
			themes[theme.Name] = theme
		}
	}

	for _, theme := range themes {
		b = theme.ToBase16()
		home, _ := os.UserHomeDir()
		themeDir := path.Join(home, ".local/share/flavours/base16/schemes/vscode/")
		themeFile := path.Join(themeDir, b.filename)

		err = os.MkdirAll(themeDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(themeFile, b.ToYaml(), 0755)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(themeFile)
		}
	}
}

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
			continue
			// os.Exit(1)
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
		themePath := path.Join(s, theme.Path)
		tmTheme, err := os.ReadFile(themePath)

		if err != nil {
			fmt.Println("error opening theme file", themePath)
			continue
		}

		ast, err := hujson.Parse(tmTheme)
		if err != nil {
			continue
		}
		ast.Standardize()
		tmTheme = ast.Pack()
		err = json.Unmarshal(tmTheme, &v)
		if err != nil {
			continue
		}
		err = json.Unmarshal(tmTheme, &v)
		if err != nil {
			continue
		}
		v.Name = theme.Label
		v.Type = theme.UITheme
		result = append(result, v)
	}
	return result
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

// Output base16 theme as yaml string
func unhash(s string) string {
	return strings.Replace(s, "#", "", 1)
}

//Convert a vscode theme to a base16 theme
func (v VsCodeTheme) ToBase16() Base16 {
	if v.Type == "vs" {
		v.Type = "light"
	} else if v.Type == "vs-dark" {
		v.Type = "dark"
	}

	fg := v.Colors.Foreground

	if len(v.Colors.EditorForeground) > len(v.Colors.Foreground) {
		fg = v.Colors.EditorForeground
	}
	var b Base16
	b.Scheme = v.Name
	b.Type = v.Type
	b.Author = "Anon"
	b.Base00 = v.Colors.EditorBackground
	b.Base01 = v.Colors.EditorBackground
	b.Base02 = v.Colors.EditorBackground
	b.Base03 = v.Colors.EditorBackground
	b.Base04 = fg
	b.Base05 = fg
	b.Base06 = fg
	b.Base07 = fg
	b.Base08 = v.Colors.TerminalAnsiRed
	b.Base09 = v.Colors.TerminalAnsiYellow
	b.Base0A = v.Colors.TerminalAnsiYellow
	b.Base0B = v.Colors.TerminalAnsiGreen
	b.Base0C = v.Colors.TerminalAnsiCyan
	b.Base0D = v.Colors.TerminalAnsiBlue
	b.Base0E = v.Colors.TerminalAnsiMagenta
	b.Base0F = v.Colors.TerminalAnsiRed

	// Generate filename
	// Convert to lowercase
	scheme := strings.ToLower(b.Scheme)
	// Replace spaces with underscores
	scheme = strings.ReplaceAll(scheme, " ", "_")
	scheme = strings.ReplaceAll(scheme, "'", "")
	// Generate filename
	b.filename = v.Type + "_" + scheme + "_vscode" + ".yaml"

	return b
}

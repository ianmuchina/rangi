package main

import (
	"fmt"
	"testing"
)

func TestLoadLinux(t *testing.T) {
	var arr []VsCodeTheme

	for _, i := range getExtensions() {
		themes := parseThemes(i)

		arr = append(arr, themes...)
	}

	for _, i := range arr {
		fmt.Println(i.Name)
	}
}

// TODO:
// 1. Create mock with multiple vscode themes
// 2. Create mock for applied styles
// 3. Add Better UX.
// 4. Test on windows
// 5. Add Fallback styles

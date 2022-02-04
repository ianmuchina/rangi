package main

import (
	"fmt"
	"testing"
)

func TestLoadLinux(t *testing.T) {
	var arr []VsCodeTheme

	for _, i := range getExtensions() {
		themes := LoadThemeJson(i)

		arr = append(arr, themes...)
	}

	for _, i := range arr {
		fmt.Println(i.Name)
	}
}

package main

import (
	"lazymod/apps"
	"os"
	"strings"
)

var statusCode string = "MAIN"

func main() {
	for {
		if statusCode == "MAIN" {
			model := apps.RunMainMenu()
			code := model.GetNextCode()
			updateStatusCode(code)
		}
		if statusCode == "ADD_CATEGORY" {
			model := apps.RunAddCategory()
			code := model.GetNextCode()
			updateStatusCode(code)
		}
		if statusCode == "ADD_STATS" {
			model := apps.RunAddStats()
			code := model.GetNextCode()
			updateStatusCode(code)
		}
		if statusCode == "EXIT" {
			os.Exit(0)
		}
	}
}

func updateStatusCode(code string) {
	if strings.TrimSpace(code) != "" {
		statusCode = code
	}
}

package main

import (
	"lazymod/apps"
	"os"
	"strings"
)

var statusCode string = "MAIN"

func main() {
	apps.RunAddStats()
	for {
		if statusCode == "MAIN" {
			model := apps.RunMainMenu()
			code := model.GetNextCode()
			handleStatusCode(code)
		}
		if statusCode == "ADD_CATEGORY" {
			model := apps.RunAddCategory()
			code := model.GetNextCode()
			handleStatusCode(code)
		}
		if statusCode == "EXIT" {
			os.Exit(0)
		}
	}
}

func handleStatusCode(code string) {
	if strings.TrimSpace(code) != "" {
		statusCode = code
	}
}

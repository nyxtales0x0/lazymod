package main

import (
	"lazymod/apps"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var statusCode string = "MAIN"
var lastModel tea.Model

func main() {
	for {
		if statusCode == "MAIN" {
			model, code := apps.RunMainMenu()
			lastModel = model
			handleStatusCode(code)
			println(statusCode)
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

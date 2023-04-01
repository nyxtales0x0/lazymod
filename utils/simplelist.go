package utils

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type listItem string

func (item listItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (delegate itemDelegate) Height() int { return 1 }

func (delegate itemDelegate) Spacing() int { return 0 }

func (delegate itemDelegate) Update(msg tea.Msg, model *list.Model) tea.Cmd { return nil }

func (delegate itemDelegate) Render(writer io.Writer, statsList list.Model, index int, stat list.Item) {
	uiString := ""
	if index == statsList.Index() {
		uiString = fmt.Sprintf("> %d. %s", index, stat)
	} else {
		uiString = fmt.Sprintf("  %d. %s", index, stat)
	}
	fmt.Fprint(writer, uiString)
}

func MakeSimpleList(items []string) list.Model {
	listItems := []list.Item{}
	for _, item := range items {
		listItems = append(listItems, listItem(item))
	}
	simpleList := list.New(listItems, itemDelegate{}, 20, 15)
	return simpleList
}

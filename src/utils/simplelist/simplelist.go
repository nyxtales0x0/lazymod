package simplelist

import "github.com/charmbracelet/bubbles/list"

type Item struct {
	title string
}

func (item Item) FilterValue() string {
	return item.title
}

func (item Item) Title() string {
	return item.title
}

func (item Item) Description() string {
	return ""
}

func New(items []string) list.Model {
	listItems := []list.Item{}
	for _, item := range items {
		listItems = append(listItems, Item{title: item})
	}
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	width := 50
	height := 20
	newList := list.New(listItems, delegate, width, height)
	// newList.DisableQuitKeybindings()
	return newList
}

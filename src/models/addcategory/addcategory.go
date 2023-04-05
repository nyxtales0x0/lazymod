package addcategory

import (
	"fmt"
	"lazymod/src/models/addstats"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Show     bool
	InputBox textinput.Model
	AddStats addstats.Model
}

func (model Model) Init() tea.Cmd {
	return tea.Batch(
		tea.ClearScreen,
		textinput.Blink,
	)
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if model.AddStats.Show {
		var _addStats tea.Model
		_addStats, cmd = model.AddStats.Update(msg)
		model.AddStats = _addStats.(addstats.Model)
		return model, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			model.AddStats.Show = true
		}
	}
	model.InputBox, cmd = model.InputBox.Update(msg)
	return model, cmd
}

func (model Model) View() string {
	if model.AddStats.Show {
		return model.AddStats.View()
	}
	return fmt.Sprintf(
		"Enter category name\n\n%s",
		model.InputBox.View(),
	)
}

func New() Model {
	inputBox := textinput.New()
	inputBox.Focus()
	return Model{
		Show:     false,
		InputBox: inputBox,
		AddStats: addstats.New(),
	}
}

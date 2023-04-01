package apps

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type editStatModel struct {
	minValue          string
	maxValue          string
	statMinValueInput textinput.Model
	statMaxValueInput textinput.Model
	activeInput       string
	nextCode          string
}

func (model editStatModel) Init() tea.Cmd {
	return textinput.Blink
}

func (model editStatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			model.activeInput = "MIN"
			model.statMaxValueInput.Blur()
			model.statMinValueInput.Focus()
		case "down":
			model.activeInput = "MAX"
			model.statMinValueInput.Blur()
			model.statMaxValueInput.Focus()
		}
	}
	var cmd tea.Cmd
	if model.activeInput == "MIN" {
		model.statMinValueInput, cmd = model.statMinValueInput.Update(msg)
		return model, cmd
	}
	if model.activeInput == "MAX" {
		model.statMaxValueInput, cmd = model.statMaxValueInput.Update(msg)
		return model, cmd
	}
	return model, nil
}

func (model editStatModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n",
		model.statMinValueInput.View(),
		model.statMaxValueInput.View(),
	)
}

func makeEditStatModel() editStatModel {
	statMinValueInput := textinput.New()
	statMinValueInput.Focus()
	statMaxValueInput := textinput.New()
	return editStatModel{
		minValue:          "0",
		maxValue:          "100",
		statMinValueInput: statMinValueInput,
		statMaxValueInput: statMaxValueInput,
		activeInput:       "MIN",
	}
}

func RunEditStat() editStatModel {
	app := tea.NewProgram(makeEditStatModel())
	model, _ := app.Run()
	return model.(editStatModel)
}

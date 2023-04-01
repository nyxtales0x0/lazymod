package apps

import (
	"lazymod/utils"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type mainMenuModel struct {
	choicesList list.Model
	nextCode    string
}

func makeMainMenu() mainMenuModel {
	choicesList := utils.MakeSimpleList([]string{
		"Add stats to mod menu",
		"Create a new category",
		"Edit existing categories",
	})
	return mainMenuModel{
		choicesList: choicesList,
	}
}

func (model mainMenuModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if model.choicesList.Index() == 1 {
				model.nextCode = "ADD_CATEGORY"
				return model, tea.Quit
			}
		}
	}
	var cmd tea.Cmd
	model.choicesList, cmd = model.choicesList.Update(msg)
	return model, cmd
}

func (model mainMenuModel) View() string {
	return model.choicesList.View()
}

func (model mainMenuModel) GetNextCode() string {
	return model.nextCode
}

func RunMainMenu() mainMenuModel {
	app := tea.NewProgram(makeMainMenu())
	model, _ := app.Run()
	return model.(mainMenuModel)
}

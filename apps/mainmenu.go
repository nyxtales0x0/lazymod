package apps

import tea "github.com/charmbracelet/bubbletea"

type mainMenuModel struct {
	choices   []string
	cursorPos int
	nextCode  string
}

func makeMainMenu() mainMenuModel {
	return mainMenuModel{
		choices: []string{
			"Create a new category",
			"Edit existing categories",
		},
	}
}

func (model mainMenuModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if model.cursorPos > 0 {
				model.cursorPos--
			}
		case "down":
			if model.cursorPos-1 < len(model.choices) {
				model.cursorPos++
			}
		case "enter":
			if model.cursorPos == 0 {
				model.nextCode = "ADD_CATEGORY"
				return model, tea.Quit
			}
		}
	}
	return model, nil
}

func (model mainMenuModel) View() string {
	uiString := "What would you like to do?\n\n"
	for index, choice := range model.choices {
		cursorString := "  "
		if index == model.cursorPos {
			cursorString = "> "
		}
		uiString += cursorString + choice + "\n"
	}
	return uiString
}

func (model mainMenuModel) GetNextCode() string {
	return model.nextCode
}

func RunMainMenu() mainMenuModel {
	app := tea.NewProgram(makeMainMenu())
	model, _ := app.Run()
	return model.(mainMenuModel)
}

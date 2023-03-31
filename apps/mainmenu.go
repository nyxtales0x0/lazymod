package apps

import tea "github.com/charmbracelet/bubbletea"

type MainMenuModel struct {
	choices   []string
	cursorPos int
}

func makeMainMenu() MainMenuModel {
	return MainMenuModel{
		choices: []string{
			"Create a new category",
			"Edit existing categories",
		},
		cursorPos: 0,
	}
}

func (model MainMenuModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return model, tea.Quit
		}
	}
	return model, nil
}

func (model MainMenuModel) View() string {
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

func RunMainMenu() (tea.Model, string) {
	app := tea.NewProgram(makeMainMenu())
	model, _ := app.Run()
	return model, "EXIT"
}

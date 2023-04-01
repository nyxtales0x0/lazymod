package apps

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type addCategoryModel struct {
	categoryName string
	inputBox     textinput.Model
	isBlinking   bool
	nextCode     string
}

func makeAddCategoryModel() addCategoryModel {
	inputBox := textinput.New()
	inputBox.Focus()
	return addCategoryModel{
		inputBox: inputBox,
	}
}

func (model addCategoryModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model addCategoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !model.isBlinking {
		model.isBlinking = true
		return model, textinput.Blink
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			model.categoryName = model.inputBox.Value()
			model.nextCode = "ADD_STATS"
			return model, tea.Quit
		}
	}
	var cmd tea.Cmd
	model.inputBox, cmd = model.inputBox.Update(msg)
	return model, cmd
}

func (model addCategoryModel) View() string {
	uiString := "What would you like to call the category?\n"
	uiString += "> " + model.inputBox.Value()
	return uiString
}

func (model addCategoryModel) GetCategoryName() string {
	return model.categoryName
}

func (model addCategoryModel) GetNextCode() string {
	return model.nextCode
}

func RunAddCategory() addCategoryModel {
	app := tea.NewProgram(makeAddCategoryModel())
	model, _ := app.Run()
	return model.(addCategoryModel)
}

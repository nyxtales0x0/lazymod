package mainmenu

import (
	"lazymod/src/models/addcategory"
	"lazymod/src/utils/simplelist"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Options     list.Model
	AddCategory addcategory.Model
}

func (model Model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	// update add category screen
	if model.AddCategory.Show {

		// Update() returns tea.Model so we create
		// a placeholder to grab the return model
		// and cast it to addcategory.Model type
		var _addCategory tea.Model

		_addCategory, cmd = model.AddCategory.Update(msg)
		model.AddCategory = _addCategory.(addcategory.Model)
		return model, cmd

	}

	// handling key press event
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":

			// index 0 means "create category"
			// option is selected
			if model.Options.Index() == 0 {
				model.AddCategory.Show = true
			}
		}
	}

	// update main menu model otherwise
	model.Options, cmd = model.Options.Update(msg)
	return model, cmd
}

func (model Model) View() string {
	uiString := ""

	// views add category screen
	// if it's set to show
	if model.AddCategory.Show {
		uiString += model.AddCategory.View()
		return uiString
	}

	// view main menu model otherwise
	uiString += model.Options.View()
	return uiString
}

func newMainMenuModel() Model {

	// creating a simple list
	// with below mentioned entries
	options := simplelist.New([]string{
		"Create new category",
		"Edit existing categories",
	})

	// setting main menu's title
	options.Title = "What would you like to do?"

	return Model{
		Options:     options,
		AddCategory: addcategory.New(),
	}
}

func RunMainMenu() {
	mainMenu := newMainMenuModel()
	app := tea.NewProgram(mainMenu)
	app.Run()
}

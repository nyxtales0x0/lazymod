package addstats

import (
	"lazymod/src/structs/scene"
	"lazymod/src/utils/simplelist"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Show      bool
	StatsList list.Model
}

func (model Model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	model.StatsList, cmd = model.StatsList.Update(msg)
	return model, cmd
}

func (model Model) View() string {
	uiString := ""
	uiString += model.StatsList.View() + "ho"
	return uiString
}

func New() Model {
	startupScene := scene.NewSceneFromFile("startup.txt.json")
	statLines := []string{}
	for _, line := range startupScene.Lines {
		if strings.Contains(line, "*create") {
			statLines = append(statLines, line)
		}
	}
	return Model{
		Show:      false,
		StatsList: simplelist.New(statLines),
	}
}

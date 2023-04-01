package apps

import (
	"lazymod/utils"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type addStatsModel struct {
	statsList list.Model
	nextCode  string
}

func makeAddStatsModel() addStatsModel {
	startupScene := utils.ParseSceneFile("startup.txt.json")
	stats := []string{}
	for _, line := range startupScene.Lines {
		if strings.Contains(line, "*create") {
			stats = append(stats, line)
		}
	}
	statsList := utils.MakeSimpleList(stats)
	return addStatsModel{
		statsList: statsList,
	}
}

func (model addStatsModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model addStatsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			model.nextCode = "EXIT"
			return model, tea.Quit
		}
	}
	var cmd tea.Cmd
	model.statsList, cmd = model.statsList.Update(msg)
	return model, cmd
}

func (model addStatsModel) View() string {
	return model.statsList.View()
}

func (model addStatsModel) GetNextCode() string {
	return model.nextCode
}

func RunAddStats() addStatsModel {
	app := tea.NewProgram(makeAddStatsModel())
	model, _ := app.Run()
	return model.(addStatsModel)
}

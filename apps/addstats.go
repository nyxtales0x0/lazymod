package apps

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type statItem string

func (stat statItem) FilterValue() string { return "" }

type statsDelegate struct{}

func (delegate statsDelegate) Height() int { return 1 }

func (delegate statsDelegate) Spacing() int { return 0 }

func (delegate statsDelegate) Update(msg tea.Msg, model *list.Model) tea.Cmd { return nil }

func (delegate statsDelegate) Render(writer io.Writer, statsList list.Model, index int, stat list.Item) {
	uiString := ""
	if index == statsList.Index() {
		uiString = fmt.Sprintf("> %d. %s", index, stat)
	} else {
		uiString = fmt.Sprintf("  %d. %s", index, stat)
	}
	fmt.Fprint(writer, uiString)
}

type addStatsModel struct {
	statsList list.Model
}

func makeAddStatsModel() addStatsModel {
	statItems := []list.Item{
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
		statItem("Hello"),
	}
	statsList := list.New(statItems, statsDelegate{}, 20, 15)
	return addStatsModel{
		statsList,
	}
}

func (model addStatsModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model addStatsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	model.statsList, cmd = model.statsList.Update(msg)
	return model, cmd
}

func (model addStatsModel) View() string {
	return model.statsList.View()
}

func RunAddStats() addStatsModel {
	app := tea.NewProgram(makeAddStatsModel())
	model, _ := app.Run()
	return model.(addStatsModel)
}

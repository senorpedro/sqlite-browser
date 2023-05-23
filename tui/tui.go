package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"senorpedro.com/sqlite-browser/db"
)

type TuiModel struct {
	SqliteReader *db.SqliteReader
	tablesList   TablesListModel
	tableContent TableContentModel
	help         help.Model
	listVisible  bool
}

func (model TuiModel) getActiveModel() string {
	if model.tableContent.Active() {
		return "content"
	}

	return "list"
}

func (model TuiModel) Init() tea.Cmd {
	return nil
}

func (model TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		height := msg.Height - HelpHeight
		model.tablesList.SetHeight(height)
		model.tableContent.SetHeight(height)

		width := msg.Width
		if model.listVisible {
			width -= TablesListWidth
		}

		model.tableContent.SetWidth(width)

		return model, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keys.T):
			model.listVisible = !model.listVisible
			if !model.listVisible {
				model.tableContent.SetActive(true)
				model.tablesList.SetActive(false)
			}

		case key.Matches(msg, Keys.Quit):
			return model, tea.Quit
		case key.Matches(msg, Keys.Tab):
			if model.listVisible {
				// switch focus to other pane
				if model.tableContent.Active() {
					model.tableContent.SetActive(false)
					model.tablesList.SetActive(true)
				} else {
					model.tableContent.SetActive(true)
					model.tablesList.SetActive(false)
				}
			}
		case key.Matches(msg, Keys.Down):
			active := model.getActiveModel()
			if active == "content" {
				model.tableContent, cmd = model.tableContent.Update(msg)
			} else {
				model.tablesList, cmd = model.tablesList.Update(msg)
				model.LoadSelectedTable()
			}
		case key.Matches(msg, Keys.Up):
			active := model.getActiveModel()
			if active == "content" {
				model.tableContent, cmd = model.tableContent.Update(msg)
			} else {
				model.tablesList, cmd = model.tablesList.Update(msg)
				model.LoadSelectedTable()
			}

		case key.Matches(msg, Keys.Help):
			model.help.ShowAll = !model.help.ShowAll
			/*
				case "enter":
					// Load file contents into the contentList
						selected := m.files[m.fileList.SelectedIndex()]
						if !selected.IsDir() {
							content, err := ioutil.ReadFile(selected.Name())
							if err != nil {
								m.contentList = list.NewModel([]string{"Error reading file"}, false)
							} else {
								m.contentList = list.NewModel(strings.Split(string(content), "\n"), false)
							}
						}
			*/
		}
	}

	return model, cmd
}

func (model TuiModel) View() string {

	tableContentView := model.tableContent.View()
	var view string

	if model.listVisible {

		tablesListView := model.tablesList.View()

		// Style the two panes
		leftPane := lipgloss.NewStyle().Render(tablesListView)
		// leftPane := lipgloss.NewStyle().Width(30).Height(100).Render(tablesListView)
		rightPane := lipgloss.NewStyle().Render(tableContentView)

		// Combine the panes horizontally
		view = lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)

	} else {
		view = tableContentView
	}

	view += Styles.Help(model.help.View(Keys))

	return view
}

func (model *TuiModel) LoadSelectedTable() {
	selectedTable := model.tablesList.SelectedTable()
	model.tableContent.Load(selectedTable)
}

func StartUI(s *db.SqliteReader) {

	tableNames := s.TableNames()

	// init ui
	tablesList := CreateTablesListModel(tableNames)
	tablesList.SetActive(true)

	tableContent := CreateTableContentModel(s)
	tableContent.SetActive(false)

	tui := TuiModel{
		SqliteReader: s,
		listVisible:  true,
		help:         help.New(),
		tablesList:   tablesList,
		tableContent: tableContent,
	}

	tui.LoadSelectedTable()

	if _, err := tea.NewProgram(tui, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}

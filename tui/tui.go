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

func (t Tui) getActiveView() string {
	if t.tableContentView.Active() {
		return "content"
	} else {
		return "list"
	}
}

type Tui struct {
	SqliteReader     *db.SqliteReader
	tablesListView   TablesListView
	tableContentView TableContentView
	help             help.Model
}

func (t Tui) Init() tea.Cmd {
	return nil
}

func (t Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		height := msg.Height - HelpHeight
		t.tablesListView.SetHeight(height)
		t.tableContentView.SetHeight(height)

		width := msg.Width - TablesListWidth
		t.tableContentView.SetWidth(width)

		return t, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keys.Quit):
			return t, tea.Quit
		case key.Matches(msg, Keys.Tab):
			// switch focus to other pane
			if t.tableContentView.Active() {
				t.tableContentView.SetActive(false)
				t.tablesListView.SetActive(true)
			} else {
				t.tableContentView.SetActive(true)
				t.tablesListView.SetActive(false)
			}
		case key.Matches(msg, Keys.Down):
			active := t.getActiveView()
			if active == "content" {
				t.tableContentView, cmd = t.tableContentView.Update(msg)
			} else {
				t.tablesListView, cmd = t.tablesListView.Update(msg)
				t.LoadSelectedTable()
			}
		case key.Matches(msg, Keys.Up):
			active := t.getActiveView()
			if active == "content" {
				t.tableContentView, cmd = t.tableContentView.Update(msg)
			} else {
				t.tablesListView, cmd = t.tablesListView.Update(msg)
				t.LoadSelectedTable()
			}

		case key.Matches(msg, Keys.Help):
			t.help.ShowAll = !t.help.ShowAll
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

	return t, cmd
}

func (t Tui) View() string {
	tablesListView := t.tablesListView.View()
	tableContentView := t.tableContentView.View()

	// Style the two panes
	leftPane := lipgloss.NewStyle().Render(tablesListView)
	// leftPane := lipgloss.NewStyle().Width(30).Height(100).Render(tablesListView)
	rightPane := lipgloss.NewStyle().Render(tableContentView)

	// Combine the panes horizontally
	view := lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)

	view += Styles.Help(t.help.View(Keys))

	return view
}

func (t *Tui) LoadSelectedTable() {
	selectedTable := t.tablesListView.SelectedTable()
	t.tableContentView.Load(selectedTable)
}

func StartUI(s *db.SqliteReader) {

	tui := Tui{SqliteReader: s}

	tableNames := tui.SqliteReader.TableNames()

	// init ui
	tablesListView := CreateTablesListView(tableNames)
	tablesListView.SetActive(true)
	tui.tablesListView = tablesListView

	tableContentView := CreateTablesContentView(s)
	tableContentView.SetActive(false)
	tui.tableContentView = tableContentView

	tui.LoadSelectedTable()

	tui.help = help.New()

	if _, err := tea.NewProgram(tui, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}

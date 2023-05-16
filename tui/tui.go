package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"senorpedro.com/sqlite-browser/db"
	"senorpedro.com/sqlite-browser/debug"
)

func (t Tui) getActiveView() string {
	if t.tableContentView.Active {
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
		debug.Log(msg.Height)
		/*
			t.tablesListView.SetHeight(msg.Height)
			t.tableContentView.SetHeight(msg.Height)
		*/
		t.tablesListView.Update(msg)
		t.tableContentView.Update(msg)

		return t, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keys.Quit):
			return t, tea.Quit
		case key.Matches(msg, Keys.Tab):
			// switch focus to other pane
			if t.tableContentView.Active {
				t.tableContentView.Active = false
				t.tablesListView.Active = true
			} else {
				t.tableContentView.Active = true
				t.tablesListView.Active = false
			}
		case key.Matches(msg, Keys.Down):
			active := t.getActiveView()
			if active == "content" {
				t.tableContentView, cmd = t.tableContentView.Update(msg)
			} else {
				t.tablesListView, cmd = t.tablesListView.Update(msg)
			}
		case key.Matches(msg, Keys.Up):
			active := t.getActiveView()
			if active == "content" {
				t.tableContentView, cmd = t.tableContentView.Update(msg)
			} else {
				t.tablesListView, cmd = t.tablesListView.Update(msg)
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

func StartUI(s *db.SqliteReader) {

	tui := Tui{SqliteReader: s}

	tableNames := tui.SqliteReader.TableNames()

	// init ui
	tablesListView := CreateTablesListView(tableNames)
	tablesListView.Active = true
	tui.tablesListView = tablesListView

	// TODO replace with real data
	header := []string{"Name", "Age", "Gender"}
	data := [][]string{
		{"Alice", "25", "Female"},
		{"Bob", "30", "Male"},
		{"Charlie", "40", "Male"},
		{"Diana", "35", "Female"},
	}

	tableContentView := CreateTablesContentView(header, data)
	tableContentView.Active = false
	tui.tableContentView = tableContentView

	tui.help = help.New()

	if _, err := tea.NewProgram(tui, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}

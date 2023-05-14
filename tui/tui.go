package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"senorpedro.com/sqlite-browser/db"
)

type Tui struct {
	SqliteReader     *db.SqliteReader
	tablesListView   TablesListView
	tableContentView TableContentView
}

func (t Tui) Init() tea.Cmd {
	return nil
}

func (t Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return t, tea.Quit
		case "enter":
			// Load file contents into the contentList
			/*
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
		default:
			t.tablesListView, cmd = t.tablesListView.Update(msg)
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

	return view
}

func StartUI(s *db.SqliteReader) {

	tui := Tui{SqliteReader: s}

	tableNames := tui.SqliteReader.TableNames()

	// init ui
	tlv := CreateTablesListView(tableNames)
	tui.tablesListView = tlv

	header := []string{"Name", "Age", "Gender"}
	data := [][]string{
		{"Alice", "25", "Female"},
		{"Bob", "30", "Male"},
		{"Charlie", "40", "Male"},
		{"Diana", "35", "Female"},
	}

	tcv := CreateTablesContentView(header, data)
	tui.tableContentView = tcv

	if _, err := tea.NewProgram(tui, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	/*
		if _, err := tea.NewProgram(tui.tableBrowser).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	*/

}
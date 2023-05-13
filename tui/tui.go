package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"senorpedro.com/sqlite-browser/db"
)

type Tui struct {
	SqliteReader *db.SqliteReader
	tableBrowser TableBrowser
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
			t.tableBrowser, cmd = t.tableBrowser.Update(msg)
		}
	}

	return t, cmd
}

func (t Tui) View() string {
	tableBrowserView := t.tableBrowser.View()
	//contentListView := m.contentList.View()

	// Style the two panes
	// leftPane := lipgloss.NewStyle().Width(30).Height(100).Render(tableBrowserView)
	// rightPane := lipgloss.NewStyle().Width(70).Height(100).Render(contentListView)

	// Combine the panes horizontally
	// view := lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)

	return tableBrowserView
}

func StartUI(s *db.SqliteReader) {

	tui := Tui{SqliteReader: s}

	tables := tui.SqliteReader.TableNames()

	// init ui
	tb := CreateTableBrowser(tables)

	tui.tableBrowser = tb

	if _, err := tea.NewProgram(tui).Run(); err != nil {
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

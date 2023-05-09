package main

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type column struct {
	name     string
	datatype string
}

type table struct {
	name    string
	columns []column
}

type database struct {
	name   string
	file   string
	tables []table
}

var db = database{
	name: "test db",
	file: "test.sqlite",
	tables: []table{
		table{
			name: "table1",
			columns: []column{
				{name: "id", datatype: "integer"},
				{name: "name", datatype: "string"},
			},
		},

		table{
			name: "table2",
			columns: []column{
				{name: "id", datatype: "integer"},
				{name: "email", datatype: "string"},
			},
		},
	},
}

type model struct {
	list                list.Model
	highlightedTableIdx int
	tableExpanded       bool
	choice              string
	quitting            bool
}

func newModel() model {
	return model{
		highlightedTableIdx: 1,
		tableExpanded:       false,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// XXX load database if it was provided as cli arg

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// react to inputs etc

	// react to window size changes

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd

	return m, nil
}

func (m model) View() string {
	// TODO add help (for current context)

	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}
	return "\n" + m.list.View()
}

func main() {
	/**
	TODO
		- WithAltScreen => fullscreen
		- compose


	*/

	// tableItems := make([]list.Item{}, len(db.tables))
	tableItems := make([]list.Item, len(db.tables))

	for i, table := range db.tables {
		tableItems[i] = item(table.name)
	}

	const defaultWidth = 20

	l := list.New(tableItems, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "What do you want for dinner?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	//
	/*
			if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
				fmt.Println("Error while running program:", err)
				os.Exit(1)
		}
	*/
}

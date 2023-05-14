package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type TableContentView struct {
	table table.Model
}

// Define a message that represents a request to render the table
type renderTableMsg struct{}

func (v TableContentView) Init() tea.Cmd {
	return nil
}

func (v TableContentView) Update(msg tea.Msg) (TableContentView, tea.Cmd) {
	return v, nil
}

func (v TableContentView) View() string {
	return v.table.View()
}

func CreateTablesContentView(columnNames []string, data [][]string) TableContentView {

	columns := make([]table.Column, len(columnNames))

	for i, column := range columnNames {
		columns[i] = table.Column{Title: column, Width: 20}
	}

	rows := make([]table.Row, len(data))

	for i, row := range data {
		rows[i] = row
	}

	/*
		initialModel := table{
			header: []string{"Name", "Age", "Gender"},
			data: [][]string{
				{"Alice", "25", "Female"},
				{"Bob", "30", "Male"},
				{"Charlie", "40", "Male"},
				{"Diana", "35", "Female"},
			},
		}
	*/

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	v := TableContentView{table: t}
	return v

}

//func (t table) Init() tea.Cmd {
//	// No command needed for initialization
//	return nil
//}
//
//func (t table) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	switch msg.(type) {
//	case tea.KeyMsg:
//		// Exit the program when the user presses any key
//		return t, tea.Quit
//	case renderTableMsg:
//		// No state update necessary, just return the same model with no command
//		return t, nil
//	default:
//		// Ignore other messages
//		return t, nil
//	}
//}
//
//func (t table) View() string {
//	// Calculate the width of each column
//	colWidths := make([]int, len(t.header))
//	for i, col := range t.header {
//		colWidths[i] = len(col)
//	}
//	for _, row := range t.data {
//		for i, val := range row {
//			if len(val) > colWidths[i] {
//				colWidths[i] = len(val)
//			}
//		}
//	}
//
//	// Render the table header
//	header := ""
//	for i, col := range t.header {
//		header += fmt.Sprintf("| %s%s ", col, style.PadRight("", colWidths[i]-len(col)))
//	}
//	header += "|\n" + style.Border
//
//	// Render the table body
//	body := ""
//	for _, row := range t.data {
//		for i, val := range row {
//			body += fmt.Sprintf("| %s%s ", val, style.PadRight("", colWidths[i]-len(val)))
//		}
//		body += "|\n"
//	}
//
//	return style.Bold(header + body + style.Border)
//}
//

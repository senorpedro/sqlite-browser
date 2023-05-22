package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"senorpedro.com/sqlite-browser/db"
)

type TableContentView struct {
	sqliteReader *db.SqliteReader
	active       bool
	table        table.Model
}

// Define a message that represents a request to render the table
type renderTableMsg struct{}

func (v TableContentView) Init() tea.Cmd {
	return nil
}

func (v TableContentView) Update(msg tea.Msg) (TableContentView, tea.Cmd) {
	var cmd tea.Cmd
	v.table, cmd = v.table.Update(msg)
	return v, cmd
}

func (v *TableContentView) SetActive(a bool) {
	if a {
		v.table.Focus()
	} else {
		v.table.Blur()
	}
	v.active = a
}

func (v TableContentView) Active() bool {
	return v.active
}

func (v *TableContentView) SetHeight(h int) {
	v.table.SetHeight(h)
}

func (v *TableContentView) SetWidth(w int) {
	v.table.SetWidth(w)
}

func (v *TableContentView) Load(tableName string) {
	v.NewTable()
	columnInfo := v.sqliteReader.TableInfo(tableName)

	columns := make([]table.Column, len(columnInfo))
	columnIdxMap := make(map[string]int)

	for i, column := range columnInfo {
		columns[i] = table.Column{Title: column.Name, Width: 20}
		columnIdxMap[column.Name] = i
	}

	tableData := v.sqliteReader.TableContent(tableName)

	rows := make([]table.Row, len(tableData))

	for idx, tableRow := range tableData {
		rows[idx] = make(table.Row, len(tableRow))

		for columnName, value := range tableRow {
			columnIdx := columnIdxMap[columnName]

			rows[idx][columnIdx] = fmt.Sprintf("%v", value)
		}
	}

	v.table.SetColumns(columns)
	v.table.SetRows(rows)
}

func (v TableContentView) View() string {
	var box RenderFunc
	if v.active {
		box = Styles.BoxActive
	} else {
		box = Styles.BoxInactive
	}

	return box(v.table.View())
}

func (v *TableContentView) NewTable() {

	v.table = table.New(
		table.WithFocused(true),
		table.WithHeight(7),
	)
}

func CreateTablesContentView(s *db.SqliteReader) TableContentView {

	v := TableContentView{sqliteReader: s}
	v.NewTable()

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

package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"senorpedro.com/sqlite-browser/db"
)

type TableContentModel struct {
	sqliteReader *db.SqliteReader
	active       bool
	table        table.Model
}

// Define a message that represents a request to render the table
type renderTableMsg struct{}

func (model TableContentModel) Init() tea.Cmd {
	return nil
}

func (model TableContentModel) Update(msg tea.Msg) (TableContentModel, tea.Cmd) {
	var cmd tea.Cmd
	model.table, cmd = model.table.Update(msg)
	return model, cmd
}

func (model *TableContentModel) SetActive(a bool) {
	if a {
		model.table.Focus()
	} else {
		model.table.Blur()
	}
	model.active = a
}

func (model TableContentModel) Active() bool {
	return model.active
}

func (model *TableContentModel) SetHeight(h int) {
	model.table.SetHeight(h)
}

func (model *TableContentModel) SetWidth(w int) {
	model.table.SetWidth(w)
}

func (model *TableContentModel) Load(tableName string) {
	model.NewTable()
	columnInfo := model.sqliteReader.TableInfo(tableName)

	columns := make([]table.Column, len(columnInfo))
	columnIdxMap := make(map[string]int)

	for i, column := range columnInfo {
		columns[i] = table.Column{Title: column.Name, Width: 20}
		columnIdxMap[column.Name] = i
	}

	tableData := model.sqliteReader.TableContent(tableName)

	rows := make([]table.Row, len(tableData))

	for idx, tableRow := range tableData {
		rows[idx] = make(table.Row, len(tableRow))

		for columnName, value := range tableRow {
			columnIdx := columnIdxMap[columnName]

			rows[idx][columnIdx] = fmt.Sprintf("%v", value)
		}
	}

	model.table.SetColumns(columns)
	model.table.SetRows(rows)
}

func (model TableContentModel) View() string {
	var box RenderFunc
	if model.active {
		box = Styles.BoxActive
	} else {
		box = Styles.BoxInactive
	}

	return box(model.table.View())
}

func (model *TableContentModel) NewTable() {

	model.table = table.New(
		table.WithFocused(true),
		table.WithHeight(7),
	)
}

func CreateTableContentModel(s *db.SqliteReader) TableContentModel {

	model := TableContentModel{sqliteReader: s}
	model.NewTable()

	return model
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

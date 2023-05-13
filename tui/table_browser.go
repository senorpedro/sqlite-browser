package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// boilerplate stuff...
type listItem string

func (i listItem) FilterValue() string { return "" }

type TableBrowser struct { // bubbletea model
	list list.Model
}

func CreateTableBrowser(tableNames []string) TableBrowser {

	tableItems := make([]list.Item, len(tableNames))

	for i, table := range tableNames {
		tableItems[i] = list.Item(listItem(table))
	}

	list := list.NewModel(tableItems, list.NewDefaultDelegate(), 14, 20)
	list.Title = "Tables"
	/*
		list.SetShowStatusBar(false)
		list.SetFilteringEnabled(false)
		list.Styles.Title = titleStyle
		list.Styles.PaginationStyle = paginationStyle
		list.Styles.HelpStyle = helpStyle
	*/

	tb := TableBrowser{
		list: list,
	}
	return tb
}

func (tb TableBrowser) Init() tea.Cmd {
	return nil
}

func (tb TableBrowser) Update(msg tea.Msg) (TableBrowser, tea.Cmd) {
	return tb, nil
}

func (tb TableBrowser) View() string {
	return "\n" + tb.list.View()
}

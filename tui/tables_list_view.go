package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// boilerplate stuff...

// TODO constants
var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type listItem string

func (i listItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, currentListItem list.Item) {
	i, ok := currentListItem.(listItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type TablesListView struct { // bubbletea model
	list   list.Model
	active bool
	height int
}

func (tlv *TablesListView) SetActive(a bool) {
	tlv.active = a
}

func (tlv TablesListView) Active() bool {
	return tlv.active
}

func (tlv *TablesListView) SetHeight(h int) {
	tlv.list.SetHeight(h + 1)
}

func CreateTablesListView(tableNames []string) TablesListView {

	tableItems := make([]list.Item, len(tableNames))

	for i, table := range tableNames {
		tableItems[i] = list.Item(listItem(table))
	}

	list := list.New(tableItems, itemDelegate{}, TablesListWidth, 20)
	list.Title = "Available Tables"
	list.SetShowHelp(false)
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	/*
		list.Styles.Title = titleStyle
		list.Styles.PaginationStyle = paginationStyle
		list.Styles.HelpStyle = helpStyle
	*/

	tlv := TablesListView{
		list: list,
	}
	return tlv
}

func (tlv TablesListView) SelectedTable() string {
	selectedTable, ok := tlv.list.SelectedItem().(listItem)
	if ok {
		return string(selectedTable)
	}
	// TODO implement proper handling
	return ""
}

func (tlv TablesListView) Init() tea.Cmd {
	return nil
}

func (tlv *TablesListView) Update(msg tea.Msg) (TablesListView, tea.Cmd) {
	var cmd tea.Cmd
	tlv.list, cmd = tlv.list.Update(msg)

	return *tlv, cmd
}

func (tlv TablesListView) View() string {
	var box RenderFunc
	if tlv.active {
		box = Styles.BoxActive
	} else {
		box = Styles.BoxInactive
	}
	return box(tlv.list.View())
}

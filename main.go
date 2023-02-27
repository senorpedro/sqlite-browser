package main

import (
	tea "github.com/charmbracelet/bubbletea"
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

type model struct {
	db database
}

func newModel() model {
	return model{
		db: database{
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
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// XXX load database if it was provided as cli arg

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// react to inputs etc

	return m, nil
}

func (m model) View() string {
	// TODO add help (for current context)

	return ""

}

func main() {
	/**
	TODO
		- WithAltScreen => fullscreen
		- compose


	*/
	//
	/*
		if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
			fmt.Println("Error while running program:", err)
			os.Exit(1)
		}
	*/
}

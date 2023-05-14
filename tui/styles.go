package tui

import "github.com/charmbracelet/lipgloss"

// ATM only for light background!
// TODO define colors for dark background as well!
type Colors struct {
	active   string
	inactive string
}

var ColorsBgLight = Colors{
	active:   "#1c1c1c",
	inactive: "#cccccc",
}

type RenderFunc func(strs ...string) string

type styles struct {
	BoxActive   RenderFunc
	BoxInactive RenderFunc
	Help        RenderFunc
}

var Styles = styles{
	BoxActive: lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(ColorsBgLight.active)).Render,
	BoxInactive: lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(ColorsBgLight.inactive)).Render,
	Help: lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render,
}

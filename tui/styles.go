package tui

import "github.com/charmbracelet/lipgloss"

var TablesListWidth = 30

// ATM only for light background!
// TODO define colors for dark background as well!
type colors struct {
	border_active   lipgloss.AdaptiveColor
	border_inactive lipgloss.AdaptiveColor
	help            lipgloss.AdaptiveColor
}

var Colors = colors{
	border_active:   lipgloss.AdaptiveColor{Light: "#1c1c1c", Dark: "#ffffff"},
	border_inactive: lipgloss.AdaptiveColor{Light: "#cccccc", Dark: "#fefefe"},
	help:            lipgloss.AdaptiveColor{Light: "#B2B2B2", Dark: "#4A4A4A"},
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
		BorderForeground(Colors.border_active).Render,
	BoxInactive: lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(Colors.border_inactive).Render,
	Help: lipgloss.NewStyle().
		Padding(1, 0, 0, 2).Render,
}

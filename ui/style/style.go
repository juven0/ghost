package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	AccentColor       = lipgloss.Color("#00ACC1")
	BackgroundColor   = lipgloss.Color("#6b6b6b")
	ActiveTextColor   = lipgloss.Color("#EEE")
	NormalTextColor   = lipgloss.Color("#CCC")
	InactiveTextColor = lipgloss.Color("#888")
)

var (
	ColumnStyle  = lipgloss.NewStyle().Padding(1, 2)
	FocusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(AccentColor))
)

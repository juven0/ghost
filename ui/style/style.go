package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	AccentColor       = lipgloss.Color("#FFE0B2")
	BackgroundColor   = lipgloss.Color("#6b6b6b")
	ActiveTextColor   = lipgloss.Color("#EEE")
	NormalTextColor   = lipgloss.Color("#CCC")
	InactiveTextColor = lipgloss.Color("#888")
)

var (
	IconPlay     = "▶"
	IconStop     = "■"
	IconLiked    = "💛"
	IconNotLiked = "🤍"
)

var (
	AccentTextStyle = lipgloss.NewStyle().Foreground(AccentColor)
)

var (
	ColumnStyle  = lipgloss.NewStyle().Padding(1, 2)
	FocusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(AccentColor))
	TrackListStyle = lipgloss.NewStyle().
			Padding(1, 2).
			MarginTop(-2)
	TrackListActiveStyle = lipgloss.NewStyle().
				Padding(0, 1).
				MarginTop(-2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(AccentColor)
	TrackTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#dcdcdc")).
			Bold(true)
)

var (
	ButtonStyle = lipgloss.NewStyle().
			Foreground(NormalTextColor).
			Background(InactiveTextColor).
			Padding(0, 3).
			MarginTop(1)
	ActiveButtonStyle = ButtonStyle.Copy().
				Foreground(InactiveTextColor).
				Background(AccentColor)
)

var (
	TrackBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#444")).
			Padding(1, 2)
	TrackVersionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#999999"))
	TrackArtistStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#dcdcdc"))

	TrackProgressStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				PaddingBottom(1)

	TrackAddInfoStyle = lipgloss.NewStyle().
				Align(lipgloss.Right).
				Width(26)
)

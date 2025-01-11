package asset

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	NormalFg  = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
	LightBlue = lipgloss.AdaptiveColor{Light: "#35FCDC", Dark: "#00d0ff"}
	DarkBlue  = lipgloss.AdaptiveColor{Light: "#04FAD3", Dark: "#4BCBE7"}
	Black     = lipgloss.AdaptiveColor{Light: "#000", Dark: "#000"}

	Indigo  = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	Cream   = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
	Fuchsia = lipgloss.Color("#F780E2")
	Green   = lipgloss.AdaptiveColor{Light: "#0dff00", Dark: "#0dff00"}
	Red     = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}

	// Emojis
	EmojiSparkles = "\U00002728" // ✨
	EmojiError    = "\U0000274C" // ❌
	EmojiTick     = "\U00002705" // ✅
	EmojiThumbsUp = "\U0001F44D" // 👍
	EmojiConfused = "\U0001F615" // 😕

	VgoLogo = lipgloss.NewStyle().
		Foreground(LightBlue).
		PaddingLeft(1).
		Bold(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(DarkBlue).
		Render(`
\  / _  _
 \/ (_](_)
    ._|
`)

	Text = lipgloss.NewStyle().
		PaddingLeft(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(DarkBlue).
		Foreground(LightBlue)
)

func SetTheme() *huh.Theme {
	t := huh.ThemeBase()

	t.Focused.Base = t.Focused.Base.BorderForeground(LightBlue)
	t.Focused.Title = t.Focused.Title.Foreground(LightBlue).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(LightBlue)
	t.Focused.Directory = t.Focused.Directory.Foreground(Indigo)
	t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(Fuchsia)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(Fuchsia)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(Fuchsia)
	t.Focused.Option = t.Focused.Option.Foreground(NormalFg)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(Fuchsia)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(LightBlue)
	t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(NormalFg)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(Cream).Background(Fuchsia)
	t.Focused.Next = t.Focused.FocusedButton
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(NormalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	t.Focused.TextInput.Text = t.Focused.TextInput.Text.Foreground(DarkBlue)
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(Green)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(Fuchsia)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}

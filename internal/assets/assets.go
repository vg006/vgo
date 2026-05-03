package asset

import (
	"charm.land/huh/v2"
	"charm.land/huh/v2/spinner"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

var (
	// Colors
	NormalFg  = compat.AdaptiveColor{Light: lipgloss.Color("235"), Dark: lipgloss.Color("252")}
	LightBlue = compat.AdaptiveColor{Light: lipgloss.Color("#35FCDC"), Dark: lipgloss.Color("#00d0ff")}
	DarkBlue  = compat.AdaptiveColor{Light: lipgloss.Color("#04FAD3"), Dark: lipgloss.Color("#4BCBE7")}
	Black     = compat.AdaptiveColor{Light: lipgloss.Color("#000"), Dark: lipgloss.Color("#000")}

	Indigo  = compat.AdaptiveColor{Light: lipgloss.Color("#5A56E0"), Dark: lipgloss.Color("#7571F9")}
	Cream   = compat.AdaptiveColor{Light: lipgloss.Color("#FFFDF5"), Dark: lipgloss.Color("#FFFDF5")}
	Fuchsia = compat.AdaptiveColor{Light: lipgloss.Color("#F780E2"), Dark: lipgloss.Color("#F780E2")}
	Green   = compat.AdaptiveColor{Light: lipgloss.Color("#0dff00"), Dark: lipgloss.Color("#0dff00")}
	Red     = compat.AdaptiveColor{Light: lipgloss.Color("#FF4672"), Dark: lipgloss.Color("#ED567A")}

	MutedFg         = compat.AdaptiveColor{Light: lipgloss.Color(""), Dark: lipgloss.Color("243")}
	SelectGreen     = compat.AdaptiveColor{Light: lipgloss.Color("#02CF92"), Dark: lipgloss.Color("#02A877")}
	BlurredButtonBg = compat.AdaptiveColor{Light: lipgloss.Color("252"), Dark: lipgloss.Color("237")}
	PlaceholderFg   = compat.AdaptiveColor{Light: lipgloss.Color("248"), Dark: lipgloss.Color("238")}

	// Emojis
	EmojiSparkles = "\U00002728" // ✨
	EmojiError    = "\U0000274C" // ❌
	EmojiTick     = "\U00002714" // ✔
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

func SetTheme(isDark bool) *huh.Styles {
	t := huh.ThemeBase(isDark)

	t.Focused.Base = t.Focused.Base.BorderForeground(LightBlue)
	t.Focused.Title = t.Focused.Title.Foreground(LightBlue).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(LightBlue)
	t.Focused.Directory = t.Focused.Directory.Foreground(Indigo)
	t.Focused.Description = t.Focused.Description.Foreground(MutedFg)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(Fuchsia)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(Fuchsia)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(Fuchsia)
	t.Focused.Option = t.Focused.Option.Foreground(NormalFg)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(Fuchsia)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(LightBlue)
	t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(LightBlue).SetString("✓ ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(MutedFg).SetString("• ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(NormalFg)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(Cream).Background(Fuchsia)
	t.Focused.Next = t.Focused.FocusedButton
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(NormalFg).Background(BlurredButtonBg)

	t.Focused.TextInput.Text = t.Focused.TextInput.Text.Foreground(DarkBlue)
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(Green)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(PlaceholderFg)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(Fuchsia)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = t.Blurred.NextIndicator.Foreground(MutedFg)
	t.Blurred.PrevIndicator = t.Blurred.PrevIndicator.Foreground(MutedFg)

	return t
}

func SetSpinnerTheme(isDark bool) *spinner.Styles {
	t := &spinner.Styles{}

	t.Spinner = t.Spinner.Foreground(LightBlue).Bold(true).
		PaddingLeft(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(DarkBlue).
		Foreground(LightBlue)

	return t
}

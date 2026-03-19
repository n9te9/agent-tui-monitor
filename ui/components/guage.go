package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/n9te9/agent-tui-monitor/ui/theme"
)

func RenderUsageGauge(rate float64, width int) string {
	if rate < 0 {
		rate = 0
	} else if rate > 1 {
		rate = 1
	}

	filledWidth := int(rate * float64(width))
	emptyWidth := width - filledWidth

	barColor := theme.ColorSuccess
	if rate > 0.8 {
		barColor = theme.ColorDanger
	} else if rate > 0.6 {
		barColor = theme.ColorWarning
	}

	filled := lipgloss.NewStyle().Foreground(barColor).Render(strings.Repeat("■", filledWidth))
	empty := lipgloss.NewStyle().Foreground(theme.ColorGaugeEmpty).Render(strings.Repeat("■", emptyWidth))

	return fmt.Sprintf("[%s%s] %3.0f%%", filled, empty, rate*100)
}

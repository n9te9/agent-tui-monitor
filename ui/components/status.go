package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/n9te9/agent-tui-monitor/domain"
	"github.com/n9te9/agent-tui-monitor/ui/theme"
)

var baseStatusStyle = lipgloss.NewStyle().
	Bold(true).
	Padding(0, 1)

func RenderStatusBadge(status domain.AgentStatus) string {
	var color lipgloss.Color

	switch status {
	case domain.AgentStatusRunning:
		color = theme.ColorSuccess
	case domain.AgentStatusError:
		color = theme.ColorDanger
	case domain.AgentStatusFinished:
		color = theme.ColorWarning
	case domain.AgentStatusStopped:
		color = theme.ColorMuted
	default:
		color = theme.ColorMuted
	}

	text := strings.ToUpper(string(status))

	return baseStatusStyle.Copy().
		Background(color).
		Foreground(lipgloss.Color("#000000")).
		Render(text)
}

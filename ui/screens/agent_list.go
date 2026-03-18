package screens

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/n9te9/agent-tui-monitor/domain"
	"github.com/n9te9/agent-tui-monitor/ui/components"
)

type AgentListModel struct {
	agents []domain.Agent
	cursor int
}

func NewAgentListModel(agents []domain.Agent) *AgentListModel {
	return &AgentListModel{
		agents: agents,
		cursor: 0,
	}
}

func (m AgentListModel) Init() tea.Cmd {
	return nil
}

func (m AgentListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.agents)-1 {
				m.cursor++
			}

		}
	}
	return m, nil
}

func (m AgentListModel) View() string {
	var b strings.Builder
	b.WriteString("Agent List\n")

	for i, agent := range m.agents {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		statusBudge := components.RenderStatusBadge(agent.Status)
		row := fmt.Sprintf("%s %-10s %-15s %s\n", cursor, agent.ID, agent.Name, statusBudge)
		b.WriteString(row)
	}

	b.WriteString("\n[q] Quit [j/k] Navigate\n")
	return b.String()
}

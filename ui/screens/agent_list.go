package screens

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
		case "enter":
			selectedAgent := m.agents[m.cursor]
			return m, func() tea.Msg {
				return PushScreenMessage{
					Screen: NewChatScreen(selectedAgent),
				}
			}
		}
	}
	return m, nil
}

func (m AgentListModel) View() string {
	var b strings.Builder
	titleStyle := lipgloss.NewStyle().Bold(true).Underline(true).Padding(0, 1)
	b.WriteString(titleStyle.Render("Agent List") + "\n\n")

	table := components.RenderTable(m.agents, m.cursor)
	b.WriteString(table)

	footerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#4C566A")).Padding(1, 0, 0, 1)
	b.WriteString(footerStyle.Render("<Enter> Chat  <q> Quit   <j/k> Navigate"))
	b.WriteString("\n")

	return b.String()
}

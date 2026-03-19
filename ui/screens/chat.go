package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/n9te9/agent-tui-monitor/domain"
)

type ChatScreen struct {
	agent domain.Agent
}

func NewChatScreen(agent domain.Agent) *ChatScreen {
	return &ChatScreen{
		agent: agent,
	}
}

func (m ChatScreen) Init() tea.Cmd {
	return nil
}

func (m ChatScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg {
				return PopScreenMessage{}
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

const chatScreenViewTemplate = `Chatting with Agent: %s (%s)

[Esc] Back to List  [Ctrl+C] Quit`

func (m ChatScreen) View() string {
	return fmt.Sprintf(chatScreenViewTemplate, m.agent.Name, m.agent.ID)
}

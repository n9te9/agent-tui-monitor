package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/n9te9/agent-tui-monitor/domain"
	"github.com/n9te9/agent-tui-monitor/internal"
	"github.com/n9te9/agent-tui-monitor/ui/screens"
)

var demoAgents = []domain.Agent{
	{ID: "1", Name: "Agent Alpha hogehogehogehogehogehogehogehogehogehoge", Status: domain.AgentStatusRunning, Provider: &internal.DemoAdapter{}},
	{ID: "2", Name: "Agent Beta", Status: domain.AgentStatusError, Provider: &internal.DemoAdapter{}},
	{ID: "3", Name: "Agent Gamma", Status: domain.AgentStatusFinished, Provider: &internal.DemoAdapter{}},
	{ID: "4", Name: "Agent Delta", Status: domain.AgentStatusStopped, Provider: &internal.DemoAdapter{}},
}

func main() {
	m := screens.NewRootModel(screens.NewAgentListModel(demoAgents))

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

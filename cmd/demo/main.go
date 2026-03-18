package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/n9te9/agent-tui-monitor/domain"
	"github.com/n9te9/agent-tui-monitor/ui/screens"
)

var demoAgents = []domain.Agent{
	{ID: "1", Name: "Agent Alpha", Status: domain.AgentStatusRunning},
	{ID: "2", Name: "Agent Beta", Status: domain.AgentStatusError},
	{ID: "3", Name: "Agent Gamma", Status: domain.AgentStatusFinished},
	{ID: "4", Name: "Agent Delta", Status: domain.AgentStatusStopped},
}

func main() {
	m := screens.NewAgentListModel(demoAgents)

	// Bubble Teaのプログラムを生成して実行
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

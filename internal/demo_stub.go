package internal

import "github.com/n9te9/agent-tui-monitor/domain"

type DemoAdapter struct{}

var _ domain.ModelProvider = (*DemoAdapter)(nil)

func NewDemoAdapter() *DemoAdapter {
	return &DemoAdapter{}
}

func (a *DemoAdapter) Name() string {
	return "Demo Model"
}

func (a *DemoAdapter) MaxContextTokens() int {
	return 4096
}

func (a *DemoAdapter) InputTokens() int {
	return 100
}

func (a *DemoAdapter) OutputTokens() int {
	return 200
}

func (a *DemoAdapter) TokenCost(input, output int) float64 {
	return float64(input+output) * 0.0001
}

func (a *DemoAdapter) SupportsThinking() bool {
	return true
}

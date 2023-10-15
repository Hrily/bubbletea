package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titles = []string{
		"Searching for something...",
	}

	headerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}).
			Margin(2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Margin(0, 2)

	messagesStyle = lipgloss.NewStyle().Padding(2).AlignVertical(lipgloss.Bottom)
)

type model struct {
	title    int
	messages []string
}

type somethingReadyMsg int
type somethingLoaded bool

func newModel() *model {
	return &model{}
}

func (m *model) Init() tea.Cmd {
	return doSomething
}

func (m *model) View() string {
	title := titleStyle.Render(titles[m.title])
	return fmt.Sprintln(headerStyle.Render(title), "\n", messagesStyle.Render(strings.Join(m.messages, "\n")))
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}
	case somethingReadyMsg:
		return m, tea.Sequence(tea.Sequence(func() tea.Msg {
			time.Sleep(2 * time.Second)
			return somethingLoaded(false)
		}), func() tea.Msg {
			time.Sleep(1 * time.Second)
			return somethingLoaded(true)
		})
	case somethingLoaded:
		if msg {
			m.messages = append(m.messages, "Something is loaded!")
		} else {
			m.messages = append(m.messages, "Something is not loaded!")
		}
	}
	return m, nil
}

func doSomething() tea.Msg {
	return somethingReadyMsg(1)
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		panic(err)
	}
}

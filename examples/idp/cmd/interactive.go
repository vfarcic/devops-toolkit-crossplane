/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive walkthrough",
	Long:  `Interactive walkthrough`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel())
		if err := p.Start(); err != nil {
			fmt.Printf("Error, run for your lives... %v", err)
			os.Exit(1)
		}
	},
}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Help  key.Binding
	Quit  key.Binding
}

type model struct {
	choices   []string
	cursor    int
	selected  map[int]struct{}
	altscreen bool
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
func initialModel() model {
	xrds := getXrds()
	choices := []string{}
	for _, xrd := range xrds.Items {
		choices = append(choices, xrd.Metadata.Name)
	}
	return model{
		choices:   choices,
		selected:  make(map[int]struct{}),
		altscreen: true,
	}
}

func (m model) Init() tea.Cmd {
	var cmd tea.Cmd
	cmd = tea.EnterAltScreen
	return cmd
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "m":
			var cmd tea.Cmd
			if m.altscreen {
				cmd = tea.ExitAltScreen
			} else {
				cmd = tea.EnterAltScreen
			}
			m.altscreen = !m.altscreen
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Which Crossplane Resource Definition (XRD) would you like to use?\n\n"
	for i, choice := range m.choices {
		cursor := ""
		if m.cursor == i {
			cursor = " <---"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("* [%s] %s %s\n", checked, choice, cursor)
	}
	s += "\nq to quit."
	s += "\nm to switch mode."
	s += "\n"
	renderer, _ := glamour.NewTermRenderer(glamour.WithStylePath("notty"))
	out, _ := renderer.Render(s)
	return out
}

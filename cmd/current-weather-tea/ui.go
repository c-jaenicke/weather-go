package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"weather-go/pkg/weather"
)

type Model struct {
	textInput     textinput.Model // location input
	forecastTable table.Model
	forecast      string
	location      string
	weather       string
	err           string
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func initialModel() Model {
	input := textinput.New()
	input.Placeholder = "Location"
	input.Focus()
	input.Width = 300

	return Model{
		textInput: input,
		forecast:  "",
		location:  "",
		weather:   "",
		err:       "",
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// If input is empty, do nothing
			if m.textInput.Value() == "" {
				m.err = "Error: No location given!"
				return m, cmd
			} else {
				m.location = m.textInput.Value()
				m.weather = currentWeather(m)
				m.forecast = getForecast(m)
				m.forecastTable = m.makeTable()
			}

		}

	}

	// Refresh input field on input
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.location == "" {
		return fmt.Sprintf("Location %s"+
			"\n%s"+
			"\n\nPress CTRL+C or q to quits", m.textInput.View(), m.err)
	} else {
		views := []string{
			textbox.Copy().Align(lipgloss.Left).Render(m.weather),
			textbox.Copy().Align(lipgloss.Left).Render(m.forecastTable.View()),
		}

		return fmt.Sprintf("Location %s"+
			"\n\n%s"+
			"\n\nPress CTRL+C or q to quit", m.textInput.View(), lipgloss.JoinHorizontal(lipgloss.Top, views...))

	}
}

func (m Model) makeTable() table.Model {
	forecastObj, _ := weather.GetForecast(m.location)

	rows := []table.Row{}
	for _, item := range forecastObj.List {
		rows = append(rows, table.Row{unixToDateTime(item.Dt), shortenFloat(item.Main.Temp)})
	}

	columns := []table.Column{
		{Title: "Time", Width: 20},
		{Title: "Temperature", Width: 10},
	}
	forecastTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
	)

	return forecastTable
}

var textbox = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")).
	Width(75)

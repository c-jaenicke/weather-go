package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"strconv"
	"time"
	"weather-go/internal/weather"
)

type Model struct {
	textInput     textinput.Model // location input
	forecastTable table.Model     // table for displaying forecast
	location      string          // location to get weather and forecast for
	weather       string          // current weather string
	err           error           // optional error string to display to user
	timeRequest   string          // timestamp of last request as string
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func initialModel() Model {
	input := textinput.New()

	location := ""
	if flag.Lookup("location").Value.String() == "" {
		location = ""
		input.Placeholder = "e.g. New York (expected format <city_name,state_code,country_code>)"
	} else {
		location = flag.Lookup("location").Value.String()
		input.Placeholder = location
		input.SetValue(location)
	}

	input.Focus()
	input.Width = widthMax

	return Model{
		textInput: input,
		location:  location,
		weather:   "",
		err:       fmt.Errorf(""),
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
				m.err = fmt.Errorf("no location given")
				return m, cmd
			} else {
				m.location = m.textInput.Value()
				m.weather, m.location, m.err = currentWeather(m)
				m.forecastTable = m.makeTable()
				m.timeRequest = m.getTime()
			}

		}
	}

	// Refresh input field on input
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	const helpString = "Press Enter to confirm your input or refresh. Press ctrl+c or q to quit!"

	if m.location == "" {
		views := []string{
			title.Copy().Align(lipgloss.Left).Render("Weather-go"),
			inputbox.Copy().Align(lipgloss.Left).Render(m.textInput.View()),
		}

		return fmt.Sprintf("%s"+
			"\nEnter a location to start"+
			"%s"+
			"\n\n%s", lipgloss.JoinVertical(lipgloss.Top, views...), m.err, helpString)

	} else {
		var views []string
		if m.err != nil {
			views = []string{
				title.Copy().Align(lipgloss.Left).Render("Weather-go"),
				inputbox.Copy().Align(lipgloss.Left).Render(m.textInput.View()),
				warning.Copy().Align(lipgloss.Center).Render(m.err.Error()),
			}
		} else {
			views = []string{
				title.Copy().Align(lipgloss.Left).Render("Weather-go"),
				inputbox.Copy().Align(lipgloss.Left).Render(m.textInput.View()),
				box.Copy().Align(lipgloss.Left).Render("Weather for " + m.location),
				box.Copy().Align(lipgloss.Left).Render(m.weather),
				box.Copy().Align(lipgloss.Left).Render(m.forecastTable.View()),
			}
		}

		return fmt.Sprintf("%s"+
			"\n\n%s\n%s", lipgloss.JoinVertical(lipgloss.Top, views...), helpString, fmt.Sprintf("Last Refresh at %s", m.timeRequest))
	}
}

// makeTable creates table for forecast information
func (m Model) makeTable() table.Model {
	forecastObj, _, err := weather.GetForecast(m.location)
	if err != nil {
		log.Println(err)
		return table.New()
	}

	rows := []table.Row{}
	for _, item := range forecastObj.List {
		rows = append(rows, table.Row{unixToDateTime(item.Dt), shortenFloat(item.Main.Temp), item.Weather[0].Description, strconv.Itoa(item.Main.Humidity), strconv.Itoa(item.Main.Pressure)})
	}

	columns := []table.Column{
		{Title: "Time", Width: 30},
		{Title: "Temp.", Width: 5},
		{Title: "Weather", Width: 20},
		{Title: "Humid.", Width: 5},
		{Title: "Pressure", Width: 10},
	}
	forecastTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(12),
	)

	// set styling of table
	tableStyle.Header = tableStyle.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("foreground")).BorderBottom(true).Bold(true)
	tableStyle.Selected = tableStyle.Selected.Foreground(lipgloss.Color("foreground")).Bold(false)
	forecastTable.SetStyles(tableStyle)

	return forecastTable
}

// getTime gets current time, used for tracking last refresh
func (m Model) getTime() string {
	dt := time.Now()
	return dt.Format("2006-01-02 15-01-05")
}

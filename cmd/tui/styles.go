package main

import "github.com/charmbracelet/lipgloss"

var box = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("12")).
	Width(100)

var inputbox = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("foreground")).
	Width(100)

var title = lipgloss.NewStyle().
	Bold(true).
	PaddingBottom(1).
	Foreground(lipgloss.Color("6"))

package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var widthMax = 80

var box = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("12")).
	Width(widthMax)

var inputbox = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("foreground")).
	Width(widthMax)

var title = lipgloss.NewStyle().
	Bold(true).
	MarginBottom(1).
	MarginTop(1).
	PaddingLeft(1).
	PaddingRight(1).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("6"))

// table style
var tableStyle = table.DefaultStyles()

var warning = lipgloss.NewStyle().
	Foreground(lipgloss.Color("foreground")).
	Background(lipgloss.Color("1")).
	PaddingLeft(1).
	PaddingRight(1)

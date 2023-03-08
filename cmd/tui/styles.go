// Package main
// Contains lipgloss styles for elements.
// See https://github.com/charmbracelet/lipgloss for more information.
package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// set max width for most elements
var widthMax = 80

// standard box for elements to be in
var box = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("12")).
	Width(widthMax)

// inputbox for input elements, differentiated by white border
var inputbox = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("foreground")).
	Width(widthMax)

// title element, with cyan background
var title = lipgloss.NewStyle().
	Bold(true).
	MarginBottom(1).
	MarginTop(1).
	PaddingLeft(1).
	PaddingRight(1).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("6"))

// table style, set by each table individually on creation
var tableStyle = table.DefaultStyles()

// warning for error display, red foreground color
var warning = lipgloss.NewStyle().
	Foreground(lipgloss.Color("foreground")).
	Background(lipgloss.Color("1")).
	PaddingLeft(1).
	PaddingRight(1)

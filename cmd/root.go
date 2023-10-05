// Package cmd is a root command.
/*
Copyright © 2023 Takafumi Miyanaga <miya.org.0309@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/dmowcomber/go-clear"
	"github.com/spf13/cobra"
)

var color string
var speed int32

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "paclear",
	Short: "paclear is a clear command with pacman animation",
	Run:   paclear,
}

func paclear(cmd *cobra.Command, args []string) {
	styledPac := style(color, pac)
	rows, cols := getSize()
	width := len(pac[0])
	height := len(pac)
	if speed < 1 {
		speed = 1
	}
	pitch := time.Duration(20 / time.Duration(speed) * time.Millisecond)
	for y := 0; y <= rows-height; y += height {
		for x := 0; x <= cols-width/3; x++ {
			for j, line := range styledPac {
				fmt.Printf("\033[%d;%dH%s", y+j+1, x, line)
			}
			time.Sleep(pitch)
			for k := 0; k < height; k++ {
				fmt.Printf("\033[%d;%dH%s", y+k+1, x, strings.Repeat(" ", width))
			}
		}
	}
	_ = clear.Clear()
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&color, "color", "c", "default", "Set pacman color (available: red, green, blue, yellow, pink)")
	rootCmd.Flags().Int32VarP(&speed, "speed", "s", 1, "Set pacman multiple speed (default: 1)")
}

func getSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	var rows, cols int
	fmt.Sscan(string(out), &rows, &cols)
	return rows, cols
}

func style(color string, lines []string) []string {
	var styled []string
	var style lipgloss.Style

	switch color {
	case "red":
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	case "green":
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	case "blue":
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF"))
	case "yellow":
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))
	case "pink":
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFC0CB"))
	default:
		style = lipgloss.NewStyle()
	}

	for _, line := range lines {
		styled = append(styled, style.Render(line))
	}
	return styled
}

var pac = []string{
	"	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣤⣤⣤⣤⣤⣀⣀⠀⠀⠀⠀⠀⠀⠀",
	"	⠀⠀⠀⠀⠀⠀⠀⢀⣠⣶⡾⠿⠛⠋⠉⠉⠉⠉⠉⠙⠛⠿⢷⣦⣄⡀⠀⠀",
	"	⠀⠀⠀⠀⠀⣠⣶⠿⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣦⣀",
	"	⠀⠀⠀⢠⣾⠟⠁⠀⠀⠀⠀⠀⠀⣴⡿⠿⣶⡀⠀⠀⠀⠀⠀⠀⠀⣠⣿⠟",
	"	⠀⠀⣰⡿⠋⠀⠀⠀⠀⠀⠀⠀⠘⣿⣇⣀⣿⠇⠀⠀⠀⠀⢀⣠⣾⠟⠁⠀",
	"	⠀⣰⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⠛⠋⠀⠀⠀⢀⣴⡿⠛⠁⠀⠀⠀",
	"	⢠⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⡿⠋⠀⠀⠀⠀⠀⠀",
	"	⣸⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⠟⠉⠀⠀⠀⠀⠀⠀⠀⠀",
	"	⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣾⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀",
	"	⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⢿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀",
	"	⢹⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀",
	"	⠘⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠿⣷⣄⠀⠀⠀⠀⠀⠀",
	"	⠀⠹⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣷⣤⡀⠀⠀⠀",
	"	⠀⠀⠹⣿⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣦⡀⠀",
	"	⠀⠀⠀⠘⢿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣿⣦",
	"	⠀⠀⠀⠀⠀⠙⠿⣶⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣾⠟⠉",
	"	⠀⠀⠀⠀⠀⠀⠀⠈⠙⠿⢷⣶⣤⣄⣀⣀⣀⣀⣀⣠⣤⣶⡾⠟⠋⠁⠀⠀",
	"	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠛⠛⠛⠛⠛⠉⠉⠀⠀⠀⠀⠀⠀⠀",
}

// SetVersionInfo sets version and date to rootCmd
func SetVersionInfo(version, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s)", version, date)
}

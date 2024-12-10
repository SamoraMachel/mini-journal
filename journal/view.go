package journal

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)

	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

func (m model) View() string {
	s := appNameStyle.Render("JOURNAL APP") + "\n\n"

	if m.err != "" {
		s += errorStyle.Render("Error: ") + errorStyle.Render(m.err) + "\n\n"
	}

	if m.state == titleView {
		s += "Journal Title: \n\n"
		s += m.textinput.View() + "\n\n"
		s += faintStyle.Render("enter - save, esc - discard")

	}

	if m.state == bodyView {
		s += "Journal: \n\n"
		s += m.textarea.View() + "\n\n"
		s += faintStyle.Render("ctrl+s - save, esc - discard")

	}

	if m.state == listView {
		for i, n := range m.journals {
			prefix := " "
			if i == m.listIndex {
				prefix = "->"
			}

			shortBody := strings.ReplaceAll(n.Text, "\n", "")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}

			s += enumeratorStyle.Render(prefix) + n.Title + " | " + faintStyle.Render(shortBody) + "\n\n"
		}

		if (m.isAuthed) {
			s += faintStyle.Render("n - New Journal, q - Quit")
		} else {
			s += faintStyle.Render("a - authenticate, q - quit")
		}
	}

	if m.state == usernameView {
		s += "Enter your username: \n"
		s += m.textinput.View() + "\n\n"
		s += faintStyle.Render("enter - proceed, ctrl+q - Quit")
	}

	if m.state == passwordView {
		s += "Enter your Password: \n"
		s += m.textinput.View() + "\n\n"
		s += faintStyle.Render("ctrl+r - Register, ctrl+l - Login, ctrl+q - Quit")
	}

	return s
}

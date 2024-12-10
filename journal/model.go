package journal

import (
	"journal/auth"
	"journal/entity"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView  uint = iota
	titleView      = 1
	bodyView       = 2
	usernameView	= 3
	passwordView = 4
)

type model struct {
	username string
	password string
	profileKey string
	err string
	state uint
	currJournal entity.JournalModel
	journals []entity.JournalModel
	isAuthed bool
	listIndex int
	textarea textarea.Model
	textinput textinput.Model
}


func NewModel(profileKey string) model {
	journalsList, err := RetrieveJournal(profileKey)
	if err != nil {
		return model{
			username: "",
			password: "",
			err: "",
			state: listView,
			isAuthed: false,
			listIndex: 0,
			textarea: textarea.New(),
			textinput: textinput.New(),
		}
	}
	return model{
		profileKey: profileKey,
		state: listView,
		isAuthed: true,
		listIndex: 0,
		err: "",
		journals: journalsList,
		textarea: textarea.New(),
		textinput: textinput.New(),
	}
	
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd tea.Cmd
	)
	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
		case tea.KeyMsg:
			key := msg.String()
			switch m.state {
			case listView:
				switch key {
				case "q":
					return m, tea.Quit
					
				case "n":
					m.textinput.SetValue("")
					m.textinput.Focus()
					m.currJournal = entity.JournalModel{}
					m.state = titleView

				case "a":
					m.textinput.SetValue("")
					m.textinput.Focus()
					m.state = usernameView

				case "up", "k":
					if m.listIndex > 0 {
						m.listIndex--
					}
				
				case "down", "j":
					if m.listIndex < len(m.journals)-1 {
						m.listIndex++
					}

				case "enter":
					m.currJournal = m.journals[m.listIndex]
					m.textarea.SetValue(m.currJournal.Text)
					m.textarea.Focus()
					m.textarea.CursorEnd()
					m.state = bodyView
					// show the text area
				}

			case titleView:
				switch key {
				case "enter":
					title := m.textinput.Value()
					if title != "" {
						m.currJournal.Title = title

						m.textarea.SetValue("")
						m.textarea.Focus()
						m.textarea.CursorEnd()

						m.state = bodyView
					}

				case "esc":
					m.state = listView
				}

			case bodyView:
				switch key {
				case "ctrl+s":
					body := m.textarea.Value()
					m.currJournal.Text = body

					var err error
					if err = AddToJournal(m.profileKey, m.currJournal); err != nil {
						m.err = err.Error()
					}

					m.journals, err = RetrieveJournal(m.profileKey)
					if err != nil {
						m.err = err.Error()
					}

					m.currJournal = entity.JournalModel{}
					m.state = listView
				case "esc":
					m.state = listView
				}

			case usernameView:
				switch key {
				case "enter":
					username := m.textinput.Value()

					if username != "" {
						m.username = username
						m.textinput.SetValue("")
						m.textinput.Focus()
						m.state = passwordView
					}

				case "ctrl+q":
					return m, tea.Quit
				}

			case passwordView:
				switch key {
				case "ctrl+r":
					password := m.textinput.Value()

					if password != "" {
						m.password = password
						err := auth.Register(m.username, m.password, "New", "User")
						if err != nil {
							m.err = err.Error()
							m.state = usernameView
						}

						if err == nil {
							m.isAuthed = true
							m.profileKey = auth.GenerateKey(m.username, m.password)
							CreateJournal(m.profileKey)
							m.state = listView
						}
					}

				case "ctrl+l":
					password := m.textinput.Value()

					if password != "" {
						m.password = password
						_, err := auth.Login(m.username, m.password)
						if err != nil {
							m.err = err.Error()
							m.state = listView
						}

						if err == nil {
							m.isAuthed = true
							m.profileKey = auth.GenerateKey(m.username, m.password)
							m.state = listView
						}
					}
					
				case "ctrl+q":
					return m, tea.Quit
				}

		}
	}

	return m, tea.Batch(cmds...)
}

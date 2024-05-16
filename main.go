package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dunkbing/kana/constants"
	"github.com/muesli/termenv"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var staticFiles embed.FS

var term = termenv.ColorProfile()

func main() {
	var kanaType string
	args := os.Args[1:]
	runServer := false

	for _, arg := range args {
		switch arg {
		case "--help":
			fmt.Println(constants.Usage)
			return
		case "--kata":
			kanaType = constants.Katakana
		case "--hira":
			kanaType = constants.Hiragana
		case "serve":
			runServer = true
		default:
			fmt.Println("Unknown option: " + arg)
			fmt.Println()
			fmt.Println(constants.Usage)
			return
		}
	}

	if kanaType == "" {
		kanaType = constants.Both
	}

	if runServer {
		startServer()
	} else {
		p := tea.NewProgram(initialModel(kanaType))
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, constants.Both)
	})
	http.HandleFunc("/katakana", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, constants.Katakana)
	})
	http.HandleFunc("/hiragana", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, constants.Hiragana)
	})
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request, kanaType string) {
	var tmplFile string
	switch kanaType {
	case constants.Katakana:
		tmplFile = "templates/katakana.html"
	case constants.Hiragana:
		tmplFile = "templates/hiragana.html"
	default:
		tmplFile = "templates/index.html"
	}
	navTmpl := "templates/navbar.html"
	tmpl, err := template.ParseFS(templates, tmplFile, navTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"KanaType": kanaType,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type model struct {
	textInput   textinput.Model
	currentWord []string
	status      string
	points      int
	kanaType    string
}

func initialModel(kanaType string) model {
	i := textinput.New()
	i.Placeholder = "Type the Romaji representation and press Enter 👆"
	i.Focus()
	i.Reset()
	i.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4e4e4e"))

	return model{
		textInput:   i,
		currentWord: newWord(kanaType),
		kanaType:    kanaType,
	}
}

func (m model) Init() tea.Cmd {
	m.currentWord = newWord(m.kanaType)
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit

		case "ctrl+h":
			if m.kanaType == constants.Hiragana {
				return m, cmd
			}
			m.kanaType = constants.Hiragana
			m.currentWord = newWord(constants.Hiragana)

		case "ctrl+k":
			if m.kanaType == constants.Katakana {
				return m, cmd
			}
			m.kanaType = constants.Katakana
			m.currentWord = newWord(constants.Katakana)

		case "ctrl+b":
			if m.kanaType == constants.Both {
				return m, cmd
			}
			m.kanaType = constants.Both
			m.currentWord = newWord(constants.Both)

		case "enter":
			correctAns := toRomaji(m.currentWord)
			ans := strings.ToLower(m.textInput.Value())
			if ans == correctAns {
				m.status = "🎉 Correct!"
				m.points++
			} else {
				m.status = fmt.Sprintf("😭 Incorrect. The answer is %s", correctAns)
			}
			m.textInput.Reset()
			m.currentWord = newWord(m.kanaType)

		default:
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

	default:
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	statusMsg := fmt.Sprintf("%v (Points: %v)", m.status, m.points)

	hiraganaMode := termenv.String(constants.Hiragana).Foreground(term.Color("10")).String()
	if m.kanaType != constants.Hiragana {
		hiraganaMode = "ctrl-(h)iragana"
	}

	katakanaMode := termenv.String(constants.Katakana).Foreground(term.Color("10")).String()
	if m.kanaType != constants.Katakana {
		katakanaMode = "ctrl-(k)atakana"
	}

	bothMode := termenv.String(constants.Both).Foreground(term.Color("10")).String()
	if m.kanaType != constants.Both {
		bothMode = "ctrl-(b)oth"
	}

	modeMsg := fmt.Sprintf("%s %s %s %s", termenv.String("Kana mode: ").Foreground(term.Color("205")).String(), hiraganaMode, katakanaMode, bothMode)

	currentWord := strings.Join(m.currentWord, "")
	return fmt.Sprintf(
		"\n%s\n\n%s\n\n%s\n\n%s\n%s\n\n",
		termenv.String("Kana Word: ").Foreground(term.Color("205")).String()+currentWord,
		m.textInput.View(),
		statusMsg,
		modeMsg,
		"(esc or ctrl-c to quit)",
	)
}

func newWord(kanaType string) []string {
	var kanaChars []string

	switch kanaType {
	case constants.Katakana:
		kanaChars = constants.KatakanaChars
	case constants.Hiragana:
		kanaChars = constants.HiraganaChars
	default:
		bothChars := constants.HiraganaChars[:0]
		bothChars = append(bothChars, constants.HiraganaChars...)
		bothChars = append(bothChars, constants.KatakanaChars...)
		kanaChars = bothChars
	}

	word := make([]string, rand.Intn(5)+1)
	for i := range word {
		word[i] = kanaChars[rand.Intn(len(kanaChars))]
	}
	return word
}

func toRomaji(s []string) string {
	romaji := ""
	for _, r := range s {
		romaji += constants.KanaMap[r]
	}
	return romaji
}

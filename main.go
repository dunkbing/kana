package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"log"
	"math/rand"
)

const usage = `kana-practice

This app displays a random Katakana or Hiragana word, and you need to type the corresponding Romaji representation. Press Enter to submit your answer.

Example:
Word displayed: あい
You type: ai (then press Enter)
`

var term = termenv.ColorProfile()

func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	textInput   textinput.Model
	currentWord string
	status      string
	points      int
}

func initialModel() model {
	i := textinput.New()
	i.Placeholder = "Type the Romaji representation and press Enter 👆"
	i.Focus()
	i.Reset()

	return model{
		textInput:   i,
		currentWord: newWord(),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.textInput.Value() == toRomaji(m.currentWord) {
				m.status = "🎉 Correct!"
				m.points++
				m.textInput.Reset()
				m.currentWord = newWord()
			} else {
				m.status = "😭 Incorrect"
			}

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
	statusMsg := m.status
	if statusMsg == "" {
		statusMsg = "Score: " + fmt.Sprintf("%d", m.points)
	}

	return fmt.Sprintf("\n\n%s\n\n%s\n\n%s\n\n%s\n\n",
		termenv.String("Kana Word: ").Foreground(term.Color("205")).String()+m.currentWord,
		m.textInput.View(),
		statusMsg,
		"(esc or ctrl-c to quit)",
	)
}

func newWord() string {
	kanaChars := []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽぁぃぅぇぉゃゅょっ")
	word := make([]rune, rand.Intn(5)+1)
	for i := range word {
		word[i] = kanaChars[rand.Intn(len(kanaChars))]
	}
	return string(word)
}

var kanaMap = map[rune]string{
	'あ': "a", 'い': "i", 'う': "u", 'え': "e", 'お': "o",
	'か': "ka", 'き': "ki", 'く': "ku", 'け': "ke", 'こ': "ko",
	'さ': "sa", 'し': "shi", 'す': "su", 'せ': "se", 'そ': "so",
	'た': "ta", 'ち': "chi", 'つ': "tsu", 'て': "te", 'と': "to",
	'な': "na", 'に': "ni", 'ぬ': "nu", 'ね': "ne", 'の': "no",
	'は': "ha", 'ひ': "hi", 'ふ': "fu", 'へ': "he", 'ほ': "ho",
	'ま': "ma", 'み': "mi", 'む': "mu", 'め': "me", 'も': "mo",
	'や': "ya", 'ゆ': "yu", 'よ': "yo",
	'ら': "ra", 'り': "ri", 'る': "ru", 'れ': "re", 'ろ': "ro",
	'わ': "wa", 'を': "o", 'ん': "n",
	'が': "ga", 'ぎ': "gi", 'ぐ': "gu", 'げ': "ge", 'ご': "go",
	'ざ': "za", 'じ': "ji", 'ず': "zu", 'ぜ': "ze", 'ぞ': "zo",
	'だ': "da", 'ぢ': "ji", 'づ': "zu", 'で': "de", 'ど': "do",
	'ば': "ba", 'び': "bi", 'ぶ': "bu", 'べ': "be", 'ぼ': "bo",
	'ぱ': "pa", 'ぴ': "pi", 'ぷ': "pu", 'ぺ': "pe", 'ぽ': "po",
	'ぁ': "a", 'ぃ': "i", 'ぅ': "u", 'ぇ': "e", 'ぉ': "o",
	'ゃ': "ya", 'ゅ': "yu", 'ょ': "yo", 'っ': "tsu",
}

func toRomaji(s string) string {
	romaji := ""
	for _, r := range s {
		romaji += kanaMap[r]
	}
	return romaji
}

package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"log"
	"math/rand"
	"os"
)

const usage = `Usage: kana [--help] [--katakana] [--hiragana]

Options:
  --help  Show this help message and exit
  --kata  Practice Katakana words
  --hira  Practice Hiragana words

If no option is provided, both Katakana and Hiragana words will be displayed.

This app displays a random Katakana or Hiragana word, and you need to type the corresponding Romaji representation. Press Enter to submit your answer.

Example:
Word displayed: あい
You type: ai (then press Enter)`

const (
	hiraganaChars = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽぁぃぅぇぉゃゅょっ"
	katakanaChars = "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヰヱヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ"
)

var term = termenv.ColorProfile()

func main() {
	var kanaType string
	args := os.Args[1:]

	for _, arg := range args {
		switch arg {
		case "--help":
			fmt.Println(usage)
			return
		case "--kata":
			kanaType = "katakana"
		case "--hira":
			kanaType = "hiragana"
		default:
			fmt.Println("Unknown option: " + arg)
			fmt.Println()
			fmt.Println(usage)
			return
		}
	}

	if kanaType == "" {
		kanaType = "both"
	}

	p := tea.NewProgram(initialModel(kanaType))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	textInput   textinput.Model
	currentWord string
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

		case "enter":
			if m.textInput.Value() == toRomaji(m.currentWord) {
				m.status = "🎉 Correct!"
				m.points++
				m.textInput.Reset()
				m.currentWord = newWord(m.kanaType)
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
	statusMsg := fmt.Sprintf("%v (Points: %v)", m.status, m.points)

	return fmt.Sprintf(
		"\n%s\n\n%s\n\n%s\n\n%s\n\n",
		termenv.String("Kana Word: ").Foreground(term.Color("205")).String()+m.currentWord,
		m.textInput.View(),
		statusMsg,
		"(esc or ctrl-c to quit)",
	)
}

func newWord(kanaType string) string {
	var kanaChars []rune

	switch kanaType {
	case "katakana":
		kanaChars = []rune(katakanaChars)
	case "hiragana":
		kanaChars = []rune(hiraganaChars)
	default:
		bothChars := hiraganaChars + katakanaChars
		kanaChars = []rune(bothChars)
	}

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
	'ア': "a", 'イ': "i", 'ウ': "u", 'エ': "e", 'オ': "o",
	'カ': "ka", 'キ': "ki", 'ク': "ku", 'ケ': "ke", 'コ': "ko",
	'サ': "sa", 'シ': "shi", 'ス': "su", 'セ': "se", 'ソ': "so",
	'タ': "ta", 'チ': "chi", 'ツ': "tsu", 'テ': "te", 'ト': "to",
	'ナ': "na", 'ニ': "ni", 'ヌ': "nu", 'ネ': "ne", 'ノ': "no",
	'ハ': "ha", 'ヒ': "hi", 'フ': "fu", 'ヘ': "he", 'ホ': "ho",
	'マ': "ma", 'ミ': "mi", 'ム': "mu", 'メ': "me", 'モ': "mo",
	'ヤ': "ya", 'ユ': "yu", 'ヨ': "yo",
	'ラ': "ra", 'リ': "ri", 'ル': "ru", 'レ': "re", 'ロ': "ro",
	'ワ': "wa", 'ヰ': "i", 'ヱ': "e", 'ヲ': "o", 'ン': "n",
	'ガ': "ga", 'ギ': "gi", 'グ': "gu", 'ゲ': "ge", 'ゴ': "go",
	'ザ': "za", 'ジ': "ji", 'ズ': "zu", 'ゼ': "ze", 'ゾ': "zo",
	'ダ': "da", 'ヂ': "ji", 'ヅ': "zu", 'デ': "de", 'ド': "do",
	'バ': "ba", 'ビ': "bi", 'ブ': "bu", 'ベ': "be", 'ボ': "bo",
	'パ': "pa", 'ピ': "pi", 'プ': "pu", 'ペ': "pe", 'ポ': "po",
}

func toRomaji(s string) string {
	romaji := ""
	for _, r := range s {
		romaji += kanaMap[r]
	}
	return romaji
}

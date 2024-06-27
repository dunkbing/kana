# kana practice
Terminal app to practice typing Kana (Japanese characters) in Romaji. Built with [bubbletea](https://github.com/charmbracelet/bubbletea)

### Demo
![demo](./demo/demo.gif)

### Install

#### Golang

```bash
go install github.com/dunkbing/kana@latest
```

or run it directly

```bash
go run github.com/dunkbing/kana@latest
```

#### Homebrew

```bash
brew install dunkbing/brews/kana
```

### Standalone Binary

Download latest archive `*.tar.gz` for your target platform from [the releases page](https://github.com/dunkbing/kana/releases/latest).

### Source

```bash
git clone https://github.com/dunkbing/kana.git
cd kana
go build -o kana .
cp kana /usr/local/bin
chmod +x /usr/local/bin/kana

# kana
```

### How to use

```bash
Usage: kana [--help] [--kata] [--hira]

Options:
  --help  Show this help message and exit
  --kata  Practice Katakana words
  --hira  Practice Hiragana words

If no option is provided, both Katakana and Hiragana words will be displayed.

This app displays a random Katakana or Hiragana word, and you need to type the corresponding Romaji representation. Press Enter to submit your answer.

Example:
Word displayed: あい
You type: ai (then press Enter)
 ```

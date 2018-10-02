package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	pinyin "github.com/mozillazg/go-pinyin"
)

var (
	pinyinArgs = pinyin.NewArgs()
)

// Hans2Pinyin
func Hans2Pinyin(str string) string {
	var returns []string
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		if unicode.Is(unicode.Scripts["Han"], r) {
			returns = append(returns, pinyin.SinglePinyin(r, pinyinArgs)...)
		} else {
			returns = append(returns, fmt.Sprintf("%c", r))
		}

		str = str[size:]
	}

	return strings.Join(returns, "")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You should have to less 1 Hans")
		return
	}

	var startIdx = 1
	if strings.HasPrefix(os.Args[1], "-style=") {
		if len(os.Args) < 3 {
			fmt.Println("You should have to less 1 Hans")
			return
		}

		param := os.Args[1][len("-style="):]

		startIdx = 2
		var style = 0
		if strings.EqualFold(param, "Tone") {
			style = pinyin.Tone
		} else if strings.EqualFold(param, "Tone2") {
			style = pinyin.Tone2
		} else if strings.EqualFold(param, "Tone3") {
			style = pinyin.Tone3
		} else if strings.EqualFold(param, "Initials") {
			style = pinyin.Initials
		} else if strings.EqualFold(param, "FirstLetter") {
			style = pinyin.FirstLetter
		} else if strings.EqualFold(param, "Finals") {
			style = pinyin.Finals
		} else if strings.EqualFold(param, "FinalsTone") {
			style = pinyin.FinalsTone
		} else if strings.EqualFold(param, "FinalsTone2") {
			style = pinyin.FinalsTone2
		} else if strings.EqualFold(param, "FinalsTone3") {
			style = pinyin.FinalsTone3
		}
		pinyinArgs.Style = style
	}

	if len(os.Args[startIdx:]) == 1 {
		fmt.Println(Hans2Pinyin(os.Args[startIdx]))
		return
	}
	for i, arg := range os.Args[startIdx:] {
		fmt.Printf("%d: %s\n", i+1, Hans2Pinyin(arg))
	}
}

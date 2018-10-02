# chapinyin

This command tool will convert hans to pinyin depends on [github.com/mozillazg/go-pinyin](https://github.com/mozillazg/go-pinyin).

## Installation

```
go get github.com/lunny/chapinyin
```

## Usage

```
chapinyin <汉字1> <汉字2>
```

```
chapinyin -style=tone3 <汉字1> <汉字2>
```

## Style lists

* Normal      = 0 // 普通风格，不带声调（默认风格）。如： zhong guo
* Tone        = 1 // 声调风格1，拼音声调在韵母第一个字母上。如： zhōng guó
* Tone2       = 2 // 声调风格2，即拼音声调在各个韵母之后，用数字 [1-4] 进行表示。如： zho1ng guo2
* Tone3       = 8 // 声调风格3，即拼音声调在各个拼音之后，用数字 [1-4] 进行表示。如： zhong1 guo2
* Initials    = 3 // 声母风格，只返回各个拼音的声母部分。如： zh g
* FirstLetter = 4 // 首字母风格，只返回拼音的首字母部分。如： z g
* Finals      = 5 // 韵母风格，只返回各个拼音的韵母部分，不带声调。如： ong uo
* FinalsTone  = 6 // 韵母风格1，带声调，声调在韵母第一个字母上。如： ōng uó
* FinalsTone2 = 7 // 韵母风格2，带声调，声调在各个韵母之后，用数字 [1-4] 进行表示。如： o1ng uo2
* FinalsTone3 = 9 // 韵母风格3，带声调，声调在各个拼音之后，用数字 [1-4] 进行表示。如： ong1 uo2
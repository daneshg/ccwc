package ccwc

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

type WCStat struct {
	lines, words, bytes, chars int
}

func (wc *WCStat) GetLines() int {
	return wc.lines
}

func (wc *WCStat) GetBytes() int {
	return wc.bytes
}

func (wc *WCStat) GetWords() int {
	return wc.words
}

func (wc *WCStat) GetChars() int {
	return wc.chars
}

func GetWCFromStdIN() (*WCStat, error) {
	stat, err := wcGetStat(os.Stdin)
	if err != nil {
		return &WCStat{}, err
	}
	return stat, nil
}

func GetWCFromFile(fileName string) (*WCStat, error) {

	fd, err := os.Open(fileName)
	if err != nil {
		return &WCStat{}, err
	}
	defer fd.Close()
	stat, err := wcGetStat(fd)
	if err != nil {
		return &WCStat{}, err
	}

	return stat, nil
}

func wcGetStat(fd *os.File) (*WCStat, error) {

	wcStat := &WCStat{
		lines: 0,
		words: 0,
		bytes: 0,
		chars: 0,
	}

	reader := bufio.NewReader(fd)
	isWord := false

	for {
		ch, n, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if isWord {
					wcStat.words++
					isWord = false
				}
				break
			} else {
				return wcStat, err
			}
		}
		wcStat.bytes += n
		wcStat.chars++
		if unicode.IsSpace(ch) {
			if isWord {
				wcStat.words++
				isWord = false
			}
			if ch == '\n' {
				wcStat.lines++
			}
		} else {
			isWord = true
		}
	}

	return wcStat, nil
}

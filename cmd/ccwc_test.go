package ccwc

import (
	"testing"
)

func TestGetWCFromFile(t *testing.T) {
	file := "testdata/test.txt"
	want := &WCStat{lines: 7145, words: 58164, bytes: 342190, chars: 339292}
	got, err := GetWCFromFile(file)
	if err != nil {
		t.Fatalf(`Failed with err %v`, err)
	}

	if want.bytes != got.bytes || want.lines != got.lines || want.words != got.words || want.chars != got.chars {
		t.Fatalf(`Test case failed want=%v but got=%v`, want, got)
	}
}

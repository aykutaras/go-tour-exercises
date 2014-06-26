package gotourexercises

import (
	"testing"
	"code.google.com/p/go-tour/wc"
)

func TestWordCount(t *testing.T) {
	wordMap := WordCount("dummy dummy")
	wc.Test(WordCount)
	if wordMap["dummy"] != 2 {
		t.Errorf("Splitted word count %v, need %v", wordMap["dummy"], 2)
	}
}
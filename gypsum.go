// Copyright 2013 Antoine Grondin. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gypsum

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func Lorem() string {
	return ManyLorem(1)[0]
}

func ManyLorem(count int) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	text := make([]string, count)
	for i, _ := range text {
		text[i] = lorem[rand.Intn(len(lorem)-1)]
	}
	return text
}

func ArticleLorem(paraCount int, paraSeparator string) string {
	return strings.Join(ManyLorem(paraCount), paraSeparator)
}

func WordLorem(wordCount int) string {
	words := make([]string, wordCount)
	rand.Seed(time.Now().UTC().UnixNano())
	paras := ManyLorem(wordCount)
	r, _ := regexp.Compile("\\W")
	for i, line := range paras {
		wordsInLine := strings.Split(line, " ")
		index := rand.Intn(len(wordsInLine))
		words[i] = r.ReplaceAllString(wordsInLine[index], "")
	}
	return strings.Join(words, " ")
}

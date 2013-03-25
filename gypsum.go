// Copyright 2013 Antoine Grondin. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gypsum

import (
	"math/rand"
)

func Lorem() string {
	return ManyLorem(1)[0]
}

func ManyLorem(count int) []string {
	text := make([]string, count)
	for i, _ := range text {
		text[i] = lorem[rand.Intn(len(lorem)-1)]
	}
	return text
}

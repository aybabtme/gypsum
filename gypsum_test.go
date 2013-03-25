// Copyright 2013 Antoine Grondin. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gypsum

import (
	"testing"
)

func TestLorem(t *testing.T) {
	text := Lorem()
	if text == "" {
		t.Error("Shouldn't be an empty string.")
	}
}

func TestManyLorem(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go testingManyLorem(t, i)
	}
}

func testingManyLorem(t *testing.T, i int) {
	manyText := ManyLorem(i)
	if len(manyText) != i {
		t.Errorf("Expected len(manyText)=<%d> but was <%d>", i, len(manyText))
		return
	}
	for _, line := range manyText {
		if line == "" {
			t.Error("Shouldn't be an empty string.")
			return
		}
	}
}

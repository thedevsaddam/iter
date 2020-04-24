package iter

import (
	"testing"
)

func TestN(t *testing.T) {

	t.Run("Should do nothing", func(t *testing.T) {
		match := 0
		for n := range N() {
			if match != n {
				t.Fail()
			}
			match++
		}
	})

	t.Run("Should generate sequence 1-9", func(t *testing.T) {
		match := 1
		for n := range N(1, 10) {
			if match != n {
				t.Fail()
			}
			match++
		}
	})

	t.Run("Should generate sequence 0-9", func(t *testing.T) {
		match := 0
		for n := range N(10) {
			if match != n {
				t.Fail()
			}
			match++
		}
	})

	t.Run("Should generate sequence 1-9 by step+=2", func(t *testing.T) {
		match := 1
		for n := range N(1, 10, 2) {
			if match != n {
				t.Fail()
			}
			match += 2
		}
	})
}

func TestL(t *testing.T) {
	output := ""
	for c := range L('a', 'e') {
		output += string(c)
	}
	if output != "abcde" {
		t.Fail()
	}
}

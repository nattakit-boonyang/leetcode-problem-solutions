package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1806"
  title: count-of-matches-in-tournament
  lang: golang
  type: large
  inputString: |-
    7
    14
    1
    200
*/

func numberOfMatches(n int) int {
	return n - 1
}

func Test_1806(t *testing.T) {
	NewTestcases(t).
		Add(6, 7).
		Add(13, 14).
		Add(0, 1).
		Add(199, 200).
		Each(func(a *assert.Assertions, td TestData) {
			actual := numberOfMatches(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}

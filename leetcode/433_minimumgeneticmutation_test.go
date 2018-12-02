package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumGeneticMutation(t *testing.T) {
	testcases := []struct {
		Start    string
		End      string
		Bank     []string
		Expected int
	}{
		{
			"AACCGGTT",
			"AACCGGTA",
			[]string{"AACCGGTA"},
			1,
		}, {
			"AACCGGTT",
			"AAACGGTA",
			[]string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			2,
		}, {
			"AAAAACCC",
			"AACCCCCC",
			[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			3,
		}, {
			"AACCGGTT",
			"AACCGGTA",
			[]string{},
			-1,
		},
	}
	for _, tc := range testcases {
		actual := minMutation(tc.Start, tc.End, tc.Bank)
		assert.Equal(t, tc.Expected, actual, "start: %s\nend: %s\nbank: %s", tc.Start, tc.End, tc.Bank)
	}
}

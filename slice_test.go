package chain

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSliceChain(t *testing.T) {
	var lines = []string{
		"the golang generic is ugly",
		"how do you think?",
	}

	data := SliceChain3[string, []string, int](lines).
		Filter(func(s string, i int) bool {
			return len(s) > 10
		}).
		Map(func(s string, _ int) []string {
			return strings.Fields(s)
		}).
		Map(func(s []string, i int) int {
			return len(s)
		}).
		Filter(func(c, _ int) bool {
			return c > 4
		})

	count := lo.Sum(data)
	assert.Equal(t, count, 5)
}

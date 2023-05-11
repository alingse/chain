package chain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceChain(t *testing.T) {
	var lines = []string{
		"the golang generic is ugly",
		"how do you think?",
	}
	var x = NewSliceChain4[string, []string, int64, any](lines).
		Map(func(s string, _ int) []string {
			return strings.Fields(s)
		}).
		Map(func(s []string, i int) int64 {
			return int64(len(s))
		})
	assert.Len(t, x, 2)
	/*
		var x = NewSliceChain3[string, []string, int](lines)
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
	*/
}

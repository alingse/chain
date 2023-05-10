package chain

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSliceChain(t *testing.T) {
	var x SliceChain3[string, []string, int] = []string{
		"the golang generic is ugly",
		"how do you think?",
	}

	data := x.
		Map(func(s string, _ int) []string {
			return strings.Fields(s)
		}).
		Map(func(s []string, i int) int {
			return len(s)
		})

	count := lo.Sum(data)
	assert.Equal(t, count, 9)
}

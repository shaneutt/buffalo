package render

import (
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/stretchr/testify/require"
)

func Test_link(t *testing.T) {
	table := []struct {
		opts tags.Options
		out  string
	}{
		{},
	}

	for _, tt := range table {
		t.Run(tt.out, func(st *testing.T) {
			r := require.New(st)

		})
	}

}

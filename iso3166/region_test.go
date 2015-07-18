package iso3166

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	s1 := strings.Trim("  C hina  ", " ")
	assert.Equal(t, "C hina", s1)
}

func TestParseRegion(t *testing.T) {
	tests := []struct {
		Str, Alpha2, Alpha3, Name string
		Err                       error
	}{
		{"china", "CN", "CHN", "China", nil},
		{"  United states of America", "US", "USA", "United States of America", nil},
	}

	for _, v := range tests {
		r, err := ParseRegion(v.Str)
		assert.Equal(t, v.Err, err)
		assert.Equal(t, v.Alpha2, r.Alpha2)
		assert.Equal(t, v.Alpha3, r.Alpha3)
		assert.Equal(t, v.Name, r.Name)
	}
}

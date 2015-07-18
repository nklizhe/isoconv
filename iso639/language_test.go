package iso639

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLanguage(t *testing.T) {
	tests := []struct {
		Str, Code1, Code2, Name string
		Err                     error
	}{
		{"en", "en", "eng", "English", nil},
		{"  chinese", "zh", "zho", "Chinese", nil},
	}

	for _, v := range tests {
		r, err := ParseLanguage(v.Str)
		assert.Equal(t, v.Err, err)
		assert.Equal(t, v.Code1, r.Code1)
		assert.Equal(t, v.Code2, r.Code2)
		assert.Equal(t, v.Name, r.Name)
	}
}

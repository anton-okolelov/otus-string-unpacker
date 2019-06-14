package string_unpacker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpacker(t *testing.T) {

	tests := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "",
		"a45":      "",
		`qwe\4\5`:  `qwe45`,
		`qwe\45`:   `qwe44444`,
		`qwe\\5`:   `qwe\\\\\`,
	}

	for packed, unpacked := range tests {
		assert.Equal(t, unpacked, Unpack(packed))
	}
}

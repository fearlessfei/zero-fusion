package cachekey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheKeyBuildKey(t *testing.T) {
	tests := []struct {
		nameSpace string
		keyPrefix string
		separator string
		keyParts  []any
		expected  string
	}{
		{
			nameSpace: "test1",
			keyPrefix: "prefix1",
			separator: ":",
			keyParts:  []any{"key1", "key2", 123},
			expected:  "test1:prefix1:key1:key2:123",
		},
		{
			nameSpace: "test2",
			keyPrefix: "prefix2",
			separator: "_",
			keyParts:  []any{"key1", "key2", "123"},
			expected:  "test2_prefix2_key1_key2_123",
		},
		{
			nameSpace: "test3",
			keyPrefix: "prefix3",
			separator: "-",
			keyParts:  []any{"key1", "key2", 123, 456},
			expected:  "test3-prefix3-key1-key2-123-456",
		},
		{
			nameSpace: "",
			keyPrefix: "prefix3",
			separator: "-",
			keyParts:  []any{"key1", "key2", 123, 456},
			expected:  "prefix3-key1-key2-123-456",
		},
		{
			nameSpace: "",
			keyPrefix: "",
			separator: "-",
			keyParts:  []any{"key1", "key2", 123, 456},
			expected:  "key1-key2-123-456",
		},
		{
			nameSpace: "",
			keyPrefix: "",
			separator: "",
			keyParts:  []any{"key1", "key2", 123, 456},
			expected:  "key1key2123456",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.nameSpace, func(t *testing.T) {
			t.Parallel()

			key := BuildKey(
				tt.keyParts,
				WithNameSpace(tt.nameSpace),
				WithKeyPrefix(tt.keyPrefix),
				WithSeparator(tt.separator),
			)

			assert.Equal(t, tt.expected, key)
		})
	}
}

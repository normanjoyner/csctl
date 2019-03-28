package label

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	type stringTest struct {
		key         string
		value       string
		expected    string
		shouldError bool
	}

	var tests = []stringTest{
		{
			key:         "",
			value:       "value1",
			shouldError: true,
		},
		{
			key:         "nodepool.containership.io/key1",
			value:       "value1",
			expected:    "nodepool.containership.io/key1=value1",
			shouldError: false,
		},
		{
			key:         "nodepool.containership.io/empty-value",
			value:       "",
			expected:    "nodepool.containership.io/empty-value=",
			shouldError: false,
		},
	}

	for _, test := range tests {
		result, err := String(test.key, test.value)

		if test.shouldError {
			assert.Error(t, err)
			continue
		}

		assert.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}

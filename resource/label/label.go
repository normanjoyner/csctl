package label

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	// NodePoolPrefix is the prefix required for any node pool labels
	NodePoolPrefix = "nodepool.containership.io"
	// ClusterPrefix is the prefix required for any cluster labels
	ClusterPrefix = "cluster.containership.io"
)

// ParseString parses a string of the form key=value and returns
// the key and value or an error.
func ParseString(labelStr string) (string, string, error) {
	fields := strings.Split(labelStr, "=")

	if len(fields) != 2 {
		return "", "", errors.Errorf("malformed label string: expected 2 fields, found %d", len(fields))
	}

	key, val := fields[0], fields[1]

	if key == "" {
		return "", "", errors.New("key cannot be empty")
	}

	return key, val, nil
}

// String returns a string representation of a label with the given
// key and value or an error.
func String(key string, value string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	return fmt.Sprintf("%s=%s", key, value), nil
}

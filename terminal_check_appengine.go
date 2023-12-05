//go:build appengine
// +build appengine

package logrusy

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}

//go:build js || nacl || plan9
// +build js nacl plan9

package logrusy

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}

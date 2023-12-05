//go:build js
// +build js

package logrusy

func isTerminal(fd int) bool {
	return false
}

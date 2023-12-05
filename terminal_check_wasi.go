//go:build wasi
// +build wasi

package logrusy

func isTerminal(fd int) bool {
	return false
}

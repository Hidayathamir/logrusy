//go:build (darwin || dragonfly || freebsd || netbsd || openbsd || hurd) && !js
// +build darwin dragonfly freebsd netbsd openbsd hurd
// +build !js

package logrusy

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}

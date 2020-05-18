// +build !windows

package cmd

import (
	"io"
	"path"
	"syscall"
)

const eolStr = "\n"

// enableVirtualTerminalProcessing does nothing.
func enableVirtualTerminalProcessing(w io.Writer) error {
	return nil
}

func getUmask() int {
	umask := syscall.Umask(0)
	syscall.Umask(umask)
	return umask
}

// makeCleanAbsSlashPath returns a clean, absolute path separated with forward
// slashes. If file is not an absolute path then it is joined on to dir.
func makeCleanAbsSlashPath(dir, file string) string {
	if !path.IsAbs(file) {
		file = path.Join(dir, file)
	}
	return path.Clean(file)
}

func trimExecutableSuffix(s string) string {
	return s
}

//go:build !windows
// +build !windows

package fs

import (
	"os"
	"syscall"
)

// fixpath returns an absolute path on windows, so restic can open long file
// names.
func fixpath(name string) string {
	return name
}

// TempFile creates a temporary file which has already been deleted (on
// supported platforms)
func TempFile(dir, prefix string) (f *os.File, err error) {
	f, err = os.CreateTemp(dir, prefix)
	if err != nil {
		return nil, err
	}

	if err = os.Remove(f.Name()); err != nil {
		return nil, err
	}

	return f, nil
}

// isNotSuported returns true if the error is caused by an unsupported file system feature.
func isNotSupported(err error) bool {
	if perr, ok := err.(*os.PathError); ok && perr.Err == syscall.ENOTSUP {
		return true
	}
	return false
}

// Chmod changes the mode of the named file to mode.
func Chmod(name string, mode os.FileMode) error {
	err := os.Chmod(fixpath(name), mode)

	// ignore the error if the FS does not support setting this mode (e.g. CIFS with gvfs on Linux)
	if err != nil && isNotSupported(err) {
		return nil
	}

	return err
}

// IsMainFile specifies if this is the main file or a secondary attached file (like ADS in case of windows)
func IsMainFile(name string) bool {
	// no-op - In case of OS other than windows this is always true
	return true
}

// SanitizeMainFileName will only keep the main file and remove the secondary file.
func SanitizeMainFileName(str string) string {
	// no-op - In case of non-windows there is no secondary file
	return str
}

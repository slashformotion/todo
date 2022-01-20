package internal

import (
	"github.com/spf13/afero"
)

var Fs = afero.NewOsFs()

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(path string) (bool, error) {
	exists, err := afero.Exists(Fs, path)
	if err != nil {
		return false, err
	}
	isDir, err := afero.IsDir(Fs, path)
	if err != nil {
		return false, err
	}
	return exists && isDir, nil
}

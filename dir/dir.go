package dir

// ref: https://github.com/k0sproject/k0s/blob/main/internal/pkg/dir/dir.go

import (
	"fmt"
	"os"
	"os/exec"
)

// IsDir check the given path exists and is a directory.
func IsDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.Mode().IsDir()
}

// GetAll returns a list of dirs in given base path.
func GetAll(base string) ([]string, error) {
	var dirs []string
	if !IsDir(base) {
		return dirs, fmt.Errorf("%s is not a directory", base)
	}

	// ReadDir reads the named directory,
	// returning all its directory entries sorted by filename.
	// If an error occurs reading the directory,
	// ReadDir returns the entries it was able to read before the error,
	// along with the error.
	dirList, err := os.ReadDir(base)
	for _, d := range dirList {
		if d.IsDir() {
			dirs = append(dirs, d.Name())
		}
	}

	if err != nil {
		return dirs, err
	}
	return dirs, nil
}

// Copy simply execute shell command "cp -r src dst".
func Copy(dst, src string) error {
	cmd := exec.Command("cp", "-r", src, dst)
	return cmd.Run()
}

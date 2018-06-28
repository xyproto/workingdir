package workingdir

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Dir struct {
	path string
}

// Create a new "working directory" struct
func New(path string) (*Dir, error) {
	// Check if we can enter the given path
	err := os.Chdir(path)
	if err != nil {
		return nil, err
	}
	// Get the full path
	fullpath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// Return a new *Dir
	return &Dir{fullpath}, nil
}

// Ensure that we are in the d.path directory
func (d *Dir) ensure() error {
	fullpath, err := os.Getwd()
	if err != nil {
		return err
	}
	if d.path == fullpath {
		return nil
	}
	return os.Chdir(d.path)
}

// List the contents of the directory
func (d *Dir) List() []string {
	var s []string
	if err := d.ensure(); err != nil {
		// Return an empty list if we can't enter the correct directory
		return s
	}
	files, _ := ioutil.ReadDir(d.path)
	for _, f := range files {
		s = append(s, f.Name())
	}
	return s
}

// Path returns the path of the directory
func (d *Dir) Path() string {
	return d.path
}

// Execute a command in the current directory.
// Return the combined output, or an empty string.
func (d *Dir) Run(cmd string) string {
	var cmdo *exec.Cmd
	if err := d.ensure(); err != nil {
		// Return an empty string if we can't enter the correct directory
		return ""
	}
	if strings.Contains(cmd, " ") {
		fields := strings.Split(cmd, " ")
		cmdo = exec.Command(fields[0], fields[1:]...)
	} else {
		cmdo = exec.Command(cmd)
	}
	b, err := cmdo.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(b)
}

// Execute a command in the current directory.
// Return the combined output, trimmed, or an empty string.
func (d *Dir) TrimRun(cmd string) string {
	return strings.TrimSpace(d.Run(cmd))
}

// Execute a command in the current directory.
// Return the combined output or an empty string.
func Run(cmd string) string {
	d, err := New(".")
	if err != nil {
		// Return an empty string if we can't enter the correct directory
		return ""
	}
	return d.Run(cmd)
}

// Execute a command in the current directory.
// Return the combined output, trimmed, or an empty string.
func TrimRun(cmd string) string {
	return strings.TrimSpace(Run(cmd))
}

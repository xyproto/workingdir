package workingdir

import (
	"io/ioutil"
	"os"
)

var curdir = "."

type Dir struct {
	path string
}

// Create a new "working directory" struct
func New(path string) (*Dir, error) {
	err := os.Chdir(path)
	if err != nil {
		return nil, err
	}
	return &Dir{path}, err
}

// List the contents of the directory
func (d *Dir) List() []string {
	var s []string
	files, _ := ioutil.ReadDir(d.path)
	for _, f := range files {
		s = append(s, f.Name())
	}
	return s
}

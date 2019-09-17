package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// File struct is general struct for all files/dirs
type File struct {
	ModifiedTime time.Time `json:"ModifiedTime"`
	IsLink       bool      `json:"IsLink"`
	IsDir        bool      `json:"IsDir"`
	LinksTo      string    `json:"LinksTo"`
	Size         int64     `json:"Size"`
	Name         string    `json:"Name"`
	Path         string    `json:"Path"`
	Children     []*File   `json:"Children"`
}

// IterateJSON takes a root path and walks the dirtree creating JSON representation
func IterateJSON(path string) string {
	rootOSFile, _ := os.Stat(path)

	//start with root file
	rootFile := toFile(rootOSFile, path)
	stack := []*File{rootFile}

	//until stack is empty,
	for len(stack) > 0 {
		//pop entry from stack
		file := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//get the children of entry
		children, _ := ioutil.ReadDir(file.Path)

		//for each child
		for _, chld := range children {
			//turn it into a File object
			child := toFile(chld, filepath.Join(file.Path, chld.Name()))

			//append it to the children of the current file popped
			file.Children = append(file.Children, child)

			//append the child to the stack, so the same process can be run again
			stack = append(stack, child)
		}
	}
	output, _ := json.MarshalIndent(rootFile, "", "     ")
	fmt.Println(string(output))
	return string(output)
}

func toFile(file os.FileInfo, path string) *File {
	JSONFile := File{ModifiedTime: file.ModTime(),
		IsDir:    file.IsDir(),
		Size:     file.Size(),
		Name:     file.Name(),
		Path:     path,
		Children: []*File{},
	}
	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		JSONFile.IsLink = true
		JSONFile.LinksTo, _ = filepath.EvalSymlinks(filepath.Join(path, file.Name()))
	}
	// Else case is the zero values of the fields
	return &JSONFile
}

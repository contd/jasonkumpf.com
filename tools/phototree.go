package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Directory struct is for directory objects
type Directory struct {
	Name     string     `json:"name"`
	Size     int64      `json:"size"`
	ModTime  time.Time  `json:"modified"`
	Path     string     `json:"path"`
	Thumb    string     `json:"thumb"`
	Items    int        `json:"items"`
	Pictures []*Picture `json:"pictures"`
}

// Exif picture data struct
type Exif struct {
	Lat   float64   `json:"lat"`
	Long  float64   `json:"long"`
	Taken time.Time `json:"taken"`
}

// Picture struct is for picture objects
type Picture struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	Type    string    `json:"type"`
	ModTime time.Time `json:"modified"`
	Path    string    `json:"path"`
	Thumb   string    `json:"thumb"`
	Width   int       `json:"width"`
	Height  int       `json:"height"`
	Exif    Exif      `json:"exif"`
}

// Files struct stores the current directory's contents in separate arrays of directores or pictures
type Files struct {
	Directories []Directory
	Pictures    []Picture
}

func main() {
	wd, e := os.Getwd()
	if e != nil {
		fmt.Printf("ERROR os.Getwd(): %v", e)
	}
	basePath := path.Join(wd, "client/static/photos/albums")

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("ERROR accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			if info.Name() != "thumbs" {
				// Here we create a Directory type/object then add to the global array
				fmt.Printf("Dir: %+v \n", info.Name())
			}
		} else {
			if !strings.Contains(path, "thumbs") {
				// Here we create a Directory type/object then add to the global array
				fmt.Printf("File: %+v  - Path: %q\n", info.Name(), path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("ERROR walking the path %q: %v\n", basePath, err)
		return
	}
}

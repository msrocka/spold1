package spold1

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFile writes the given EcoSpold data set to the file with the given path.
func WriteFile(es *EcoSpold, filePath string) error {
	data, err := xml.MarshalIndent(es, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, os.ModePerm)
}

// ReadFile reads an EcoSpold model from the given XML file.
func ReadFile(path string) (*EcoSpold, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var es EcoSpold
	if err := xml.Unmarshal(data, &es); err != nil {
		return nil, err
	}
	return &es, err
}

// EachFile processes each EcoSpold file in the given folder and calls the given
// function until there is no more file in the folder or the given function
// returns false.
func EachFile(dir string, fn func(spold *EcoSpold) bool) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		path := filepath.Join(dir, f.Name())
		spold, err := ReadFile(path)
		if err != nil {
			return err
		}
		if !fn(spold) {
			break
		}
	}
	return nil
}

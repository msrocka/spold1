package spold1

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// WriteFile writes the given EcoSpold data set to the file with the given path.
func WriteFile(es *EcoSpold, filePath string) error {
	data, err := xml.MarshalIndent(es, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, os.ModePerm)
}

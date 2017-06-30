// Package easycsv provides a quick wrapper for opening a csv file and returning
// the data in an easy to use, iterable way.
package easycsv

import (
	"encoding/csv"
	"os"
)

// Open takes the given path to a csv file and returns the rows as a [][]string which you can then
// iterate over.
func Open(path string) ([][]string, error) {
	fn, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer fn.Close()

	reader := csv.NewReader(fn)
	fields, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return fields, nil
}

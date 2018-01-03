// Package easycsv provides a quick wrapper for opening a csv file and returning
// the data in an easy to use, iterable way.
package easycsv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// CSV contains our buffer for writing to when creating a CSV file.
type CSV struct {
	Buffer *bufio.Writer
}

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

// NewCSV creates a new CSV file for writing to.
func NewCSV(name string) (*CSV, error) {
	csv, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	defer csv.Close()

	buf := bufio.NewWriter(csv)

	return &CSV{
		Buffer: buf,
	}, nil
}

// Write appends the given content to the newly created CSV file.
func (c *CSV) Write(content string) {
	fmt.Fprintf(c.Buffer, content)
}

// End finishes the writing to the newly created CSV file.
func (c *CSV) End() {
	c.Buffer.Flush()
}

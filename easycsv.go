// Package easycsv provides an easy to use wrapper for reading from and writing to CSV files.
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
	File   *os.File
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
	reader.Comment = '#'

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

	buf := bufio.NewWriter(csv)

	return &CSV{
		Buffer: buf,
		File:   csv,
	}, nil
}

// Write appends the given content to the newly created CSV file.
func (c *CSV) Write(content string) {
	fmt.Fprintf(c.Buffer, content)
}

// End finishes the writing to the newly created CSV file, and closes it.
func (c *CSV) End() {
	c.Buffer.Flush()
	c.File.Close()
}

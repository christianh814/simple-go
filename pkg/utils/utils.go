package utils

import (
	"bufio"
	"io"
	"os"
)

// readFile returns the contents of a file as a slice of string
func ReadFile(file string) ([]string, error) {
	// set up the return slice of string
	var s = []string{}

	// open the file, report any errors
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	// close the file when the function exists
	defer f.Close()

	// read the file line by line and convert it into a slice of string
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		//if we get to the end of the file, break out of this loop
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		// append contents into a slice of string
		s = append(s, string(line))

	}

	return s, nil
}

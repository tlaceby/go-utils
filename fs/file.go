package fs

import (
	"io"
	"os"
	"strings"
)

// Read a textfiles contents to a string
func ReadTextFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", nil
	}

	buffer, err := io.ReadAll(file)

	if err != nil {
		return "", nil
	}

	return string(buffer), nil
}

// Returns a slice of each line inside a file
func GetLines(filename string) ([]string, error) {
	contents, err := ReadTextFile(filename)

	if err != nil {
		return make([]string, 0), err
	}

	return strings.Split(contents, "\n"), nil
}

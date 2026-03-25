// Package input read args from stdin
package input

import (
	"bufio"
	"io"
)

func ReadInput(in io.Reader, nullDelimited bool) ([]string, error) {
	scanner := bufio.NewScanner(in)

	if nullDelimited {
		scanner.Split(scanNullTerminated)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	items := make([]string, 0)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func scanNullTerminated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i, b := range data {
		if b == 0 {
			return i + 1, data[:i], nil
		}
	}

	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}

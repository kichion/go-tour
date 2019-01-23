package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (a MyReader) Read(rb []byte) (l int, e error) {
	for l, e = 0, nil; l < len(rb); l++ {
		rb[l] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}

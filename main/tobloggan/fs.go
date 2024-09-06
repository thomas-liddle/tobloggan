package main

import "os"

type FSWriter struct{}

func (FSWriter) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}
func (FSWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

package utils

import (
	"io"
	"github.com/schollz/progressbar/v3"
	"os"
	"io/ioutil"
)

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

// creating a new file to show progress bar 
func ReadFileWithProgress(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, _ := file.Stat()
	barProgress := progressbar.DefaultBytes(
		info.Size(),
		"reading",
	)
	return io.ReadAll(io.TeeReader(file, barProgress))
}
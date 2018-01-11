package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/Users/shibayu36/development/go/src/github.com/mackerelio/mkr", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		fmt.Println(path)
		return nil
	})
}

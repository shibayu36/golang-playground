package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	dir, err := ioutil.TempDir("/tmp", "hoge-")
	fmt.Println(dir, err)
}

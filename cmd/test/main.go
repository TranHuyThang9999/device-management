package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "http://127.0.0.1:8080/manager/files/export/f57387fe-412e-4f29-800e-2223feedac02.png"
	fmt.Println(filepath.Base(path))
}

package savefiles

import (
	"device_management/common/log"
	"os"
)

func ListFilesInDirectory(path string) []string {
	dir, err := os.Open(path)
	if err != nil {
		log.Errorf(err, "failed to open directory: %w")
		return nil
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Error(err, "failed to read directory: %w")
		return nil
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames
}

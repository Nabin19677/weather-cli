package utils

import (
	"fmt"
	"os"
)

var TEMP_FOLDER = "weather-cli-temp"

func FileExistsT(filename string) bool {
	filePath := fmt.Sprintf("%s/%s/%s", os.TempDir(), TEMP_FOLDER, filename)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func CreateFolderT() {
	folderPath := fmt.Sprintf("%s/%s", os.TempDir(), TEMP_FOLDER)
	err := os.Mkdir(folderPath, 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func WriteToFileT(filename string, body []byte) error {
	filePath := fmt.Sprintf("%s/%s/%s", os.TempDir(), TEMP_FOLDER, filename)
	return os.WriteFile(filePath, body, 0777)
}

func ReadFileT(filename string) ([]byte, error) {
	filePath := fmt.Sprintf("%s/%s/%s", os.TempDir(), TEMP_FOLDER, filename)
	return os.ReadFile(filePath)
}

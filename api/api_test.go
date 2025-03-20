package api

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCsvFile(t *testing.T) {
	tmpDir := CreateTmpDir()
	names := GetFilesNames(tmpDir)
	RemoveAllFilesFromPath(tmpDir, names)

	createCsv(tmpDir, "")

	files := GetFilesNames(tmpDir)
	assert.Equal(t, 1, len(files), "should be one file in this folder "+tmpDir)

	RemoveAllFilesFromPath(tmpDir, files)
}

func CreateTmpDir() string {
	userHome, err := os.UserHomeDir()

	if err != nil {
		log.Println("could not find user home dir")
	}

	tmpDir, err := os.MkdirTemp(userHome, "tmp")

	if err != nil {
		log.Println("could not create dir temp", err)
	}
	return tmpDir
}

func GetFilesNames(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal("could not open path=" + path)
	}

	files, err := f.Readdirnames(0)
	return files
}

func RemoveAllFilesFromPath(path string, names []string) {
	for _, name := range names {
		os.Remove(path + "/" + name)
	}
}

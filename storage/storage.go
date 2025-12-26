package storage

import (
	"fmt"
	"log"
	"os"
)

var (
	VaultDir string
)

func Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory:", err)
		os.Exit(1)
	}
	VaultDir = fmt.Sprintf("%s/.gotion", homeDir)

	// Ensure directory exists
	err = os.MkdirAll(VaultDir, 0750)
	if err != nil {
		log.Fatal("Error creating vault directory:", err)
		os.Exit(1)
	}
}

func GetVaultDir() string {
	if VaultDir == "" {
		Init()
	}
	return VaultDir
}

func ListFiles() ([]os.DirEntry, error) {
	if VaultDir == "" {
		Init()
	}
	return os.ReadDir(VaultDir)
}

func DeleteNote(filename string) error {
	if VaultDir == "" {
		Init()
	}
	return os.Remove(fmt.Sprintf("%s/%s", VaultDir, filename))
}

func ExportNotes() (string, error) {
	if VaultDir == "" {
		Init()
	}

	files, err := ListFiles()
	if err != nil {
		return "", err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	exportDir := fmt.Sprintf("%s/gotion-notes", currentDir)

	err = os.MkdirAll(exportDir, 0750)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() {
			sourceFile := fmt.Sprintf("%s/%s", VaultDir, file.Name())
			destFile := fmt.Sprintf("%s/%s", exportDir, file.Name())

			content, err := os.ReadFile(sourceFile)
			if err != nil {
				return "", err
			}

			err = os.WriteFile(destFile, content, 0644)
			if err != nil {
				return "", err
			}
		}
	}

	return exportDir, nil
}

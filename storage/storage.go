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

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var folderStructure = map[string]interface{}{
	"app": []string{},
	"db": []string{
		"migrations",
	},
	"delivery": map[string]interface{}{
		"http": map[string]interface{}{
			"controller": []string{},
			"dto": []string{
				"request",
				"response",
			},
			"middleware": []string{},
			"router":     []string{},
		},
	},
	"domain": []string{
		"entity",
		"repository",
		"usecase",
	},
	"shared": []string{
		"config",
		"logger",
		"token",
		"utils",
	},
}

func createFolderStructure(root string, structure map[string]interface{}) error {
	for folder, subfolders := range structure {
		folderPath := filepath.Join(root, folder)
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			return err
		}

		switch subfolders := subfolders.(type) {
		case []string:
			for _, subfolder := range subfolders {
				subfolderPath := filepath.Join(folderPath, subfolder)
				if err := os.MkdirAll(subfolderPath, 0755); err != nil {
					return err
				}
			}
		case map[string]interface{}:
			if err := createFolderStructure(folderPath, subfolders); err != nil {
				return err
			}
		}

		fmt.Printf("Folder '%s' berhasil dibuat.\n", folderPath)
	}

	fmt.Println("Struktur folder berhasil dibuat.")
	return nil
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Terjadi kesalahan: %s\n", err)
		return
	}

	rootDirectory := filepath.Dir(exePath)

	err = createFolderStructure(rootDirectory, folderStructure)
	if err != nil {
		fmt.Printf("Terjadi kesalahan: %s\n", err)
	}
}

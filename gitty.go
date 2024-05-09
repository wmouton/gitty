package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: gitty <username>/<repository>")
		return
	}

	repoUrl := args[1]
	url := "https://github.com/" + repoUrl
	dest := "~/cloned/" + repoUrl

	// Expand home directory in destination path
	expandedDest, err := expandHomeDirectory(dest)
	if err != nil {
		fmt.Println("Failed to expand home directory:", err)
		return
	}

	// Check if destination directory already exists
	if _, err := os.Stat(expandedDest); !os.IsNotExist(err) {
		fmt.Println("Destination directory already exists:", expandedDest)
		return
	}

	cmd := exec.Command("git", "clone", url, expandedDest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to clone repository:", err)
		return
	}

	fmt.Println("Repository cloned successfully")
}

// expandHomeDirectory expands the "~" symbol in the path to the user's home directory.
func expandHomeDirectory(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, path[1:]), nil
}

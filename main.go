package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Error Log
func check(err error) {
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		os.Exit(1)
	}
}

// List all file
func listFiles(p string, f *[]string) error {
	return filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			*f = append(*f, path)
		}
		return nil
	})
}

// Copy File
func copyFile(p string, f []string) {
	fileCount := len(f)

	// Convert fileCount to string
	fileCountStr := strconv.Itoa(fileCount)

	for i, file := range f {
		// Convert current number to string
		currentNumStr := strconv.Itoa(i + 1)

		// Pad the current number with leading zeros
		paddedNum := strings.Repeat("0", len(fileCountStr)-len(currentNumStr)) + currentNumStr

		fmt.Printf("Loop iteration: %s %s \n", paddedNum, file)

		srcFile, err := os.Open(file)
		check(err)
		defer srcFile.Close()

		destFile, err := os.Create(p + "/" + paddedNum + ".txt") // creates if file doesn't exist
		check(err)
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
		check(err)

		err = destFile.Sync()
		check(err)
	}

}

func main() {
	definedPath := "./test-dir"
	newDesination := definedPath + " (combined)"
	var (
		fileList []string
	)

	os.MkdirAll(newDesination, 0755)

	err := listFiles(definedPath, &fileList)

	sort.Strings(fileList)

	copyFile(newDesination, fileList)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

}

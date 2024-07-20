package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/akamensky/argparse"
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

		// Sort the slice after populating it
		sort.Strings(*f)

		return nil
	})
}

// Copy File
func copyFile(p string, f []string) {
	fileCount := len(f)

	// Convert fileCount to string
	fileCountStr := strconv.Itoa(fileCount)

	for i, file := range f {
		// Get the file extension
		ext := filepath.Ext(file)

		// Convert current number to string
		currentNumStr := strconv.Itoa(i + 1)

		// Pad the current number with leading zeros
		paddedNum := strings.Repeat("0", len(fileCountStr)-len(currentNumStr)) + currentNumStr

		fmt.Printf("%s %s \n", paddedNum, file)

		srcFile, err := os.Open(file)
		check(err)
		defer srcFile.Close()

		destFile, err := os.Create(p + "/" + paddedNum + ext) // creates if file doesn't exist
		check(err)
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
		check(err)

		err = destFile.Sync()
		check(err)
	}

}

func main() {
	// Create new parser object
	parser := argparse.NewParser("args", "Combine all file in a specified path")

	// Create string flag
	s := parser.String("p", "string", &argparse.Options{Required: true, Help: "A path that you want to combine"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	definedPath := *s
	newDesination := definedPath + " (combined)"
	var (
		fileList []string
	)

	os.MkdirAll(newDesination, 0755)

	err = listFiles(definedPath, &fileList)

	copyFile(newDesination, fileList)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

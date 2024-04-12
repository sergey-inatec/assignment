package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

// Shred overwrites the given file with random data 3 times and then deletes the file.
func Shred(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	randomData := make([]byte, fileSize)
	for i := 0; i < 3; i++ {
		if _, err := rand.Read(randomData); err != nil {
			return err
		}
		if _, err := file.WriteAt(randomData, 0); err != nil {
			return err
		}
	}
	if err := file.Sync(); err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./shred <file1> [<file2> ...]")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		if err := Shred(arg); err != nil {
			fmt.Printf("Error shredding file %s: %v\n", arg, err)
		} else {
			fmt.Printf("File %s shredded successfully\n", arg)
		}
	}
}


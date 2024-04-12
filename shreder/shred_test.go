package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestShredEmptyFile tests shredding an empty file.
func TestShredEmptyFile(t *testing.T) {
	// Create a temporary empty file
	file, err := ioutil.TempFile("", "emptyfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Shred the file
	err = Shred(file.Name())
	if err != nil {
		t.Errorf("Shred failed: %v", err)
	}

	// Check if file still exists
	if _, err := os.Stat(file.Name()); !os.IsNotExist(err) {
		t.Error("File still exists after shredding")
	}
}

// TestShredNonExistentFile tests shredding a non-existent file.
func TestShredNonExistentFile(t *testing.T) {
	// Shred a non-existent file
	err := Shred("nonexistentfile.txt")
	if err == nil {
		t.Error("Expected error but got nil")
	} else {
		t.Logf("Expected error: %v", err)
	}
}

// TestShredSmallFile tests shredding a small file.
func TestShredSmallFile(t *testing.T) {
	// Create a temporary file with some data
	data := []byte("hello world")
	file, err := ioutil.TempFile("", "smallfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(data); err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Shred the file
	err = Shred(file.Name())
	if err != nil {
		t.Errorf("Shred failed: %v", err)
	}

	// Check if file still exists
	if _, err := os.Stat(file.Name()); !os.IsNotExist(err) {
		t.Error("File still exists after shredding")
	}
}

// TestShredFileInUse tests shredding a file that is open for reading by another process.
func TestShredFileInUse(t *testing.T) {
	// Create a temporary file
	file, err := ioutil.TempFile("", "fileinuse")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Open the file for reading (simulate another process)
	go func() {
		f, err := os.Open(file.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		// Wait for a while to simulate file being open
	}()

	// Shred the file
	err = Shred(file.Name())
	if err != nil {
		t.Errorf("Shred failed: %v", err)
	}

	// Check if file still exists
	if _, err := os.Stat(file.Name()); !os.IsNotExist(err) {
		t.Error("File still exists after shredding")
	}
}

// TestShredLargeFile tests shredding a large file.
func TestShredLargeFile(t *testing.T) {
	// Create a temporary large file
	data := make([]byte, 1024*1024) // 1MB
	file, err := ioutil.TempFile("", "largefile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(data); err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Shred the file
	err = Shred(file.Name())
	if err != nil {
		t.Errorf("Shred failed: %v", err)
	}

	// Check if file still exists
	if _, err := os.Stat(file.Name()); !os.IsNotExist(err) {
		t.Error("File still exists after shredding")
	}
}

func TestMain(m *testing.M) {
	// Run the tests
	exitCode := m.Run()

	// Exit with the same code as the test execution
	os.Exit(exitCode)
}


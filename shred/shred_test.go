package shred 

import (
   	"fmt"
	"testing"
	"os"
	"math/rand" // faster than "crypto/rand" so I use this to create random test files
)

// Used for creating random testfiles
const min = 8
const max = 1024

func TestFileDeleted(t *testing.T) {
	fmt.Println("TestFileDeleted started")
	testFileName := "test_file_1"
	// Create a test file
	f, error := os.Create(testFileName)
	if error != nil {
		t.Fatalf("Error creating random file: %v", error)
	}
	defer f.Close()

	fileSize := rand.Intn(max - min) + min
	data := make([]byte, fileSize)
    _, error = rand.Read(data)
    if error != nil{
    	t.Fatalf("Test failed: %v", error)
    }

    // Write random data to file
    _, error = f.Write(data)
	if error != nil {
		t.Fatalf("Test failed: %v", error)
	}

	// Sync file with disk
	error = f.Sync()
	if error != nil {
		t.Fatalf("Test failed: %v", error)
	}

	res := Shred(testFileName)
	// If not nil, there was an error
	if res != nil{
		t.Fatalf("Test failed: %v", res)
	}
	// Check if the file got deleted
	_, error = os.Stat(testFileName);
	if !os.IsNotExist(error) {
		t.Fatalf("The random file was not deleted!")
	}
}

func TestInvalidFilepath(t *testing.T) {
	fmt.Println("TestInvalidFilepath started")
	testFileName := "test_file_2"
	// Try to Shred a non existing file
	res := Shred(testFileName)
	// It should return IsNotExist error
	if !os.IsNotExist(res){
		t.Fatalf("Test failed: %v", res)
	}
}

func TestFolderpath(t *testing.T) {
	fmt.Println("TestFolderpath started")
	test_folder_name := "test_folder"
	// Create a test folder
	error := os.Mkdir(test_folder_name, os.ModePerm)
	if error != nil {
		t.Fatalf("Test failed: %v", error)
	}
	// Try to Shred a folder
	res := Shred(test_folder_name)
	// It shouldn't return nil because os.OpenFile() should fail 
	if res == nil {
		t.Fatalf("Test failed: %v", res)
	}

	// Check if the folder got deleted
	_, error = os.Stat(test_folder_name)
	if os.IsNotExist(error) {
		t.Fatalf("The folder was deleted: %v", error)
	}

	// Delete test folder
	error = os.Remove(test_folder_name)
	if error != nil {
		t.Fatalf("Error removing the new folder: %v", error)
	}
}

func TestZeroSizeFile(t *testing.T) {
	fmt.Println("TestZeroSizeFile started")
	testFileName := "test_file_3"
	// Create a test file
	f, error := os.Create(testFileName)
	if error != nil {
		t.Fatalf("Error creating random file: %v", error)
	}
	defer f.Close()

	// Sync file with disk
	error = f.Sync()
	if error != nil {
		t.Fatalf("Test failed: %v", error)
	}

	res := Shred(testFileName)
	// If not nil, there was an error
	if res != nil{
		t.Fatalf("Test failed: %v", res)
	}
	// Check if the file got deleted
	_, error = os.Stat(testFileName);
	if !os.IsNotExist(error) {
		t.Fatalf("The random file was not deleted!")
	}
}

func TestLargeFile(t *testing.T) {
	fmt.Println("TestLargeFile started for size: 0.5GB")
	testFileName := "test_file_4"
	// Create a test file
	f, error := os.Create(testFileName)
	if error != nil {
		t.Fatalf("Error creating random file: %v", error)
	}
	defer f.Close()

	fileSize := 536870912 // 0.5 GB file
	data := make([]byte, fileSize)
    _, error = rand.Read(data)
    if error != nil{
    	t.Fatalf("Test failed: %v", error)
    }

    // Write random data to file
    _, error = f.Write(data)
	if error != nil {
		t.Fatalf("Test failed: %v", error)
	}

	// Sync file with disk
	error = f.Sync()
	if error != nil {
		t.Fatalf("Test failed: %v", error)
	}

	res := Shred(testFileName)
	// If not nil, there was an error
	if res != nil{
		t.Fatalf("Test failed: %v", res)
	}
	
	// Check if the file got deleted
	_, error = os.Stat(testFileName);
	if !os.IsNotExist(error) {
		t.Fatalf("The random file was not deleted!")
	}
}
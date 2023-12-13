package shred 

import (
   	"fmt"
	"testing"
	"os"
	"math/rand"
)

const min = 8
const max = 1024

func TestFileDeleted(t *testing.T) {
	fmt.Println("TestFileDeleted started")
	testFileName := "test_file_1"
	// Create a random file
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
	res := Shred(testFileName)
	if !os.IsNotExist(res){
		t.Fatalf("Test failed: %v", res)
	}
}

func TestFolderpath(t *testing.T) {
	fmt.Println("TestFolderpath started")
	test_folder_name := "test_folder"
	error := os.Mkdir(test_folder_name, os.ModePerm)
	if error != nil {
		t.Fatalf("Test failed: %v", error)
	}
	res := Shred(test_folder_name)
	if res == nil {
		t.Fatalf("Test failed: %v", res)
	}

	_, error = os.Stat(test_folder_name)
	if os.IsNotExist(error) {
		t.Fatalf("The folder was deleted: %v", error)
	}

	error = os.Remove(test_folder_name)
	if error != nil {
		t.Fatalf("Error removing the new folder: %v", error)
	}
}

func TestZeroSizeFile(t *testing.T) {
	fmt.Println("TestZeroSizeFile started")
	testFileName := "test_file_3"
	// Create a random file
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
	// Create a random file
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
	if res != nil{
		t.Fatalf("Test failed: %v", res)
	}
	// Check if the file got deleted
	_, error = os.Stat(testFileName);
	if !os.IsNotExist(error) {
		t.Fatalf("The random file was not deleted!")
	}
}
package shred

import (
	// "fmt"
	"os"
	"crypto/rand" // better vs "math/rand" for cryptographic random generator (but slower)
)


func Shred(filepath string) error {

	// Check if 'filepath' is a valid Path that leads to an existing file
	fileInfo, error := os.Stat(filepath)
	if os.IsNotExist(error) {
		return error
	}
	// Get the size of the file
	fileSize := fileInfo.Size()
    
	// Try to open the file
	f, error := os.OpenFile(filepath, os.O_RDWR, 0644)
    if error != nil {
        return error
    }
    defer f.Close() // close the file if any other error occurs

	// 3 iterations
	for i:=0; i<3; i++ {

		// Move at the beginning of the file
	    _, error = f.Seek(0, 0)
		if error != nil {
			return error
		}

	    // Generate a random number for the file size
	    data := make([]byte, fileSize)
	    _, error = rand.Read(data)
	    if error != nil{
	    	return error
	    }

	    // Write random data to file
	    _, error = f.Write(data)
		if error != nil {
			return error
		}

		// Sync file with disk
		error = f.Sync()
		if error != nil {
			return error
		}
	}

	// Close and delete the file
	f.Close()
	error = os.Remove(filepath)
	if error != nil {
		return error
	}

	// Success
	return nil
}


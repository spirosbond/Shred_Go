# Shred tool in GO

This repo implements an example of a Shred function. It comes with a simple main function to help calling the Shred function as well as test cases.

The Shred function takes a file path as an input and overwrites it fully 3 times, before deleting it. This ensures the data of the file cannot be recovered by scanning the raw bytes of the disk and reassembling them ([File Carving](https://en.wikipedia.org/wiki/File_carving))


## Possible Use Cases

This technique can be useful for deleting confidential files and making it impossible to recover them with [File Carving](https://en.wikipedia.org/wiki/File_carving). By writing 3 times random bytes on the same memory space, it is almost guaranteed that there will be no way to reconstruct the original information.

Possible use cases:

- Secure file deletion of sensitive data
- Data Sanitization before releasing resources (selling or disposing devices)
- User Privacy Protection to ensure compliance with data polocies

## Advantages / Disadvantages of this approach

✔️ Simple to create and maintain

✔️ Uses the "crypto/rand" random number generator which is more secure than the "math/rand"

✔️ 3 overwrites make very difficult to recover the deleted file

✖️ Can be slow for larger files

✖️ Can create wear on the disk due to the multiple writes

✖️ This implementation is not using any type of parallelism/multithreading. We could assign different memory segments of the file to different threads to speed it up

## Tested on
Manjaro 6.1.67-2-MANJARO (64-bit) with Go 2:1.21.5-1 package

## How to Execute

- Clone this repository
- Navigate to the root of the repository
- Run the `main.go`:
```bash
go run main.go <path_to_file_you_want_to_shred>
```

## How to Execute Tests

- Navigate into the folder `shred`
- Run:
```bash
go test
```

## Implemented Tests

- Deletion of Random file
- Confirm no crash if wrong path is given
- Confirm no crash if folder is given instead of a file
- Deletion of zero byte file
- Deletion of large file (0.5GB)

## Further Info

This implementation is done by using modules:

```
- spirosbond/main
	|
	--> spirosbond/shred
```

The following command has been used to enable running `main.go` without publishing the modules:

```bash
go mod edit -replace spirosbond/shred=./shred
```

## Useful Sources
- [Go Dev Docs](https://go.dev/doc/tutorial/)
- [Go by example](https://gobyexample.com/)
- [Go Samples](https://gosamples.dev/)
- [File Carving - Wikipedia](https://en.wikipedia.org/wiki/File_carving)
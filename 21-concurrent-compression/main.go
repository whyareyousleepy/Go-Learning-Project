package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open file %s, %s\n", filename, err)
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		fmt.Printf("Unable to create file %s, %s\n", filename+".gz", err)
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}

func main() {

	var wg sync.WaitGroup

	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			compress(file)
			wg.Done()
		}(file)
	}

	wg.Wait()
	fmt.Printf("Compressed %d files\n", len(os.Args[1:]))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	rootDir := "C:\\Users\\TO11RC\\OneDrive - ING\\Downloads\\"
	pattern := "TIBCO_SECURITY_VENDOR"

	var wg sync.WaitGroup

	callback := func(name string, info os.FileInfo, err error) error {
		wg.Add(1)

		go func() {
			file := grep(name, pattern)
			if file != "NOT_FOUND" {
				//file = strings.Replace(file, rootDir, "", -1)
				fmt.Println(file)
			}
			defer wg.Done()
		}()

		return nil
	}

	filepath.Walk(rootDir, callback)

	wg.Wait()

	diff := time.Since(start)
	fmt.Println("duration: " + diff.String())
}

func grep(filename string, pattern string) string {
	file, _ := os.Open(filename)

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	fileName := "NOT_FOUND"
	for Scanner.Scan() {
		word := Scanner.Text()
		if word == pattern {
			fileName = file.Name()
			break
		}
	}
	return fileName
}

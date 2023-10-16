package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestWithGrep(t *testing.T) {
	file, err := os.Open("C:/Users/TO11RC/OneDrive - ING/Downloads/job_10359060.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileName := "not found"

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		word := Scanner.Text()

		if word == "TASK" {
			//fileName = path.Base(file.Name())
			fileName = file.Name()
			break
		}
	}
	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileName)
}

func TestWithMain(t *testing.T) {
	file := grep_test("C:/Users/TO11RC/OneDrive - ING/Downloads/job_10359060.txt", "TASK")
	fmt.Println(file)
}

func grep_test(filename string, pattern string) string {
	file, _ := os.Open(filename)

	fileName := "not found"

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		word := Scanner.Text()
		if word == pattern {
			fileName = file.Name()
			break
		}
	}
	return fileName
}

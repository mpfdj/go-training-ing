package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name string
	sum  int
}

func main() {
	dat, _ := os.ReadFile("C:/Users/TO11RC/OneDrive - ING/miel/workspace/go/go-training-ing/exercise_no_space_left_on_device/input.txt")
	lines := strings.Split(string(dat), "\r\n")

	var stack []dir
	var sum = 0

	for index, line := range lines {

		if strings.HasPrefix(line, "$ cd") {

			if strings.HasPrefix(line, "$ cd ..") {
				dir := stack[len(stack)-1]

				// pop
				n := len(stack) - 1
				stack = stack[:n]

				stack[len(stack)-1].sum = stack[len(stack)-1].sum + dir.sum

			} else {

				for i := index + 1; i < len(lines); i++ {
					lineInner := lines[i]

					if isChangeDirectory(lineInner) {
						break
					}

					if !strings.HasPrefix(lineInner, "$") && !strings.HasPrefix(lineInner, "dir") {
						size := strings.Split(lineInner, " ")[0]
						sizeInt, _ := strconv.Atoi(size)
						sum = sum + sizeInt
					}

				}

				// push
				name := strings.Split(line, " ")[2]
				dir := dir{name, sum}

				stack = append(stack, dir)
				sum = 0

			} // eof else

			printStack(stack)
			fmt.Println()
		} // eof if

	} //eof for

	// Calculate total for root directory
	if len(stack) == 2 {
		dir := stack[len(stack)-1]

		// pop
		n := len(stack) - 1
		stack = stack[:n]

		stack[len(stack)-1].sum = stack[len(stack)-1].sum + dir.sum
		printStack(stack)
	}

}

func isChangeDirectory(s string) bool {
	if strings.HasPrefix(s, "$ cd") {
		if strings.HasPrefix(s, "$ cd ..") {
			return false
		} else {
			return true
		}
	}
	return false
}

func printStack(stack []dir) {
	var s = ""
	for _, value := range stack {
		if value.name == "/" {
			s = s + value.name
		} else {
			s = s + value.name + "/"
		}
		s = s + "(" + strconv.Itoa(value.sum) + ")"
	}
	fmt.Print(s)
}

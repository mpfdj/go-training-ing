package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type State struct {
	words []string
	count int
}

type LogProcessor struct {
	states map[string]*State
}

//var re = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (INFO|WARNING|NOTICE) (.+)$`)

func main() {
	runProcessor()

}
func runProcessor() {
	lp := NewLogProcessor()

	content, err := os.ReadFile("C:/Users/TO11RC/OneDrive - ING/miel/workspace/go/go-training-ing/training-concurrency-3/log.txt")
	if err != nil {
		panic(err)
	}

	lp.processLogFile(content)
	lp.outputSummary()
}

func NewLogProcessor() *LogProcessor {
	return &LogProcessor{
		states: make(map[string]*State),
	}
}

func (lp *LogProcessor) processLogFile(content []byte) {
	lines := bytes.Split(content, []byte("\n"))

	for _, line := range lines {
		lp.processLogEntry(string(line))
	}
}

func (lp *LogProcessor) processLogEntry(line string) {
	re := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (INFO|WARNING|NOTICE) (.+)$`)
	matches := re.FindStringSubmatch(line)

	if len(matches) == 0 {
		return
	}

	words := strings.Split(matches[3], " ")

	for _, word := range words {
		lp.processWord(matches[2], word)
	}
}

func (lp *LogProcessor) processWord(state, word string) {
	if lp.states[state] == nil {
		lp.states[state] = &State{
			words: []string{},
			count: 0,
		}
	}

	if lp.containsWord(state, word) {
		return
	}

	entry := *lp.states[state]
	entry.count++
	entry.words = append(entry.words, word)

	lp.states[state] = &entry
}

func (lp *LogProcessor) containsWord(state, word string) bool {
	for _, existingWord := range lp.states[state].words {
		if existingWord == word {
			return true
		}
	}

	return false
}

func (lp *LogProcessor) outputSummary() {
	var output string

	for name, state := range lp.states {
		if name == "notice" {
			continue
		}

		output += fmt.Sprintf("%s: %d\n", name, state.count)
	}

	fmt.Println(output)
}

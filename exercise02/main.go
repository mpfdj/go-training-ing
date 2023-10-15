package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type vCard struct {
	N     string
	FN    string
	TEL   string
	EMAIL string
}

var data = `BEGIN:VCARD
FN:John Doe
TEL:1234567890
EMAIL:john.doe@example.com
END:VCARD

BEGIN:VCARD
FN:Jane Smith
TEL:9876543210
EMAIL:jane.smith@outlook.com
END:VCARD

BEGIN:VCARD
FN:Bob Johnson
TEL:1112233445
EMAIL:bob.johnson@live.com
END:VCARD

BEGIN:VCARD
FN:Alice Roberts
TEL:6667788990
EMAIL:alice.roberts@gmail.com
END:VCARD

BEGIN:VCARD
FN:Charlie Brown
TEL:5554443322
EMAIL:charlie.brown@msn.nl
END:VCARD

BEGIN:VCARD
FN:Emily Davis
TEL:3332221111
EMAIL:emily.davis@ing.com
END:VCARD

BEGIN:VCARD
FN:Frank Martin
TEL:7778889999
EMAIL:frank.martin@example.com
END:VCARD

BEGIN:VCARD
FN:Grace Lee
TEL:2223334444
EMAIL:grace.lee@example.com
END:VCARD`

func main() {
	list := parseVcardString(data)
	vcards := parseVcardList(list)
	printVcardList(vcards)
}

func parseVcardString(s string) []string {
	var list []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			list = append(list, line)
		}
	}
	return list
}

func parseVcardList(list []string) []vCard {
	var vcards []vCard
	for i := 0; i < len(list); i++ {
		if i%5 == 0 {
			fn := strings.Split(list[i+1], ":")[1]
			tel := strings.Split(list[i+2], ":")[1]
			email := strings.Split(list[i+3], ":")[1]

			arr := strings.Split(fn, " ")
			n := arr[1] + ";" + arr[0]

			vcard := vCard{N: n, FN: fn, TEL: tel, EMAIL: email}
			vcards = append(vcards, vcard)
		}
	}
	return vcards
}

func printVcardList(vcards []vCard) {
	sort.Slice(vcards, func(i, j int) bool {
		return vcards[i].N < vcards[j].N
	})

	for i, v := range vcards {
		fmt.Printf("%d = %v\n", i, v)
	}
}

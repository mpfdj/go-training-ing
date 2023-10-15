package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testdata = `BEGIN:VCARD
FN:John Doe
TEL:1234567890
EMAIL:john.doe@example.com
END:VCARD

BEGIN:VCARD
FN:Jane Smith
TEL:9876543210
EMAIL:jane.smith@outlook.com
END:VCARD`

func TestWithVcards(t *testing.T) {
	list := parseVcardString(testdata)
	vcards := parseVcardList(list)

	fmt.Printf("%#v\n", list)
	fmt.Printf("%#v\n", vcards)

	assert.Equal(t, 10, len(list))
	assert.Equal(t, 2, len(vcards))

	jd := vcards[0]
	assert.Equal(t, "Doe;John", jd.N)
	assert.Equal(t, "John Doe", jd.FN)
	assert.Equal(t, "1234567890", jd.TEL)
	assert.Equal(t, "john.doe@example.com", jd.EMAIL)

	js := vcards[1]
	assert.Equal(t, "Smith;Jane", js.N)
	assert.Equal(t, "Jane Smith", js.FN)
	assert.Equal(t, "9876543210", js.TEL)
	assert.Equal(t, "jane.smith@outlook.com", js.EMAIL)
}

func TestSomething(t *testing.T) {
	var a string = "Hello"
	var b string = "Helloxxx"
	assert.Equal(t, a, b, "The two words should be the same.")
}

package test_url

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestParseURL(t *testing.T) {

	//s := "/"
	//s := "/?hello=world&name=miel"
	s := "/DOC-EXAMPLE-BUCKET/]"

	result := strings.Split(s, "/")

	fmt.Println("len: " + strconv.Itoa(len(result)))
	fmt.Println("element 0: " + result[0])
	fmt.Println("element 1: " + result[1])

	//if len(s) > 2 {
	//	params, _ := url.ParseQuery(s[2:])
	//	fmt.Println(params)
	//} else {
	//	fmt.Println("No params found")
	//}

	//for _, p := range getUrlQueryParams(s) {
	//	p.
	//}

	//values := r.URL.Query()
	//for k, v := range values {
	//	fmt.Println(k, " => ", v)
	//}

}

func getUrlQueryParams(urlString string) url.Values {
	if len(urlString) > 2 {
		params, _ := url.ParseQuery(urlString[2:])
		return params
	}
	return nil
}

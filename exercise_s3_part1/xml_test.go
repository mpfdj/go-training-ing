package main

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var data = `<?xml version="1.0" encoding="UTF-8"?>
<ListAllMyBucketsResult>
	<Name>test</Name>
</ListAllMyBucketsResult>`

type BucketList struct {
	XMLName xml.Name `xml:"ListAllMyBucketsResult"`
	Name    string   `xml:"Name"`
}

type Bucket struct {
	XMLName      xml.Name `xml:"Bucket"`
	CreationDate string   `xml:"CreationDate"`
	Name         string   `xml:"Name"`
}

func TestMarshalXML(t *testing.T) {

	var bucketList BucketList
	if err := xml.Unmarshal([]byte(data), &bucketList); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", bucketList)

}

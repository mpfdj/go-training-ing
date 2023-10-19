package tests_unmarshal

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// XML to Go
// https://jsonformatter.org/xml-to-go?utm_content=cmp-true

var data = `<?xml version="1.0" encoding="UTF-8"?>
<ListAllMyBucketsResult>
   <Buckets>
      <Bucket>
         <CreationDate>timestamp</CreationDate>
         <Name>string</Name>
      </Bucket>
   </Buckets>
   <Owner>
      <DisplayName>string</DisplayName>
      <ID>string</ID>
   </Owner>
</ListAllMyBucketsResult>`

func (b BucketList) String() string {
	return fmt.Sprintf("CreationDate=%v", b.Buckets[0].Bucket.CreationDate)
}

type BucketList struct {
	Buckets []Buckets `xml:"Buckets"`
	Owner   Owner     `xml:"Owner"`
}

type Buckets struct {
	Bucket Bucket `xml:"Bucket"`
}

type Bucket struct {
	CreationDate string `xml:"CreationDate"`
	Name         string `xml:"Name"`
}

type Owner struct {
	DisplayName string `xml:"DisplayName"`
	ID          string `xml:"ID"`
}

func TestUnmarshalXML(t *testing.T) {

	var bucketList BucketList
	if err := xml.Unmarshal([]byte(data), &bucketList); err != nil {
		panic(err)
	}
	fmt.Println(bucketList)
	fmt.Printf("%#v\n", bucketList)

}

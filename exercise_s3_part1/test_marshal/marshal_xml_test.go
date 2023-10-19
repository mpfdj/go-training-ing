package tests_unmarshall

import (
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

type ListAllMyBucketsResult struct {
	Buckets []Bucket
	Owner   Owner
}

//type Buckets struct {
//	Bucket Bucket `xml:"Bucket"`
//}

type Bucket struct {
	CreationDate string
	Name         string
}

type Owner struct {
	DisplayName string
	ID          string
}

func TestMarshalXML(t *testing.T) {
	var buckets []Bucket
	bucket1 := Bucket{"creationDate1", "name1"}
	bucket2 := Bucket{"creationDate2", "name2"}

	buckets = append(buckets, bucket1)
	buckets = append(buckets, bucket2)

	owner := Owner{"displayName", "id"}

	bucketList := ListAllMyBucketsResult{Buckets: buckets, Owner: owner}

	fmt.Printf("%#v\n", bucketList)

}

package main

import (
	"encoding/xml"
	"net/http"
)

var data = `<?xml version="1.0" encoding="UTF-8"?>
<ListAllMyBucketsResult>
   <Buckets>
      <Bucket>
         <CreationDate>2019-12-11T23:32:47+00:00</CreationDate>
         <Name>DOC-EXAMPLE-BUCKET</Name>
      </Bucket>
      <Bucket>
         <CreationDate>2019-11-10T23:32:13+00:00</CreationDate>
         <Name>DOC-EXAMPLE-BUCKET2</Name>
      </Bucket>
   </Buckets>
   <Owner>
      <DisplayName>Miel de Jaeger</DisplayName>
      <ID>AIDACKCEVSQ6C2EXAMPLE</ID>
   </Owner>
</ListAllMyBucketsResult>`

type ListAllMyBucketsResult struct {
	Buckets Buckets `json:"Buckets"`
	Owner   Owner   `json:"Owner"`
}

type Buckets struct {
	Bucket []Bucket `json:"Bucket"`
}

type Bucket struct {
	CreationDate string `json:"CreationDate"`
	Name         string `json:"Name"`
}

type Owner struct {
	DisplayName string `json:"DisplayName"`
	ID          string `json:"ID"`
}

func list(w http.ResponseWriter, req *http.Request) {
	bucket1 := Bucket{"2019-12-11T23:32:47+00:00", "DOC-EXAMPLE-BUCKET"}
	bucket2 := Bucket{"2019-11-10T23:32:13+00:00", "DOC-EXAMPLE-BUCKET2"}

	var buckets []Bucket
	buckets = append(buckets, bucket1)
	buckets = append(buckets, bucket2)

	b := Buckets{Bucket: buckets}

	owner := Owner{"Miel de Jaeger", "AIDACKCEVSQ6C2EXAMPLE"}

	bucketList := ListAllMyBucketsResult{Buckets: b, Owner: owner}

	x, _ := xml.MarshalIndent(bucketList, "", "  ")

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)

}

func main() {
	http.HandleFunc("/", list)
	http.ListenAndServe(":7000", nil)

}

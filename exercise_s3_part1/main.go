package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
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

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("url: " + req.URL.String())
	fmt.Println("url path: " + req.URL.Path)

	w.Header().Set("Content-Type", "application/xml")

	url := req.URL.String()
	path := strings.Split(req.URL.Path, "/")

	if len(path) == 3 {
		bucket := path[1]
		if bucket == "DOC-EXAMPLE-BUCKET" &&
			strings.Contains(url, "delimiter=") &&
			strings.Contains(url, "encoding-type=") &&
			strings.Contains(url, "prefix=") &&
			strings.Contains(url, "versions=") {

			fmt.Println("Listing bucket DOC-EXAMPLE-BUCKET...")
			fmt.Println()

			http.ServeFile(w, req, "C:/Users/TO11RC/OneDrive - ING/miel/workspace/go/go-training-ing/exercise_s3_part1/s3/bucket_mock.xml")
		} else {
			fmt.Println()
			http.ServeFile(w, req, "C:/Users/TO11RC/OneDrive - ING/miel/workspace/go/go-training-ing/exercise_s3_part1/s3/bucket_empty_mock.xml")
		}
	} else {
		resp := listAllBucketsMock()
		w.Write(resp)
	}

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7001", nil)

}

func listAllBucketsMock() []byte {
	bucket1 := Bucket{"2019-12-11T23:32:47+00:00", "DOC-EXAMPLE-BUCKET"}
	bucket2 := Bucket{"2019-11-10T23:32:13+00:00", "DOC-EXAMPLE-BUCKET2"}

	var buckets []Bucket
	buckets = append(buckets, bucket1)
	buckets = append(buckets, bucket2)

	b := Buckets{Bucket: buckets}

	owner := Owner{"Miel de Jaeger", "AIDACKCEVSQ6C2EXAMPLE"}

	bucketList := ListAllMyBucketsResult{Buckets: b, Owner: owner}

	x, _ := xml.MarshalIndent(bucketList, "", "  ")
	return x
}

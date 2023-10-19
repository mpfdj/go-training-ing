https://brainhive.notion.site
https://training.brainhive.nl

# Advent of code 2022
adventofcode.com/2022/day/7
https://github.com/anolivei/advent_of_code_2022
https://nickymeuleman.netlify.app/garden/aoc2022-day07

# vCard specification
https://datatracker.ietf.org/doc/html/rfc6350
https://datatracker.ietf.org/doc/html/rfc6350#section-6.2.2

# Init a go module
cd exercise02
go mod init go-training/exercise02
go mod tidy
go get -u github.com/stretchr/testify/assert


# AWS S3 ListBuckets
https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html


# MinIO
https://min.io/docs/minio/linux/reference/minio-mc-admin.html?ref=docs#installation
https://min.io/docs/minio/linux/reference/minio-mc.html#id3
mc alias set localS3 http://localhost:7000
mc alias list
mc ls localS3


# XML to Go
https://jsonformatter.org/xml-to-go?utm_content=cmp-true


https://github.com/dmeijboom/training-connect-go


cd exercise_grpc_todo
go mod init go-training/exercise_grpc_todo

go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
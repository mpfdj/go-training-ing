curl -s --header "Content-Type: application/json" --data "{}" http://localhost:8080/todo.v1.TodoService/ListTasks | jq
curl -s --header "Content-Type: application/json" --data @new_task.json http://localhost:8080/todo.v1.TodoService/CreateTask | jq


# IntelliJ plugin gopilot



https://github.com/dmeijboom/training-connect-go

go mod tidy

go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

cd training-connect-go
training-connect-go

buf generate
cd training-connect-go/gen


go run ./cmd/server

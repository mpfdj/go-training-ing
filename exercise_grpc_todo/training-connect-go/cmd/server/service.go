package main

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"

	todov1 "github.com/dillendev/training-connect-go/gen/todo/v1"
)

type TodoService struct{}

func (s *TodoService) ListTasks(ctx context.Context, req *connect.Request[todov1.ListTasksRequest]) (*connect.Response[todov1.ListTasksResponse], error) {

	tasklist := []*todov1.ListTasksResponse_Task{}

	for _, t := range Tasks {
		task := todov1.ListTasksResponse_Task{Id: t.id, Name: t.name}
		tasklist = append(tasklist, &task)
	}

	return connect.NewResponse(&todov1.ListTasksResponse{Tasks: tasklist}), nil

	//return connect.NewResponse(&todov1.ListTasksResponse{Tasks: []*todov1.ListTasksResponse_Task{
	//	{
	//		Id:   1,
	//		Name: "Buy milk",
	//	},
	//}}), nil
}

func (s *TodoService) CreateTask(ctx context.Context, req *connect.Request[todov1.CreateTaskRequest]) (*connect.Response[todov1.CreateTaskResponse], error) {

	task := Task{id: req.Msg.Id, name: req.Msg.Name}
	Tasks = append(Tasks, task)

	fmt.Printf("id=%d name=%s\n", req.Msg.Id, req.Msg.Name)
	return connect.NewResponse(&todov1.CreateTaskResponse{Status: "OK"}), nil
}

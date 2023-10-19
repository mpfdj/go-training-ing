// Code generated by protoc-gen-connect-go.exe. DO NOT EDIT.
//
// Source: todo/v1/todo.proto

package todov1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/dillendev/training-connect-go/gen/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TodoServiceName is the fully-qualified name of the TodoService service.
	TodoServiceName = "todo.v1.TodoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TodoServiceCreateTaskProcedure is the fully-qualified name of the TodoService's CreateTask RPC.
	TodoServiceCreateTaskProcedure = "/todo.v1.TodoService/CreateTask"
	// TodoServiceListTasksProcedure is the fully-qualified name of the TodoService's ListTasks RPC.
	TodoServiceListTasksProcedure = "/todo.v1.TodoService/ListTasks"
)

// TodoServiceClient is a client for the todo.v1.TodoService service.
type TodoServiceClient interface {
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
}

// NewTodoServiceClient constructs a client for the todo.v1.TodoService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTodoServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TodoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &todoServiceClient{
		createTask: connect_go.NewClient[v1.CreateTaskRequest, v1.CreateTaskResponse](
			httpClient,
			baseURL+TodoServiceCreateTaskProcedure,
			opts...,
		),
		listTasks: connect_go.NewClient[v1.ListTasksRequest, v1.ListTasksResponse](
			httpClient,
			baseURL+TodoServiceListTasksProcedure,
			opts...,
		),
	}
}

// todoServiceClient implements TodoServiceClient.
type todoServiceClient struct {
	createTask *connect_go.Client[v1.CreateTaskRequest, v1.CreateTaskResponse]
	listTasks  *connect_go.Client[v1.ListTasksRequest, v1.ListTasksResponse]
}

// CreateTask calls todo.v1.TodoService.CreateTask.
func (c *todoServiceClient) CreateTask(ctx context.Context, req *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// ListTasks calls todo.v1.TodoService.ListTasks.
func (c *todoServiceClient) ListTasks(ctx context.Context, req *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return c.listTasks.CallUnary(ctx, req)
}

// TodoServiceHandler is an implementation of the todo.v1.TodoService service.
type TodoServiceHandler interface {
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
}

// NewTodoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTodoServiceHandler(svc TodoServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	todoServiceCreateTaskHandler := connect_go.NewUnaryHandler(
		TodoServiceCreateTaskProcedure,
		svc.CreateTask,
		opts...,
	)
	todoServiceListTasksHandler := connect_go.NewUnaryHandler(
		TodoServiceListTasksProcedure,
		svc.ListTasks,
		opts...,
	)
	return "/todo.v1.TodoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TodoServiceCreateTaskProcedure:
			todoServiceCreateTaskHandler.ServeHTTP(w, r)
		case TodoServiceListTasksProcedure:
			todoServiceListTasksHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTodoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTodoServiceHandler struct{}

func (UnimplementedTodoServiceHandler) CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.CreateTask is not implemented"))
}

func (UnimplementedTodoServiceHandler) ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.ListTasks is not implemented"))
}

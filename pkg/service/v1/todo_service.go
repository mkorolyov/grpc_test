package v1

import (
	"context"
	"math/rand"

	"github.com/mkorolyov/grpc_test/pkg/api/v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type toDoServiceServer struct {
}

// NewToDoServiceServer creates ToDo service
func NewToDoServiceServer() v1.ToDoServiceServer {
	return &toDoServiceServer{}
}

// Create new todo task
func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	return &v1.CreateResponse{
		Id: rand.Int63(),
	}, nil
}

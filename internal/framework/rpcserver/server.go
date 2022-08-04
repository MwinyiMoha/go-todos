package rpcserver

import (
	"context"
	"go-todos/internal/core/ports"
	pb "go-todos/internal/framework/rpcserver/proto"
	"go-todos/internal/utils/exceptions"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RPCServer struct {
	TodoService ports.TodoService
}

func NewRPCServer(service ports.TodoService) *RPCServer {
	return &RPCServer{
		TodoService: service,
	}
}

func (r *RPCServer) GetTodos(ctx context.Context, req *pb.GetTodosRequest) (*pb.GetTodosResponse, error) {
	var todoItems []*pb.Todo

	todos, err := r.TodoService.GetTodos()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, todoDoc := range *todos {
		todoItems = append(
			todoItems,
			&pb.Todo{
				Id:          todoDoc.ID,
				Description: todoDoc.Description,
				CreatedAt:   timestamppb.New(todoDoc.CreatedAt),
			},
		)
	}

	return &pb.GetTodosResponse{TodoItems: todoItems}, nil
}

func (r *RPCServer) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	todoID := req.GetId()

	todoDoc, err := r.TodoService.GetTodo(todoID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.GetTodoResponse{
		TodoItem: &pb.Todo{
			Id:          todoDoc.ID,
			Description: todoDoc.Description,
			CreatedAt:   timestamppb.New(todoDoc.CreatedAt),
		},
	}, nil
}

func (r *RPCServer) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	description := req.GetDescription()

	todoDoc, err := r.TodoService.CreateTodo(description)
	if err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			if verr.StatusCode == 400 {
				return nil, status.Error(codes.InvalidArgument, verr.Description)
			}

			return nil, status.Error(codes.Internal, verr.Description)
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateTodoResponse{
		TodoItem: &pb.Todo{
			Id:          todoDoc.ID,
			Description: todoDoc.Description,
			CreatedAt:   timestamppb.New(todoDoc.CreatedAt),
		},
	}, nil
}

func (r *RPCServer) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	todoID := req.GetId()
	description := req.GetDescription()

	todoDoc, err := r.TodoService.UpdateTodo(todoID, description)
	if err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			if verr.StatusCode == 400 {
				return nil, status.Error(codes.InvalidArgument, verr.Description)
			}

			return nil, status.Error(codes.Internal, verr.Description)
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateTodoResponse{
		TodoItem: &pb.Todo{
			Id:          todoDoc.ID,
			Description: todoDoc.Description,
			CreatedAt:   timestamppb.New(todoDoc.CreatedAt),
		},
	}, nil
}

func (r *RPCServer) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	todoID := req.GetId()

	err := r.TodoService.DeleteTodo(todoID)
	if err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			if verr.StatusCode == 404 {
				return nil, status.Error(codes.NotFound, verr.Description)
			}

			return nil, status.Error(codes.Internal, verr.Description)
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteTodoResponse{}, nil
}

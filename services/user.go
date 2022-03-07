package services

import (
	"context"
	"fmt"

	"github.com/ridneto/fc2-grpc/pb"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//}

type UserService struct {
	// para nao dar pau se adc um serviço que não existe
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

package grpc

import (
	"context"
	"user-service/internal/adapters/api/service"
	"user-service/pkg/grpc"
)

type userServer struct {
	grpc.UnimplementedUserGrpcServiceServer
	s service.UserService
}

func NewUserServer(s service.UserService) grpc.UserGrpcServiceServer {
	return &userServer{
		s: s,
	}
}

func (serv *userServer) RegisterUser(ctx context.Context, dto *grpc.UserRegisterDTO) (*grpc.UserRegisterResponse, error) {
	result, err := serv.s.AddUser(dto.GetUsername(), dto.GetPassword())
	if err != nil {
		return &grpc.UserRegisterResponse{
			Details: &grpc.Details{
				Success: false,
				Mess: err.Error(),
			},
			User: nil,
		}, err
	}
	user := &grpc.UserDTO{
		Id: result.Id,
		Username: result.Username,
		Token: result.Tocken,
	}
	return &grpc.UserRegisterResponse{
		Details: nil,
		User: user,
	}, nil
}

func (serv *userServer) LoginUser(context.Context, *grpc.UserLoginDTO) (*grpc.UserLoginResponse, error) {
	return nil, nil
}
func (serv *userServer) VerifyUser(context.Context, *grpc.UserToken) (*grpc.UserVerifyResponse, error) {
	// result, err := serv.s.VerifyUser(dto.Token)
	return nil, nil
}
func (serv *userServer) GetUserById(context.Context, *grpc.UserId) (*grpc.UserResponse, error) {
	return nil, nil
}
func (serv *userServer) GetUserByTocken(context.Context, *grpc.UserToken) (*grpc.UserResponse, error) {
	return nil, nil
}

    
    
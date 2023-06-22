package user_grpc

import (
	"context"
	"user-service/internal/adapters/api/service"
	user_grpc "user-service/pkg/grpc/user"
)

type userServer struct {
	user_grpc.UnimplementedUserGrpcServiceServer
	s service.UserService
}

func NewUserServer(s service.UserService) user_grpc.UserGrpcServiceServer {
	return &userServer{
		s: s,
	}
}

func (serv *userServer) RegisterUser(ctx context.Context, dto *user_grpc.UserRegisterDTO) (*user_grpc.UserRegisterResponse, error) {
	result, err := serv.s.RegisterUser(dto.GetUsername(), dto.GetPassword())
	if err != nil {
		return &user_grpc.UserRegisterResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			User: nil,
		}, nil
	}
	user := &user_grpc.UserDTO{
		Id:       result.Id,
		Username: result.Username,
		Token:    result.Token,
	}
	return &user_grpc.UserRegisterResponse{
		Details: &user_grpc.Details{
			Success: true,
			Mess: "",
		},
		User:    user,
	}, nil
}

func (serv *userServer) LoginUser(ctx context.Context, dto *user_grpc.UserLoginDTO) (*user_grpc.UserLoginResponse, error) {
	result, err := serv.s.LoginUser(dto.GetUsername(), dto.GetPassword())
	if err != nil {
		return &user_grpc.UserLoginResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess: err.Error(),
			},
			Token: "",
		}, nil
	}
	token := result.Token
	return &user_grpc.UserLoginResponse{
		Details: &user_grpc.Details{
			Success: true,
			Mess: "",
		},
		Token: token,
	}, nil
}
func (serv *userServer) VerifyUser(ctx context.Context, dto *user_grpc.UserToken) (*user_grpc.UserVerifyResponse, error) {
	result, err := serv.s.VerifyUser(dto.GetToken())
	if result == nil {
		return &user_grpc.UserVerifyResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess: "bad token",
			},
			User: nil,
		}, nil
	}
	if err != nil {
		return &user_grpc.UserVerifyResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess: err.Error(),
			},
			User: nil,
		}, nil
	}
	user := &user_grpc.UserDTO{
		Id: result.Id,
		Username: result.Username,
		Token: result.Token,
	}
	return &user_grpc.UserVerifyResponse{
		Details: &user_grpc.Details{
			Success: true,
			Mess: "",
		},
		User: user,
	}, nil
}
func (serv *userServer) GetUserById(ctx context.Context, dto *user_grpc.UserId) (*user_grpc.UserResponse, error) {
	result, err := serv.s.GetUserById(dto.GetId())
	if result == nil {
		return &user_grpc.UserResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess: "bad user id",
			},
			User: nil,
		}, nil
	}
	if err != nil {
		return &user_grpc.UserResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess: err.Error(),
			},
			User: nil,
		}, nil
	}
	user := &user_grpc.UserDTO{
		Id: result.Id,
		Username: result.Username,
		Token: result.Token,
	}
	return &user_grpc.UserResponse{
		Details: &user_grpc.Details{
			Success: true,
			Mess: "",
		},
		User: user,
	}, nil
}
func (serv *userServer) GetUserByToken(ctx context.Context, dto *user_grpc.UserToken) (*user_grpc.UserResponse, error) {
	result, err := serv.s.GetUserByToken(dto.GetToken())
	if result == nil {
		return &user_grpc.UserResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess: "bad user id",
			},
			User: nil,
		}, nil
	}
	if err != nil {
		return &user_grpc.UserResponse{
			Details: &user_grpc.Details{
				Success: false,
				Mess: err.Error(),
			},
			User: nil,
		}, nil
	}
	user := &user_grpc.UserDTO{
		Id: result.Id,
		Username: result.Username,
		Token: result.Token,
	}
	return &user_grpc.UserResponse{
		Details: &user_grpc.Details{
			Success: true,
			Mess: "",
		},
		User: user,
	}, nil
}

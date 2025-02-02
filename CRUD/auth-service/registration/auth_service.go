package registration

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rd "common-libs/redis"
)

type RegistrationData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

// func Authentication(ctx context.Context) error {
// 	passhash, err := GetUserPasshash(req.Login)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "could not get user: %v", err)
// 	}

// 	if req.Passhash == passhash {
// 		token, err := GenerateToken(req.Login, time.Hour)
// 		if err != nil {
// 			return nil, status.Errorf(codes.Internal, "could not generate token: %v", err)
// 		}
// 		return &pb.AuthResponse{Token: token, ExpiresIn: int32(time.Hour.Seconds())}, nil
// 	}

// 	return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
// }

func Register(ctx context.Context, data RegistrationData) error {
	passhash, err := GetUserPasshash(data.Login)
	if err != nil {
		return errors.Errorf("can't get user: %v", err)
	} else if passhash != "" {
		return errors.Errorf("user already exists")
	}

	err = RegisterClient(data.Login, data.Passhash)
	if err != nil {
		return status.Errorf(codes.Internal, "could not register user: %v", err)
	}

	err = rd.Client.Set(ctx, data.Taskid, "success", 1*time.Hour).Err()
	if err != nil {
		return errors.Errorf("can't push result into Redis: %v", err)
	}

	return nil
}

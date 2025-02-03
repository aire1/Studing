package registration

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	auth "crud/auth-service/authorization"
)

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

func Register(ctx context.Context, data auth.AuthData) error {
	passhash, err := auth.GetUserPasshash(data.Login)
	if err != nil && err.Error() != "no rows in result set" {
		return errors.Errorf("can't get user: %v", err)
	} else if passhash != "" {
		return errors.Errorf("user already exists")
	}

	err = RegisterClient(data.Login, data.Passhash)
	if err != nil && err.Error() != "no rows in result set" {
		return status.Errorf(codes.Internal, "could not register user: %v", err)
	}

	return nil
}

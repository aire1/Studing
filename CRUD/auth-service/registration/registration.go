package registration

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	auth "crud/auth-service/getAuthorization"
)

type RegData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

func Register(ctx context.Context, data RegData) error {
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

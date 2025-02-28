package middleware

import (
	"net/http"

	"bishack.dev/services/user"
	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func (o *userServiceMock) AccountDetails(token string) *user.User {
	args := o.Called(token)

	resp := args.Get(0)
	if resp == nil {
		return nil
	}

	return resp.(*user.User)
}

type sessionMock struct {
	mock.Mock
}

func (o *sessionMock) GetUser(r *http.Request) map[string]string {
	args := o.Called(r)

	resp := args.Get(0)
	if resp == nil {
		return nil
	}

	return resp.(map[string]string)
}

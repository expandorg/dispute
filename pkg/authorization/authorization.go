package authorization

import (
	"os"
	"strconv"
	"strings"

	"github.com/expandorg/dispute/pkg/authentication"
)

type Authorizer interface {
	SetAuthData(data authentication.AuthData)
	IsModerator() (bool, error)
	GetUserID() uint64
	GetAuthToken() string
}

type authorizor struct {
	authData authentication.AuthData
}

func NewAuthorizer() Authorizer {
	return &authorizor{
		authentication.AuthData{},
	}
}

func (a *authorizor) SetAuthData(data authentication.AuthData) {
	a.authData = data
}

func (a *authorizor) IsModerator() (bool, error) {
	moderators := os.Getenv("MODERATOR_IDS")
	ids := strings.Split(moderators, ",")

	if !hasID(a.authData.UserID, ids) {
		return false, UnauthorizedAccess{}
	}
	return true, nil
}

func (a *authorizor) GetUserID() uint64 {
	return a.authData.UserID
}

func (a *authorizor) GetAuthToken() string {
	return a.authData.Token
}

func hasID(id uint64, ids []string) bool {
	for _, m := range ids {
		mID, _ := strconv.ParseUint(m, 10, 64)
		if mID == id {
			return true
		}
	}
	return false
}

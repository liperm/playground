package builder

import (
	"errors"

	"github.com/go-builder-pattern/internal/consts"
	"github.com/go-builder-pattern/internal/model"
)

type IBuilder interface {
	SetName(*model.User, string)
	SetEmail(*model.User, string)
	SetAccessLevel(*model.User)
	SetId(*model.User)
}

func GetBuilder(userType consts.UserType) (IBuilder, error) {
	switch userType {
	case consts.USER:
		return UserBuilder{}, nil
	case consts.ADMIN:
		return AdminBuilder{}, nil
	}

	return nil, errors.New("builder_not_implemeted")
}

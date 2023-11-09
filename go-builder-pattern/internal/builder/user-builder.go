package builder

import (
	"github.com/go-builder-pattern/internal/model"
	"github.com/google/uuid"
)

type UserBuilder struct{}

func (b UserBuilder) SetName(user *model.User, name string) {
	user.Name = name
}

func (b UserBuilder) SetEmail(user *model.User, email string) {
	user.Email = email
}

func (b UserBuilder) SetAccessLevel(user *model.User) {
	user.AccessLevel = "user"
}

func (b UserBuilder) SetId(user *model.User) {
	user.Id = uuid.New()
}

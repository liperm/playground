package builder

import (
	"github.com/go-builder-pattern/internal/model"
	"github.com/google/uuid"
)

type AdminBuilder struct{}

func (b AdminBuilder) SetName(user *model.User, name string) {
	user.Name = name
}

func (b AdminBuilder) SetEmail(user *model.User, email string) {
	user.Email = email
}

func (b AdminBuilder) SetAccessLevel(user *model.User) {
	user.AccessLevel = "admin"
}

func (b AdminBuilder) SetId(user *model.User) {
	user.Id = uuid.New()
}

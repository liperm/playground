package director

import (
	"github.com/go-builder-pattern/internal/builder"
	"github.com/go-builder-pattern/internal/model"
)

func CreteUser(builder builder.IBuilder) *model.User {
	var user model.User

	builder.SetId(&user)
	builder.SetName(&user, "Zé leleco")
	builder.SetAccessLevel(&user)
	builder.SetEmail(&user, "seila@gmail.com")

	return &user
}

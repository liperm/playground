package main

import (
	"fmt"

	"github.com/go-builder-pattern/internal/builder"
	"github.com/go-builder-pattern/internal/consts"
	"github.com/go-builder-pattern/internal/director"
)

func main() {
	admBuilder, err := builder.GetBuilder(consts.ADMIN)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := director.CreteUser(admBuilder)
	fmt.Println("Admin user: ", *user)

	userBuilder, err := builder.GetBuilder(consts.USER)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user = director.CreteUser(userBuilder)
	fmt.Println("Normal user: ", *user)

	_, err = builder.GetBuilder(10)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

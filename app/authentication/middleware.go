package authentication

import (
	"fmt"
	"nx-go-api/app"
)

type AuthMiddleware struct {
}

func New() *AuthMiddleware {

	return &AuthMiddleware{}
}

func init() {
	err := app.Injector.Provide(New)
	if err != nil {
		fmt.Println("Error providing AuthMiddleware:", err)
		panic(err)
	}
}

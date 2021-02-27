package users

import (
	"github.com/labstack/echo/v4"
	"pdf-backend/context"
)

func NewHttp(e *echo.Echo, ctx context.IContext) {
	SetupCreateUserHandler(e, ctx)
	SetupUpdateUserHandler(e, ctx)
	SetupDeleteUserHandler(e, ctx)
	SetupGetAllUserHandler(e, ctx)
	SetupGetUserDetailHandler(e, ctx)
}

package authentication

import (
	"github.com/labstack/echo/v4"
	"pdf-backend/context"
)

func NewHttp(e *echo.Echo, ctx context.IContext) {
	SetupLoginHandler(e, ctx)
}

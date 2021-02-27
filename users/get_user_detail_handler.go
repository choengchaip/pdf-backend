package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupGetUserDetailHandler(e *echo.Echo, ctx context.IContext) {
	e.GET("/user/:id", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUserStore(ctx)
		service := NewUserService(ctx, store)

		result, err := service.FindUser(params["id"].(string))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, result)
	}, middleware.Authorization(ctx))
}

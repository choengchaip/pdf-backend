package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupDeleteUserHandler(e *echo.Echo, ctx context.IContext) {
	e.DELETE("/user/:id/delete", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUserStore(ctx)
		service := NewUserService(ctx, store)

		err := service.DeleteUser(params["id"].(string))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, http.StatusOK)
	}, middleware.Authorization(ctx))
}

package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupUpdateUserHandler(e *echo.Echo, ctx context.IContext) {
	e.PUT("/user/:id/update", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUserStore(ctx)
		service := NewUserService(ctx, store)

		result, err := service.UpdateUser(params["id"].(string), params["first_name"].(string), params["last_name"].(string), params["role"].(string))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, result)
	}, middleware.Authorization(ctx))
}

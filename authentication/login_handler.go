package authentication

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pdf-backend/context"
)

func SetupLoginHandler(e *echo.Echo, ctx context.IContext) {
	e.POST("/login", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewLoginStore(ctx)
		service := NewLoginService(ctx, store)

		result, err := service.Login(params["username"].(string), params["password"].(string))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, result)
	})
}

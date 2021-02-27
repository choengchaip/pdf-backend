package users

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupGetAllUserHandler(e *echo.Echo, ctx context.IContext) {
	e.GET("/users", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUserStore(ctx)
		service := NewUserService(ctx, store)

		results, err := service.GetAllUser()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, bson.M{
			"items": results,
			"total": len(results),
		})
	}, middleware.Authorization(ctx))
}

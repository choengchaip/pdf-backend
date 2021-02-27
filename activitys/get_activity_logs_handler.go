package activitys

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
	"pdf-backend/users"
)

func SetupGetActivityLogsHandler(e *echo.Echo, ctx context.IContext) {
	e.GET("/activity-logs", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		userStore := users.NewUserStore(ctx)
		userService := users.NewUserService(ctx, userStore)

		store := NewActivityStore(ctx)
		service := NewActivityService(ctx, store, userService)

		results, err := service.GetAllActivities()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, bson.M{
			"items": results,
			"total": len(results),
		})
	}, middleware.Authorization(ctx))
}

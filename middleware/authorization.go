package middleware

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"pdf-backend/context"
)

func Authorization(ctx context.IContext) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			params := context.Request().Header
			if params["Authorization"] == nil {
				return context.JSON(http.StatusUnauthorized, bson.M{
					"message": "Token is invalid",
				})
			}

			m := ctx.MongoDB(ctx.Config().MongoConfig())
			result, err := m.FindOne("tokens", bson.M{
				"token": params["Authorization"][0],
			})
			if err != nil {
				return context.JSON(http.StatusUnauthorized, bson.M{
					"message": http.StatusText(http.StatusUnauthorized),
				})
			}
			if result == nil {
				return context.JSON(http.StatusUnauthorized, bson.M{
					"message": "User Not Found",
				})
			}

			ctx.SetUserID(result["user_id"].(string))
			return next(context)
		}
	}
}

package pdf

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupPDFGetListHandler(e *echo.Echo, ctx context.IContext) {
	e.GET("/upload-pdfs", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUploadStore(ctx)
		service := NewUploadService(ctx, store)

		filter := bson.M{}
		filter = ctx.BindString(filter, params, "q")
		filter = ctx.BindString(filter, params, "approver_id")
		filter = ctx.BindString(filter, params, "file_name")
		filter = ctx.BindTime(filter, params, "upload_at_start")
		filter = ctx.BindTime(filter, params, "upload_at_end")
		filter = ctx.BindTime(filter, params, "expire_at_start")
		filter = ctx.BindTime(filter, params, "expire_at_end")

		results, err := service.GetAllPDF(filter)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, bson.M{
			"items": results,
			"total": len(results),
		})
	}, middleware.Authorization(ctx))
}

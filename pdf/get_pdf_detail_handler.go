package pdf

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupPDFGetDetailHandler(e *echo.Echo, ctx context.IContext) {
	e.GET("/upload-pdf/:id", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUploadStore(ctx)
		service := NewUploadService(ctx, store)

		result, err := service.FindPDF(params["id"].(string))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, result)
	}, middleware.Authorization(ctx))
}

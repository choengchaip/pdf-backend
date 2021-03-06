package pdf

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupApprovePDFHandler(e *echo.Echo, ctx context.IContext) {
	e.PUT("/approve-pdf/:id", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUploadStore(ctx)
		service := NewUploadService(ctx, store)

		err := service.ApprovePDF(params["user_id"].(string), params["id"].(string))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, http.StatusOK)
	}, middleware.Authorization(ctx))
}

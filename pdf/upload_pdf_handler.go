package pdf

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pdf-backend/context"
	"pdf-backend/middleware"
)

func SetupPDFUploadHandler(e *echo.Echo, ctx context.IContext) {
	e.POST("/upload-pdf", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUploadStore(ctx)
		service := NewUploadService(ctx, store)

		file, err := context.FormFile("file")
		if err != nil {
			return err
		}

		result, err := service.PDFUpload(params["file_name"].(string), params["approver_id"].(string), params["expire_at"].(string), file)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return context.JSON(http.StatusOK, result)
	}, middleware.Authorization(ctx))
}

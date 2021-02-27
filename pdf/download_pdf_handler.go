package pdf

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"pdf-backend/context"
)

func SetupPDFDownloadHandler(e *echo.Echo, ctx context.IContext) {
	e.GET("/download-pdf/:id", func(context echo.Context) error {
		params := map[string]interface{}{}
		if err := context.Bind(&params); err != nil {
			return err
		}

		store := NewUploadStore(ctx)
		service := NewUploadService(ctx, store)

		downloadURL, err := service.PDFDownload(url.QueryEscape(params["id"].(string)))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		if downloadURL == "" {
			return context.JSON(http.StatusInternalServerError, "The file is not exists")
		}

		return context.File(downloadURL)
	})
}

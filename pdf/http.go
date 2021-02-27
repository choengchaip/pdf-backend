package pdf

import (
	"github.com/labstack/echo/v4"
	"pdf-backend/context"
)

func NewHttp(e *echo.Echo, ctx context.IContext) {
	SetupPDFUploadHandler(e, ctx)
	SetupPDFDownloadHandler(e, ctx)
	SetupPDFDeleteHandler(e, ctx)
	SetupPDFGetDetailHandler(e, ctx)
	SetupPDFGetListHandler(e, ctx)
	SetupApprovePDFHandler(e, ctx)
}

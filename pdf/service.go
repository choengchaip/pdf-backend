package pdf

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"pdf-backend/context"
	"pdf-backend/utils"
)

type IUploadService interface {
	PDFUpload(fileName string, approverID string, expireAt string, file *multipart.FileHeader) (bson.M, error)
	PDFDownload(fileAlias string) (string, error)
	PDFDelete(fileName string) error
	ApprovePDF(userID string, fileName string) error
	FindPDF(fileName string) (bson.M, error)
	GetAllPDF(filter bson.M) ([]bson.M, error)
}

type UploadService struct {
	ctx   context.IContext
	store IUploadStore
}

func NewUploadService(ctx context.IContext, store IUploadStore) IUploadService {
	return &UploadService{
		ctx:   ctx,
		store: store,
	}
}

func (s *UploadService) PDFUpload(fileName string, approverID string, expireAt string, file *multipart.FileHeader) (bson.M, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileAlias := utils.GetUUID()
	path := fmt.Sprintf(`tmp/%s`, fileAlias)
	destination, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer destination.Close()

	if _, err = io.Copy(destination, src); err != nil {
		return nil, err
	}

	result, err := s.store.PDFUpload(fileName, approverID, fileAlias, fmt.Sprintf(`%s/download-pdf/%s`, s.ctx.Config().BaseURL(), url.QueryEscape(fileAlias)), utils.NewTimeStampT(expireAt))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UploadService) PDFDelete(fileName string) error {
	err := s.store.PDFDelete(fileName)
	if err != nil {
		return err
	}

	return nil
}

func (s *UploadService) ApprovePDF(userID string, fileName string) error {
	pdf, err := s.FindPDF(fileName)
	if err != nil {
		return err
	}
	if pdf == nil {
		return nil
	}

	err = s.store.PDFUpdate(fileName, bson.M{
		"status": "Approved",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UploadService) FindPDF(fileAlias string) (bson.M, error) {
	result, err := s.store.FindPDF(fileAlias)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UploadService) PDFDownload(fileAlias string) (string, error) {
	file, err := s.FindPDF(fileAlias)
	if err != nil {
		return "", err
	}
	if file == nil {
		return "", nil
	}

	return file["internal_path"].(string), nil
}

func (s *UploadService) GetAllPDF(filter bson.M) ([]bson.M, error) {
	results, err := s.store.GetAllPDF(filter)
	if err != nil {
		return nil, err
	}

	return results, nil
}

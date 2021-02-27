package pdf

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"pdf-backend/context"
	"time"
)

type IUploadStore interface {
	PDFUpload(fileName string, approverID string, fileAlias string, fileSrc string, expireAt time.Time) (bson.M, error)
	PDFDelete(fileName string) error
	PDFUpdate(fileName string, params bson.M) error
	FindPDF(fileName string) (bson.M, error)
	GetAllPDF(filter bson.M) ([]bson.M, error)
}

type UploadStore struct {
	ctx context.IContext
}

func NewUploadStore(ctx context.IContext) IUploadStore {
	return &UploadStore{
		ctx: ctx,
	}
}

func (s *UploadStore) PDFUpload(fileName string, approverID string, fileAlias string, fileSrc string, expireAt time.Time) (bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	_, err := m.InsertOne("files", bson.M{
		"file_name":     fileName,
		"file_alias":    fileAlias,
		"download_url":  fileSrc,
		"internal_path": fmt.Sprintf(`tmp/%s`, fileAlias),
		"approver_id":   approverID,
		"status":        "Not Approve Yet",
		"expire_at":     expireAt.Unix(),
		"upload_at":     time.Now().Unix(),
		"upload_by":     s.ctx.UserID(),
		"is_active":     true,
	})
	if err != nil {
		return nil, err
	}

	return bson.M{
		"file_name":     fileName,
		"file_alias":    fileAlias,
		"download_url":  fileSrc,
		"internal_path": fmt.Sprintf(`tmp/%s`, fileAlias),
		"approver_id":   approverID,
		"status":        "Not Approve Yet",
		"expire_at":     expireAt.Unix(),
		"upload_at":     time.Now().Unix(),
		"upload_by":     s.ctx.UserID(),
		"is_active":     true,
	}, nil
}

func (s *UploadStore) PDFDelete(fileName string) error {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	_, err := m.UpdateOne("files", bson.M{
		"file_name": fileName,
	}, bson.M{
		"$set": bson.M{
			"is_active":  true,
			"deleted_at": time.Now().Unix(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UploadStore) PDFUpdate(fileName string, params bson.M) error {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	_, err := m.UpdateOne("files", bson.M{
		"file_alias": fileName,
	}, bson.M{
		"$set": params,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UploadStore) FindPDF(fileAlias string) (bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	result, err := m.FindOne("files", bson.M{
		"file_alias": fileAlias,
		"is_active":  true,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UploadStore) GetAllPDF(filter bson.M) ([]bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())

	if ok := filter["q"]; ok != nil {
		filter["$or"] = bson.M{
			"file_name":  filter["q"],
			"file_alias": filter["q"],
			"upload_by":  filter["q"],
		}
		delete(filter, "q")
	}
	if ok := filter["upload_at_start"]; ok != nil {
		filter["upload_at"] = bson.M{
			"$gt": filter["upload_at_start"].(time.Time).Unix(),
		}
		delete(filter, "upload_at_start")
	}
	if ok := filter["upload_at_end"]; ok != nil {
		filter["upload_at"] = bson.M{
			"$lt": filter["upload_at_end"].(time.Time).Unix(),
		}
		delete(filter, "upload_at_end")
	}
	if ok := filter["expire_at_start"]; ok != nil {
		filter["expire_at"] = bson.M{
			"$gt": filter["expire_at_start"].(time.Time).Unix(),
		}
		delete(filter, "expire_at_start")
	}
	if ok := filter["expire_at_end"]; ok != nil {
		filter["expire_at"] = bson.M{
			"$lt": filter["expire_at_end"].(time.Time).Unix(),
		}
		delete(filter, "expire_at_end")
	}

	results, err := m.Find("files", filter)
	if err != nil {
		return nil, err
	}

	return results, nil
}

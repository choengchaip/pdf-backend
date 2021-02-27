package pdf

import "mime/multipart"

type IApproveModel struct {
	FileAlias string `json:"id"`
	UserID    string `json:"user_id"`
}

type IDeleteModel struct {
	FileAlias string `json:"id"`
}

type IUpdateModel struct {
	FileName   string                `json:"file_name" bson:"file_name"`
	ApproverID string                `json:"approver_id" bson:"approver_id"`
	ExpireAt   string                `json:"expire_at" bson:"approver_id"`
	File       *multipart.FileHeader `json:"file" bson:"approver_id"`
}

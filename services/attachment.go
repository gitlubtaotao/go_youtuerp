package services

import (
	"youtuerp/models"
	"youtuerp/repositories"
	"youtuerp/tools/uploader"
)

type IAttachment interface {
	Delete(id uint) error
	Create(attachment models.Attachment) (models.Attachment, error)
	FindAll(attr map[string]interface{}) ([]models.Attachment, error)
}

type Attachment struct {
	repo repositories.IAttachment
}

func (a Attachment) Delete(id uint) error {
	attach, err := a.repo.First(id)
	if err != nil {
		return err
	}
	if err := a.repo.Delete(id); err != nil {
		return err
	}
	up := uploader.NewQiNiuUploaderDefault()
	go up.DeleteFile(attach.Key)
	return nil
}

func (a Attachment) FindAll(attr map[string]interface{}) ([]models.Attachment, error) {
	attachments, err := a.repo.FindAll(attr)
	if err != nil {
		return nil, err
	}
	up := uploader.NewQiNiuUploaderDefault()
	for i := 0; i < len(attachments); i++ {
		attachments[i].Url = up.PrivateReadURL(attachments[i].Key)
	}
	return attachments, nil
}

func (a Attachment) Create(attachment models.Attachment) (models.Attachment, error) {
	return a.repo.Create(attachment)
}

func NewAttachment() IAttachment {
	return Attachment{repo: repositories.NewAttachment()}
}

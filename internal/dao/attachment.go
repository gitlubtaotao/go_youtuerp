package dao

import (
	"youtuerp/database"
	"youtuerp/internal/models"
)

type IAttachment interface {
	First(id uint) (models.Attachment, error)
	//删除附件
	Delete(id uint) error
	//查询所有的附件
	FindAll(attr map[string]interface{}) ([]models.Attachment, error)
	//创建附件
	Create(attachment models.Attachment) (models.Attachment, error)
}

type Attachment struct {
}

func (a Attachment) First(id uint) (models.Attachment, error) {
	var attach models.Attachment
	err := database.GetDBCon().First(&attach, "id = ?", id).Error
	return attach, err
}

func (a Attachment) Delete(id uint) error {
	return database.GetDBCon().Where("id = ?", id).Delete(models.Attachment{}).Error
}

func (a Attachment) FindAll(attr map[string]interface{}) ([]models.Attachment, error) {
	var attachments []models.Attachment
	err := database.GetDBCon().Where(attr).Find(&attachments).Error
	return attachments, err
}

func (a Attachment) Create(attachment models.Attachment) (models.Attachment, error) {
	err := database.GetDBCon().Create(&attachment).Error

	return attachment, err
}

func NewAttachment() IAttachment {
	return Attachment{}
}

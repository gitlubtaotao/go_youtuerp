package dao

import (
	"database/sql"
	"youtuerp/database"
	"youtuerp/internal/models"
)

type ICrmTrack interface {
	Create(track models.CrmTrack) (models.CrmTrack, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string, isCount bool) ([]models.CrmTrack, int64, error)
}
type CrmTrack struct {
	BaseRepository
}

func (c CrmTrack) Create(track models.CrmTrack) (models.CrmTrack, error) {
	err := database.GetDBCon().Create(&track).Error
	return track, err
}

func (c CrmTrack) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string, isCount bool) (tracks []models.CrmTrack, total int64, err error) {
	var rows *sql.Rows
	sqlCon := database.GetDBCon().Model(&models.CrmTrack{})
	if isCount {
		if total, err = c.Count(sqlCon, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"crm_tracks.*"}
	}
	sqlCon = c.crud.Where(sqlCon, filter, selectKeys, c.Paginate(per, page), c.OrderBy(orders))
	rows, err = sqlCon.Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.CrmTrack
		err = sqlCon.ScanRows(rows, &data)
		tracks = append(tracks, data)
	}
	return
}

func NewCrmTrack() ICrmTrack {
	return &CrmTrack{}
}

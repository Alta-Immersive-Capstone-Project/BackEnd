package export

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ExportRepo struct {
	db *gorm.DB
}

func NewExportRepo(DB *gorm.DB) *ExportRepo {
	return &ExportRepo{
		db: DB,
	}
}

func (e *ExportRepo) ReportTransaction(time int) ([]entities.House, error) {
	var report []entities.House
	err := e.db.Preload("transactions").Find(&report)
	if err.Error != nil {
		log.Warn(err.Error)
		return report, err.Error
	}
	return report, nil
}

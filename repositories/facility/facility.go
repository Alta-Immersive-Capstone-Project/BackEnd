package facility

import (
	"fmt"
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type FacilityDB struct {
	Db *gorm.DB
}

// Depen
func NewFacilityDB(db *gorm.DB) *FacilityDB {
	return &FacilityDB{
		Db: db,
	}
}

// Create Facility
func (f *FacilityDB) CreateFacility(New entities.Facility) (entities.Facility, error) {
	if err := f.Db.Create(&New).Error; err != nil {
		log.Warn(err)
		return New, err
	}
	return New, nil
}

// Get All Facility
func (f *FacilityDB) GetAllFacility(DistrictID uint) ([]entities.Facility, error) {
	var facilities []entities.Facility
	err := f.Db.Where("district_id = ?", DistrictID).Find(&facilities).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return facilities, err
	}
	return facilities, nil
}

// Get Facility By ID
func (f *FacilityDB) GetFacilityID(id uint) (entities.Facility, error) {
	var facility entities.Facility
	err := f.Db.Where("id= ?", id).First(&facility).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return facility, err
	}
	return facility, nil
}

// Update Facility By ID
func (f *FacilityDB) UpdateFacility(id uint, UpdateFacility entities.Facility) (entities.Facility, error) {
	var facility entities.Facility

	err := f.Db.Where("id =?", id).First(&facility).Updates(&UpdateFacility).Find(&facility).Error
	if err != nil {
		log.Warn("Update Error", err)
		return facility, err
	}

	return facility, nil
}

// Delete Facility By ID
func (f *FacilityDB) DeleteFacility(id uint) error {
	var delete entities.Facility

	err := f.Db.Where("id = ?", id).First(&delete).Delete(&delete).Error
	if err != nil {
		log.Warn("Delete Error", err)
		return err
	}
	return nil
}

// Update Facility By ID
func (f *FacilityDB) GetNearFacility(HouseID uint) ([]entities.NearFacility, error) {

	var respond []entities.NearFacility
	var data []map[string]interface{}

	createView := f.Db.Exec("CREATE OR REPLACE VIEW distance AS SELECT f.name, ST_Distance_Sphere(point(h.latitude, h.longitude - 90), point(f.latitude, f.longitude - 90)) as distance FROM houses h LEFT JOIN facilities f ON f.district_id = h.district_id;")
	if createView.Error != nil {
		log.Warn(createView.Error)
		return []entities.NearFacility{}, createView.Error
	}

	selectFacility := f.Db.Table("distance").Where("distance < 500").Order("distance").Limit(5).Find(&data)
	if selectFacility.Error != nil {
		log.Warn(selectFacility.Error)
		return []entities.NearFacility{}, selectFacility.Error
	}
	for _, v := range data {
		nameFacility := fmt.Sprint(v["name"])

		radiusFacility := v["distance"].(float64)
		nearFacility := entities.NearFacility{Name: nameFacility, Radius: radiusFacility}
		respond = append(respond, nearFacility)
	}
	return respond, nil
}

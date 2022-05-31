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

	var house entities.House
	if err := f.Db.Where("id = ?", HouseID).First(&house).Error; err != nil {
		fmt.Println(err)
		return []entities.NearFacility{}, err
	}

	List, _ := f.GetAllFacility(house.DistrictID)

	sphere := map[string]interface{}{}
	var houseLongitude float64
	if house.Longitude > 90 {
		houseLongitude = house.Longitude - 90
	} else {
		houseLongitude = house.Longitude
	}

	var result []entities.NearFacility
	var nearFacility entities.NearFacility

	for _, v := range List {
		var facilityLongitude float64
		if v.Longitude > 90 {
			facilityLongitude = v.Longitude - 90
		} else {
			facilityLongitude = v.Longitude
		}
		point1 := fmt.Sprintf("POINT(%f %f)", v.Latitude, facilityLongitude)
		point2 := fmt.Sprintf("POINT(%f %f)", house.Latitude, houseLongitude)
		err := f.Db.Raw("SELECT ST_Distance_Sphere(ST_GeomFromText(?),ST_GeomFromText(?)) AS tes", point1, point2)
		err.Find(&sphere)

		jarak := sphere["tes"]
		if jarak.(float64) <= 500 {
			nearFacility.Name = v.Name
			nearFacility.Radius = jarak.(float64)
			result = append(result, nearFacility)
		}
		fmt.Println("result>>>", result)
	}
	return result, nil
}

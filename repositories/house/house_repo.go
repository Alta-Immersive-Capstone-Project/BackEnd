package house

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type HouseRepo struct {
	Db *gorm.DB
}

func NewHouseRepo(db *gorm.DB) *HouseRepo {
	return &HouseRepo{
		Db: db,
	}
}

func (hr *HouseRepo) CreateHouse(addNew entities.House) (entities.House, error) {
	err := hr.Db.Create(&addNew).Error
	if err != nil {
		log.Warn(err)
		return entities.House{}, err
	}
	return addNew, nil
}

func (hr *HouseRepo) UpdateHouse(id uint, update entities.House) (entities.House, error) {
	var house entities.House
	err := hr.Db.Where("id = ?", id).First(&house).Updates(&house).Error
	if err != nil {
		log.Warn("Update Error", err)
		return house, err
	}
	return house, nil
}

func (hr *HouseRepo) DeleteHouse(id uint) error {
	var house entities.House
	err := hr.Db.Where("id = ?", id).First(&house).Delete(&house).Error
	if err != nil {
		log.Warn("Delete Error", err)
		return err
	}
	return nil
}
func (hr *HouseRepo) GetHouseID(id uint) (entities.House, error) {
	var house entities.House
	err := hr.Db.Where("id = ?", id).First(&house).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return house, err
	}
	return house, nil
}
func (hr *HouseRepo) GetAllHouseByDistrict(dist_id uint) ([]entities.House, error) {
	var house []entities.House
	err := hr.Db.Where("district_id = ?", dist_id).Find(&house).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return house, err
	}
	return house, nil
}

func (hr *HouseRepo) GetAllHouseByCities(cid uint) ([]entities.House, error) {
	var houses []entities.House
	err := hr.Db.Joins("JOIN districts ON districts.id = houses.district_id AND districts.city_id = ?", cid).Joins("JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.room_id = rooms.id").Select("houses.id , houses.title , houses.brief , houses.address, houses.district_id, districts.name as district, rooms.type , MIN(rm.price) as price, reviews.id as review_id, AVG(rv.rating) as rating").Group("title").Having("MIN(rm.price) < ?", 1000000).Find(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

func (hr *HouseRepo) GetAllHouseByDstAndCty(cid uint, dist_id uint) ([]entities.House, error) {
	var houses []entities.House
	err := hr.Db.Joins("JOIN districts ON districts.id = houses.district_id AND districts.city_id = ?", cid).Joins("JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.room_id = rooms.id").Select("houses.id , houses.title , houses.brief , houses.address, houses.district_id, districts.name as district, rooms.type , MIN(rm.price) as price, reviews.id as review_id, AVG(rv.rating) as rating").Where("districts.id = ?", dist_id).Find(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

func (hr *HouseRepo) SelectAllHouse() ([]entities.House, error) {
	var houses []entities.House
	err := hr.Db.Joins("JOIN districts ON districts.id = houses.district_id").Joins("JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.room_id = rooms.id").Select("houses.id , houses.title , houses.brief , houses.address, houses.district_id, districts.name as district, rooms.type , MIN(rm.price) as price, reviews.id as review_id, AVG(rv.rating) as rating").Group("title").Having("MIN(rm.price) < ?", 1000000).Find(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}
func (hr *HouseRepo) FindHouseByTitle(title string) ([]entities.House, error) {
	var houses []entities.House
	err := hr.Db.Joins("JOIN districts ON districts.id = houses.district_id").Joins("JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.room_id = rooms.id").Select("houses.id , houses.title , houses.brief , houses.address, houses.district_id, districts.name as district, rooms.type , MIN(rm.price) as price, reviews.id as review_id, AVG(rv.rating) as rating").Where("houses.title like ?", "%"+title+"%").Find(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}
func (hr *HouseRepo) FindHouseByLocation(lat float64, long float64) ([]entities.House, error) {
	var houses []entities.House
	err := hr.Db.Joins("JOIN districts ON districts.id = houses.district_id").Joins("JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.room_id = rooms.id").Select("houses.id , houses.title , houses.brief , houses.address, houses.district_id, districts.name as district, rooms.type , MIN(rm.price) as price, reviews.id as review_id, AVG(rv.rating) as rating").Where("houses.latitude = ? AND house.longitude = ?", lat, long).Find(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

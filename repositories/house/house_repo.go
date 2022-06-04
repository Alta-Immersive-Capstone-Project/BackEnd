package house

import (
	"fmt"
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
	err := hr.Db.Where("id = ?", id).First(&house).Updates(&update).Error
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

func (hr *HouseRepo) GetAllHouseByDist(dist_id uint) ([]entities.HouseResponseGet, error) {
	var houses []entities.HouseResponseGet
	err := hr.Db.Table("houses").Select("houses.*,districts.name as district,districts.city_id as city_id,cities.city as city, MIN(rooms.price) as price, rooms.type as room_type, AVG(reviews.rating) as rating").Group("title").Joins("LEFT JOIN districts ON districts.id = houses.district_id").Joins("LEFT JOIN cities ON districts.city_id = cities.id").Joins("LEFT JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.house_id = houses.id").Where("houses.district_id = ?", dist_id).Scan(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

func (hr *HouseRepo) GetHouseID(id uint) (entities.HouseResponseGetByID, error) {
	var house entities.HouseResponseGetByID
	err := hr.Db.Table("houses").Select("houses.*,districts.name as district,districts.city_id as city_id,cities.city as city").Joins("LEFT JOIN districts ON districts.id = houses.district_id").Joins("LEFT JOIN cities ON districts.city_id = cities.id").Where("houses.id = ?", id).First(&house).Error
	if err != nil {
		log.Warn("Error Get By ID", err)
		return house, err
	}
	return house, nil
}

func (hr *HouseRepo) GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseGet, error) {
	var houses []entities.HouseResponseGet
	err := hr.Db.Table("houses").Select("houses.*,districts.name as district,districts.city_id as city_id,cities.city as city, MIN(rooms.price) as price, rooms.type as room_type, AVG(reviews.rating) as rating").Group("title").Joins("JOIN districts ON districts.id = houses.district_id").Joins("LEFT JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN cities ON districts.city_id = cities.id").Joins("LEFT JOIN reviews ON reviews.house_id = houses.id").Where("houses.district_id = ?", dist_id).Scan(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

func (hr *HouseRepo) GetAllHouseByCities(cid uint) ([]entities.HouseResponseGet, error) {
	var houses []entities.HouseResponseGet
	err := hr.Db.Table("houses").Select("houses.*,districts.name as district,districts.city_id as city_id,cities.city as city, MIN(rooms.price) as price, rooms.type as room_type, AVG(reviews.rating) as rating").Group("title").Joins("JOIN districts ON districts.id = houses.district_id AND districts.city_id = ?", cid).Joins("LEFT JOIN cities ON districts.city_id = cities.id").Joins("LEFT JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.house_id = houses.id").Scan(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

func (hr *HouseRepo) GetAllHouseByDstAndCty(cid uint, dist_id uint) ([]entities.HouseResponseGet, error) {
	var houses []entities.HouseResponseGet
	err := hr.Db.Table("houses").Select("houses.*,districts.name as district,districts.city_id as city_id,cities.city as city, MIN(rooms.price) as price, rooms.type as room_type, AVG(reviews.rating) as rating").Group("title").Joins("JOIN districts ON districts.id = houses.district_id AND districts.city_id = ?", cid).Joins("LEFT JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN cities ON districts.city_id = cities.id").Joins("LEFT JOIN reviews ON reviews.house_id = houses.id").Where("houses.district_id = ?", dist_id).Scan(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

func (hr *HouseRepo) SelectAllHouse() ([]entities.HouseResponseGet, error) {
	var response []entities.HouseResponseGet
	err := hr.Db.Table("houses").Select("houses.*,districts.name as district,districts.city_id as city_id,cities.city as city, MIN(rooms.price) as price, rooms.type as room_type, AVG(reviews.rating) as rating").Group("title").Joins("JOIN districts ON districts.id = houses.district_id").Joins("LEFT JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN cities ON districts.city_id = cities.id").Joins("LEFT JOIN reviews ON reviews.house_id = houses.id").Scan(&response).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return response, err
	}
	fmt.Println(response)
	return response, nil
}

func (hr *HouseRepo) FindHouseByTitle(title string) ([]entities.HouseResponseGet, error) {
	fmt.Println("masuk")
	var houses []entities.HouseResponseGet
	err := hr.Db.Table("houses").Select("houses.*,districts.name as district,districts.city_id as city_id,cities.city as city, MIN(rooms.price) as price, rooms.type as room_type, AVG(reviews.rating) as rating").Group("title").Joins("JOIN districts ON districts.id = houses.district_id").Joins("LEFT JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.house_id = houses.id").Joins("LEFT JOIN cities ON districts.city_id = cities.id").Where("houses.title LIKE ?", "%"+title+"%").Scan(&houses).Error
	if err != nil {
		log.Warn("Error Get Data", err)
		return houses, err
	}
	return houses, nil
}

// func (hr *HouseRepo) FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error) {
// 	var houses []entities.HouseResponseJoin
// 	err := hr.Db.Table("houses").Select("houses.id, houses.title, houses.brief, houses.owner_name, houses.owner_phone, houses.address, houses.available, houses.district_id, districts.name as district, rooms.id as room_id, MIN(rooms.price) as price, rooms.type as type, AVG(reviews.rating) as rating").Joins("JOIN districts ON districts.id = houses.district_id").Joins("JOIN rooms ON rooms.house_id = houses.id").Joins("LEFT JOIN reviews ON reviews.room_id = rooms.id").Where("houses.latitude = ? AND house.longitude = ?", lat, long).Scan(&houses).Error
// 	if err != nil {
// 		log.Warn("Error Get Data", err)
// 		return houses, err
// 	}
// 	return houses, nil
// }

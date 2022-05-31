package city

import (
	"kost/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type cityDB struct {
	Db *gorm.DB
}

func NewCityDB(db *gorm.DB) *cityDB {
	return &cityDB{
		Db: db,
	}
}

func (c *cityDB) CreateCity(new entities.City) (entities.City, error) {
	err := c.Db.Create(&new).Error
	if err != nil {
		log.Warn(err)
		return entities.City{}, err
	}
	return new, nil
}

func (c *cityDB) GetAllCity() ([]entities.City, error) {
	var citys []entities.City
	err := c.Db.Find(&citys).Error
	if err != nil {
		log.Warn(err)
		return []entities.City{}, err
	}
	return citys, nil
}
func (c *cityDB) GetAllCityDistricts(cityID uint) ([]entities.City, error) {
	var citys []entities.City
	err := c.Db.Preload("Districts").Find(&citys).Error
	if err != nil {
		log.Warn(err)
		return []entities.City{}, err
	}
	return citys, nil
}

func (c *cityDB) GetCity(cityID uint) (entities.City, error) {
	var city entities.City
	err := c.Db.Where("id=?", cityID).First(&city).Error
	if err != nil {
		log.Warn(err)
		return entities.City{}, err
	}
	return city, nil
}

func (c *cityDB) DeleteCity(cityID uint) error {
	var city entities.City
	err := c.Db.Where("id = ?", cityID).Delete(&city).Error
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func (c *cityDB) UpdateCity(id uint, new entities.City) (entities.City, error) {
	var city entities.City
	err := c.Db.Where("id = ?", id).First(&city).Updates(&new).Find(&city).Error
	if err != nil {
		log.Warn(err)
		return entities.City{}, err
	}
	return city, nil
}

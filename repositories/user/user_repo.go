package user
<<<<<<< HEAD

import (
	"backend/be8/entities"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (ur *UserRepository) InsertUser(newUser entities.User) (entities.User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot insert data")
	}

	log.Info()
	return newUser, nil
}

func (ur *UserRepository) GetUserID(ID int) (entities.User, error) {
	arrUser := []entities.User{}

	if err := ur.Db.Where("id = ?", ID).Find(&arrUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot select data")
	}

	if len(arrUser) == 0 {
		log.Warn("not found data")
		return entities.User{}, errors.New("not found data")
	}

	log.Info()
	return arrUser[0], nil
}

func (ur *UserRepository) UpdateUser(ID int, update entities.User) (entities.User, error) {
	var res entities.User
	if err := ur.Db.Where("id = ?", ID).Updates(&update).Find(&res).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot update data")
	}

	log.Info()
	return res, nil
}

func (ur *UserRepository) DeleteUser(ID int) (entities.User, error) {
	var user []entities.User
	res, err := ur.GetUserID(ID)
	if err != nil {
		return entities.User{}, err
	}

	if err := ur.Db.Delete(&user, "id = ?", ID).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("cannot delete data")
	}
	log.Info()
	return res, nil

}
=======
>>>>>>> c4ca72a1ed7c4b21751f758877248c248310c2f9

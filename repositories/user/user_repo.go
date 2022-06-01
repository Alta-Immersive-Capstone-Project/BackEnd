package user

import (
	"kost/entities"

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
	tx := ur.Db.Create(&newUser)
	if tx.Error != nil {
		return entities.User{}, tx.Error
	}

	return newUser, nil
}

func (ur *UserRepository) GetUserID(id uint) (entities.User, error) {
	var arrUser entities.User
	tx := ur.Db.Where("id = ?", id).First(&arrUser)
	if tx.Error != nil {
		return entities.User{}, tx.Error
	}

	log.Info()
	return arrUser, nil
}

func (ur *UserRepository) GetAllUser() ([]entities.User, error) {
	var arrUser []entities.User
	err := ur.Db.Where("role = ? OR role = ?", "consultant", "supervisor").Find(&arrUser).Error
	if err != nil {
		log.Warn("error find")
		return []entities.User{}, err
	}

	log.Info(" repo")
	return arrUser, nil
}

func (ur *UserRepository) FindByUser(value string) (entities.User, error) {
	user := entities.User{}
	tx := ur.Db.Where("email = ?", value).First(&user)
	if tx.Error != nil {
		return entities.User{}, tx.Error
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(id uint, user entities.User) (entities.User, error) {

	tx := ur.Db.Save(&user)
	if tx.Error != nil {
		// return Kode 500 jika error
		return entities.User{}, tx.Error
	}
	return user, nil
}

func (ur *UserRepository) DeleteUser(id uint) error {

	// Delete from database
	tx := ur.Db.Delete(&entities.User{}, id)
	if tx.Error != nil {

		// return kode 500 jika error
		return tx.Error
	}
	return nil
}

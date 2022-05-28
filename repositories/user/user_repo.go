package user

import (
	"kost/entities"
	"kost/entities/web"

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
		return entities.User{}, web.WebError{Code: 500, ProductionMessage: "server error", DevelopmentMessage: tx.Error.Error()}
	}

	return newUser, nil
}

func (ur *UserRepository) GetUserID(id int) (entities.User, error) {
	var arrUser []entities.User
	tx := ur.Db.Where("id = ?", id).Find(&arrUser)
	if tx.Error != nil {
		return entities.User{}, web.WebError{Code: 500, ProductionMessage: "server error", DevelopmentMessage: tx.Error.Error()}
	}

	if len(arrUser) == 0 {
		log.Warn("not found data")
		return entities.User{}, web.WebError{Code: 400, ProductionMessage: "bad request", DevelopmentMessage: "data not exist"}
	}

	log.Info()
	return arrUser[0], nil
}
func (ur *UserRepository) FindByUser(field string, value string) (entities.User, error) {
	user := entities.User{}
	tx := ur.Db.Where(field+" = ?", value).Find(&user)
	if tx.Error != nil {

		// return kode 500 jika terjadi error
		return entities.User{}, web.WebError{Code: 500, ProductionMessage: "server error", DevelopmentMessage: tx.Error.Error()}
	} else if tx.RowsAffected <= 0 {

		// return kode 400 jika tidak ditemukan
		return entities.User{}, web.WebError{Code: 400, ProductionMessage: "bad request", DevelopmentMessage: "data not exist"}
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(id int, user entities.User) (entities.User, error) {

	tx := ur.Db.Save(&user)
	if tx.Error != nil {
		// return Kode 500 jika error
		return entities.User{}, web.WebError{Code: 500, ProductionMessage: "server error", DevelopmentMessage: tx.Error.Error()}
	}
	return user, nil
}

func (ur *UserRepository) DeleteUser(id int) error {

	// Delete from database
	tx := ur.Db.Delete(&entities.User{}, id)
	if tx.Error != nil {

		// return kode 500 jika error
		return web.WebError{Code: 500, ProductionMessage: "server error", DevelopmentMessage: tx.Error.Error()}
	}
	return nil
}

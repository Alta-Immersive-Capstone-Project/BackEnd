package reviews

import (
	"kost/entities"
)

type ReviewModel interface {
	Create(review *entities.Review) (*entities.Review, error)
	GetByRoomID(room_id uint) ([]entities.Review, error)
	GetByUserID(user_id uint) (entities.Customer, error)
}

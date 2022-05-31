package reviews

import (
	"kost/entities"
)

type ReviewModel interface {
	Create(review entities.Review) (entities.Review, error)
	GetByRoomID(room_id uint) ([]entities.ReviewJoin, error)
	GetRating(room_id uint) ([]int, float32, error)
}

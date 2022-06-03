package reviews

import (
	"kost/entities"
)

type ReviewModel interface {
	Create(review entities.Review) (entities.Review, error)
	GetByHouseID(HouseID uint) ([]entities.ReviewJoin, error)
	GetRating(HouseID uint) ([]int, float32, error)
}

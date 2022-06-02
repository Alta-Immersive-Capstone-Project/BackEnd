package reviews

import "kost/entities"

type ReviewService interface {
	AddComment(customer_id uint, request entities.ReviewRequest) (entities.ReviewResponse, error)
	GetByHouseID(HouseID uint) ([]entities.ReviewJoin, error)
	GetRating(HouseID uint) ([]int, float32)
}

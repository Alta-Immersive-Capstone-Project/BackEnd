package reviews

import "kost/entities"

type ReviewService interface {
	AddComment(customer_id uint, request *entities.ReviewRequest) (entities.ReviewResponse, error)
	GetByRoomID(room_id uint) ([]entities.ReviewGetResponse, error)
	GetRating(room_id uint) float32
}

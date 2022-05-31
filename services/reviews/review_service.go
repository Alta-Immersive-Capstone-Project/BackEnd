package reviews

import (
	"kost/entities"
	repo "kost/repositories/reviews"

	"github.com/jinzhu/copier"
)

type reviewService struct {
	rm repo.ReviewModel
}

func NewReviewService(rm repo.ReviewModel) *reviewService {
	return &reviewService{
		rm: rm,
	}
}

func (rs *reviewService) AddComment(customer_id uint, request entities.ReviewRequest) (entities.ReviewResponse, error) {
	var response entities.ReviewResponse

	review := entities.Review{
		RoomID:  request.RoomID,
		Comment: request.Comment,
		Rating:  request.Rating,
		UserID:  customer_id,
	}

	result, err := rs.rm.Create(review)
	if err != nil {
		return entities.ReviewResponse{}, err
	}

	copier.Copy(&response, &result)
	return response, nil
}

func (rs *reviewService) GetByRoomID(room_id uint) ([]entities.ReviewJoin, error) {
	responses, err := rs.rm.GetByRoomID(room_id)
	if err != nil {
		return []entities.ReviewJoin{}, err
	}

	return responses, nil
}

func (rs *reviewService) GetRating(room_id uint) ([]int, float32) {
	count, total, err := rs.rm.GetRating(room_id)
	if err != nil {
		return []int{}, 0
	}

	return count, total
}

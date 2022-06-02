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
		HouseID: request.HouseID,
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

func (rs *reviewService) GetByRoomID(HouseID uint) ([]entities.ReviewJoin, error) {
	responses, err := rs.rm.GetByRoomID(HouseID)
	if err != nil {
		return []entities.ReviewJoin{}, err
	}

	return responses, nil
}

func (rs *reviewService) GetRating(HouseID uint) ([]int, float32) {
	count, total, err := rs.rm.GetRating(HouseID)
	if err != nil {
		return []int{}, 0
	}

	return count, total
}

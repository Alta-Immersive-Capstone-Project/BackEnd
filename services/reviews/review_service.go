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

func (rs *reviewService) AddComment(customer_id uint, request *entities.ReviewRequest) (entities.ReviewResponse, error) {
	var response entities.ReviewResponse

	review := entities.Review{
		RoomID:     request.RoomID,
		Comment:    request.Comment,
		Rating:     request.Rating,
		CustomerID: customer_id,
	}

	result, err := rs.rm.Create(&review)
	if err != nil {
		return entities.ReviewResponse{}, err
	}

	copier.Copy(&response, &result)
	return response, nil
}

func (rs *reviewService) GetByRoomID(room_id uint) ([]entities.ReviewGetResponse, error) {
	var responses []entities.ReviewGetResponse

	result, err := rs.rm.GetByRoomID(room_id)
	if err != nil {
		return []entities.ReviewGetResponse{}, err
	}

	for _, r := range result {
		var response entities.ReviewGetResponse
		copier.Copy(&response, &r)

		result, _ := rs.rm.GetByUserID(response.CustomerID)
		response.Name = result.Name

		responses = append(responses, response)
	}

	return responses, nil
}

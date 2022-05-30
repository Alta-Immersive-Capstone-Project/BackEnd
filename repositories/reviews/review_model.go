package reviews

import (
	"kost/entities"

	"gorm.io/gorm"
)

type reviewModel struct {
	db *gorm.DB
}

func NewReviewModel(db *gorm.DB) *reviewModel {
	return &reviewModel{
		db: db,
	}
}

func (m *reviewModel) Create(review entities.Review) (entities.Review, error) {
	record := m.db.Create(&review)

	if record.RowsAffected == 0 {
		return entities.Review{}, record.Error
	}

	return review, nil
}

func (m *reviewModel) GetByRoomID(room_id uint) ([]entities.Review, error) {
	var reviews []entities.Review

	record := m.db.Where("room_id = ?", room_id).Find(&reviews)

	if record.RowsAffected == 0 {
		return []entities.Review{}, record.Error
	}

	return reviews, nil
}

func (m *reviewModel) GetByUserID(user_id uint) (entities.User, error) {
	var customer entities.User

	record := m.db.Where("id = ?", user_id).Find(&customer)

	if record.RowsAffected == 0 {
		return entities.User{}, record.Error
	}

	return customer, nil
}

func (m *reviewModel) GetRating(room_id uint) ([]int, float32, error) {
	var reviews []entities.Review
	var count []int
	record := m.db.Where("room_id = ?", room_id).Find(&reviews)

	if record.RowsAffected == 0 {
		return count, 0, record.Error
	}

	var sum float32
	var one, two, three, four, five int
	for _, review := range reviews {
		switch review.Rating {
		case 1:
			one++
		case 2:
			two++
		case 3:
			three++
		case 4:
			four++
		case 5:
			five++
		}

		sum += review.Rating
	}

	count = []int{one, two, three, four, five}
	total := sum / float32(len(reviews))

	return count, total, nil
}

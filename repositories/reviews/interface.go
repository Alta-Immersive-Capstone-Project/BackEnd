package reviews

import (
	"kost/entities"
)

type ReviewModel interface {
	Create(review entities.Review) (entities.Review, error)
<<<<<<< HEAD
	GetByRoomID(room_id uint) ([]entities.Review, error)
	GetByUserID(user_id uint) (entities.User, error)
=======
	GetByRoomID(room_id uint) ([]entities.ReviewJoin, error)
>>>>>>> 3d2f172cae4224571053c1b5658836fe1402c6a9
	GetRating(room_id uint) ([]int, float32, error)
}

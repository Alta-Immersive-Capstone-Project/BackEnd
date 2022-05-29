package entities

type Image struct {
	RoomID uint   `json:"roomID" validate:"required"`
	Url    string `json:"url"`
}

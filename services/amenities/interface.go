package amenities

import "kost/entities"

type AmenitiesControl interface {
	CreateAmenities(Insert entities.AddAmenities) (entities.RespondAmenities, error)
	GetAmenitiesID(RoomID uint) (entities.RespondAmenities, error)
	UpdateAmenities(RoomID uint, update entities.UpdateAmenities) (entities.RespondAmenities, error)
	DeleteAmenities(RoomID uint) error
}

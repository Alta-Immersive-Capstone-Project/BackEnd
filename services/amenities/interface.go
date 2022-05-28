package amenities

import "kost/entities"

type AmenitiesControl interface {
	CreateAmenities(Insert entities.AddAmenities) (entities.RespondAmenities, error)
	GetAllAmenities(RoomID uint) ([]entities.RespondAmenities, error)
	GetAmenitiesID(id uint) (entities.RespondAmenities, error)
	UpdateAmenities(id uint, update entities.UpdateAmenities) (entities.RespondAmenities, error)
	DeleteAmenities(id uint) error
}

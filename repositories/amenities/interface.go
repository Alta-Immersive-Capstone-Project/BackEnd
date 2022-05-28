package amenities

import "kost/entities"

type RepoAmenities interface {
	CreateAmenities(New entities.Amenities) (entities.Amenities, error)
	GetAllAmenities(RoomID uint) ([]entities.Amenities, error)
	GetAmenitiesID(id uint) (entities.Amenities, error)
	UpdateAmenities(id uint, UpdateAmenities entities.Amenities) (entities.Amenities, error)
	DeleteAmenities(id uint) error
}

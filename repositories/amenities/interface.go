package amenities

import "kost/entities"

type RepoAmenities interface {
	CreateAmenities(New entities.Amenities) (entities.Amenities, error)
	GetAmenitiesID(RoomID uint) (entities.Amenities, error)
	UpdateAmenities(RoomID uint, UpdateAmenities entities.Amenities) (entities.Amenities, error)
	DeleteAmenities(RoomID uint) error
}

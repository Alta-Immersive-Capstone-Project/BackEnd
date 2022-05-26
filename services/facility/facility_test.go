package facility

import (
	"errors"
	"kost/entities"

	mocks "kost/mocks/repositories/facility"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Create Sample Data
var MockFacility = []entities.Facility{
	{
		Model:     gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:      "Rumah Sakit Primaya",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		HouseID:   1,
	},
	{
		Model:     gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:      "Grand Chandra Karya",
		Longitude: -6.19395219376102,
		Latitude:  106.85925435178284,
		HouseID:   1,
	},
}

var MockRespondFacility = []entities.RespondFacility{
	{
		ID:        1,
		Name:      "Rumah Sakit Primaya",
		Longitude: -6.168273696181832,
		Latitude:  106.86491520706296,
		HouseID:   1,
	},
	{
		ID:        2,
		Name:      "Grand Chandra Karya",
		Longitude: -6.19395219376102,
		Latitude:  106.85925435178284,
		HouseID:   1,
	},
}

func TestCreateFacility(t *testing.T) {
	t.Run("Success Create Facility", func(t *testing.T) {

		FacilityRepo := mocks.NewRepoFacility(t)
		var New = entities.Facility{
			Name:      "Rumah Sakit Primaya",
			Longitude: -6.168273696181832,
			Latitude:  106.86491520706296,
			HouseID:   1,
		}

		FacilityRepo.On("CreateFacility", New).Return(MockFacility[0], nil).Once()

		var NewFacility = entities.AddNewFacility{
			Name:      "Rumah Sakit Primaya",
			Longitude: -6.168273696181832,
			Latitude:  106.86491520706296,
			HouseID:   1,
		}

		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.CreateFacility(NewFacility)
		assert.Nil(t, err)
		assert.Equal(t, "Rumah Sakit Primaya", result.Name)

	})
	t.Run("Error Access Database", func(t *testing.T) {
		FacilityRepo := mocks.NewRepoFacility(t)
		var New = entities.Facility{
			Name:      "Rumah Sakit Primaya",
			Longitude: -6.168273696181832,
			Latitude:  106.86491520706296,
			HouseID:   1,
		}

		FacilityRepo.On("CreateFacility", New).Return(MockFacility[0], errors.New("Error Access Database")).Once()

		var NewFacility = entities.AddNewFacility{
			Name:      "Rumah Sakit Primaya",
			Longitude: -6.168273696181832,
			Latitude:  106.86491520706296,
			HouseID:   1,
		}

		FacilityService := NewServiceFacility(FacilityRepo)
		result, err := FacilityService.CreateFacility(NewFacility)
		assert.NotNil(t, err)
		assert.NotEqual(t, "Rumah Sakit Primaya", result.Name)
	})
}

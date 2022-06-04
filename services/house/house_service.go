package house

import (
	"fmt"
	"kost/entities"
	"kost/repositories/house"
	"kost/repositories/room"
	"kost/utils/s3"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

type HouseService struct {
	repo house.IRepoHouse
	room room.RoomRepo
	s3   s3.S3Control
}

func NewHouseService(Repo house.IRepoHouse, Room room.RoomRepo, S3 s3.S3Control) *HouseService {
	return &HouseService{
		repo: Repo,
		room: Room,
		s3:   S3,
	}
}

var ImagesDefault string = "https://belajar-be.s3.ap-southeast-1.amazonaws.com/room/1653973008.png"

func (hs *HouseService) CreateHouse(Insert entities.AddHouse, url string) (entities.HouseResponse, error) {
	var house entities.House
	copier.Copy(&house, &Insert)
	if url == "" {
		house.Image = ImagesDefault
	} else {
		house.Image = url
	}
	res, err := hs.repo.CreateHouse(house)
	if err != nil {
		log.Warn(err)
		return entities.HouseResponse{}, err
	}
	var result entities.HouseResponse
	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) UpdateHouse(id uint, update entities.House) (entities.HouseResponse, error) {
	// var UpdateHouse entities.House
	// copier.Copy(&UpdateHouse, &update)
	res, err := hs.repo.UpdateHouse(id, update)
	if err != nil {
		log.Warn(err)
		return entities.HouseResponse{}, err
	}
	fmt.Println(res)
	var result entities.HouseResponse

	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) DeleteHouse(id uint) error {

	res, err := hs.repo.GetHouseID(id)
	if err != nil {
		log.Warn(err)
		return err
	}
	fmt.Println(res.Image)
	if res.Image != ImagesDefault {
		file := strings.Replace(res.Image, "https://belajar-be.s3.ap-southeast-1.amazonaws.com/", "", 1)
		fmt.Println(file)
		err = hs.s3.DeleteFromS3(file)
		if err != nil {
			return err
		}
	}

	err = hs.repo.DeleteHouse(id)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func (hs *HouseService) GetHouseID(id uint) (entities.HouseResponse, error) {
	res, err := hs.repo.GetHouseID(id)
	if err != nil {
		log.Warn(err)
		return entities.HouseResponse{}, err
	}

	var result entities.HouseResponse
	copier.Copy(&result, &res)

	resul, err := hs.room.GetbyHouse(id)
	if err != nil {
		log.Warn(err)
		return entities.HouseResponse{}, err
	}
	result.Rooms = resul

	return result, nil
}

func (hs *HouseService) GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseGetAll, error) {
	res, err := hs.repo.GetAllHouseByDist(dist_id)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseGetAll{}, err
	}
	var result []entities.HouseResponseGetAll

	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) FindAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseGetAll, error) {
	res, err := hs.repo.GetAllHouseByDistrict(dist_id)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseGetAll{}, err
	}
	var result []entities.HouseResponseGetAll
	copier.Copy(&result, &res)
	return result, nil
}

func (hs *HouseService) FindAllHouseByCities(cid uint) ([]entities.HouseResponseGetAll, error) {
	res, err := hs.repo.GetAllHouseByCities(cid)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseGetAll{}, err
	}
	var result []entities.HouseResponseGetAll
	copier.Copy(&result, &res)
	return result, nil
}
func (hs *HouseService) FindAllHouseByCtyAndDst(cid uint, dist_id uint) ([]entities.HouseResponseGetAll, error) {
	res, err := hs.repo.GetAllHouseByDstAndCty(cid, dist_id)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseGetAll{}, err
	}
	var result []entities.HouseResponseGetAll
	copier.Copy(&result, &res)
	return result, nil
}
func (hs *HouseService) SelectAllHouse() ([]entities.HouseResponseGetAll, error) {
	res, err := hs.repo.SelectAllHouse()
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseGetAll{}, err
	}
	var result []entities.HouseResponseGetAll
	copier.Copy(&result, &res)
	return result, nil
}
func (hs *HouseService) FindHouseByTitle(title string) ([]entities.HouseResponseGetAll, error) {
	res, err := hs.repo.FindHouseByTitle(title)
	if err != nil {
		log.Warn(err)
		return []entities.HouseResponseGetAll{}, err
	}
	var result []entities.HouseResponseGetAll
	copier.Copy(&result, &res)
	return result, nil
}

// func (hs *HouseService) FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error) {
// 	res, err := hs.repo.FindHouseByLocation(lat, long)
// 	if err != nil {
// 		log.Warn(err)
// 		return []entities.HouseResponseJoin{}, err
// 	}
// 	var result []entities.HouseResponseJoin
// 	copier.Copy(&result, &res)
// 	return result, nil
// }

package handlers

import (
	"kost/entities"
	services "kost/services/reviews"
	"net/http"
	"strconv"

	"kost/deliveries/helpers"
	"kost/deliveries/middlewares"
	validation "kost/deliveries/validations"

	"github.com/labstack/echo/v4"
)

type reviewHandler struct {
	rs services.ReviewService
	v  validation.Validation
}

func NewReviewHandler(rs services.ReviewService, v validation.Validation) *reviewHandler {
	return &reviewHandler{
		rs: rs,
		v:  v,
	}
}

func (rh *reviewHandler) InsertComment(c echo.Context) error {
	var request entities.ReviewRequest
	user_id := uint(middlewares.ExtractTokenUserId(c))

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
	}

	err = rh.v.Validation(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	response, err := rh.rs.AddComment(user_id, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
	}

	return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Created Comment", response))
}

func (rh *reviewHandler) GetByRoomID(c echo.Context) error {
	HouseID, _ := strconv.Atoi(c.Param("id"))

	response, _ := rh.rs.GetByRoomID(uint(HouseID))

	if len(response) == 0 {
		return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Comment Not Found"))
	}

	count, total := rh.rs.GetRating(uint(HouseID))

	return c.JSON(http.StatusOK, helpers.StatusOKReview("Success Get By Room ID", response, count, total))
}

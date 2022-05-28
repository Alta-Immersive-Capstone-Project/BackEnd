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
	user_id := uint(middlewares.ExtractTokenUserId(c))
	var request entities.ReviewRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
	}

	err = rh.v.Validation(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	response, err := rh.rs.AddComment(user_id, &request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
	}

	return c.JSON(http.StatusCreated, helpers.StatusCreated("Success Created Comment", response))
}

func (rh *reviewHandler) GetByRoomID(c echo.Context) error {
	room_id, _ := strconv.Atoi(c.Param("room_id"))

	response, err := rh.rs.GetByRoomID(uint(room_id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.StatusNotFound("Comment Not Found"))
	}

	return c.JSON(http.StatusOK, helpers.StatusOK("Success Get By Room ID", response))
}

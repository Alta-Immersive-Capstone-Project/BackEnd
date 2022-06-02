package handlers

import "github.com/labstack/echo/v4"

type RoomHanlder interface {
	CreateRoom() echo.HandlerFunc
	GetAllRoom() echo.HandlerFunc
	GetIDRoom() echo.HandlerFunc
	UpdateRoom() echo.HandlerFunc
	DeleteRoom() echo.HandlerFunc
	DeleteImageUpdate() echo.HandlerFunc
}

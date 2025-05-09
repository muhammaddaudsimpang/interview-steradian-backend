package handler

import (
	"interview/dto"
	."interview/constant"
	"interview/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	carUsecase usecase.CarUsecase
}

func NewCarHandler(carUsecase usecase.CarUsecase) *CarHandler {
	return &CarHandler{
		carUsecase: carUsecase,
	}
}

func (h *CarHandler) GetAllCars(c *gin.Context) {
	cars, err := h.carUsecase.GetAllCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Message: ErrGettingData + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Message: MsgDataOK,
		Data:    cars,
	})
}

func (h *CarHandler) CreateCar(c *gin.Context) {
	var req dto.CreateCarReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Message: ErrNotFound + err.Error(),
		})
		return
	}

	car, err := h.carUsecase.CreateOneCar(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Message: ErrBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Message: MsgCarCreated,
		Data:    car,
	})
}

func (h *CarHandler) DeleteCar(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Message: ErrBadRequest,
		})
		return
	}

	err = h.carUsecase.DeleteOneCar(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ApiResponse{
			Message: ErrNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Message: MsgCarDeleted,
	})
}

func (h *CarHandler) GetCarById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Message: ErrBadRequest,
		})
		return
	}

	car, err := h.carUsecase.GetCarById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ApiResponse{
			Message: ErrNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Message: MsgDataOK,
		Data:    car,
	})
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Message: ErrBadRequest,
		})
		return
	}

	var req dto.UpdateCarReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Message: ErrBadRequest,
		})
		return
	}

	car, err := h.carUsecase.UpdateOneCar(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ApiResponse{
			Message: ErrNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Message: MsgCarUpdated,
		Data:    car,
	})
}

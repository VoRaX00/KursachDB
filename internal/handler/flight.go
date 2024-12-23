package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
	"strconv"
)

type Flight interface {
	Add(schedule services.AddFlight) error
	Delete(id int) error
	GetAll() []models.Flight
}

// @Summary AddFlight
// @Tags flight
// @Description Add new flight
// @ID add-flight
// @Accept json
// @Produce json
// @Router /api/flight/add [post]
func (h *Handler) AddFlight(ctx *gin.Context) {
	var input services.AddFlight
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := validateAddFlight(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := h.services.Flight.Add(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddFlight(schedule services.AddFlight) error {
	panic("implement me")
}

// @Summary DeleteFlight
// @Tags flight
// @Description Delete flight
// @ID delete-flight
// @Accept json
// @Produce json
// @Router /api/flight/delete [delete]
func (h *Handler) DeleteFlight(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err = h.services.Flight.Delete(id); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

// @Summary GetAllFlight
// @Tags flight
// @Description Get all flight
// @ID get-all-flight
// @Accept json
// @Produce json
// @Router /api/flight/ [get]
func (h *Handler) GetAllFlight(ctx *gin.Context) {
	res := h.services.Flight.GetAll()
	ctx.JSON(http.StatusOK, res)
}
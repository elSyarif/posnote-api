package plants

import (
	"context"
	"net/http"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service ports.PlantsService
}

func NewPlantHandler(service ports.PlantsService) *handler {
	return &handler{
		service: service,
	}
}

func (handler *handler) AddPlant(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var plant *domain.Plants
	err := ctx.ShouldBindJSON(&plant)
	if err != nil {
		ctx.Error(err)
		return
	}

	id, err := handler.service.AddPlants(c, plant)
	if err != nil {
		helper.HTTPResponseError(ctx, http.StatusBadRequest, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, http.StatusCreated, gin.H{
		"plantId": id,
	})
}

func (handler *handler) GetPlant(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	name := ctx.Query("name")
	plants, err := handler.service.GetPlants(c, name)
	if err != nil {
		helper.HTTPResponseError(ctx, http.StatusBadRequest, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"plants": plants,
	})
}

func (handler *handler) GetPlantById(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := ctx.Param("id")
	plant, err := handler.service.GetById(c, id)
	if err != nil {
		helper.HTTPResponseError(ctx, http.StatusBadRequest, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"plant": plant,
	})
}

func (handler *handler) Update(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var plant *domain.Plants
	id := ctx.Param("id")

	err := ctx.ShouldBindJSON(&plant)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = handler.service.Update(c, id, plant)
	if err != nil {
		helper.HTTPResponseError(ctx, http.StatusBadRequest, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccess(ctx, 200, "Berhasil memperbaharui plant")
}

func (handler *handler) Delete(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := ctx.Param("id")

	err := handler.service.Delete(c, id)
	if err != nil {
		helper.HTTPResponseError(ctx, http.StatusBadRequest, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccess(ctx, 200, "Berhasil menghapus plant")
}

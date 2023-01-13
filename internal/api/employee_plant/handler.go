package employeeplant

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service ports.EmployeePlantService
}

func NewHandlerEmployeePlant(service ports.EmployeePlantService) *handler {
	return &handler{
		service: service,
	}
}

/**
* TODO: Menambahkan Employee ke plant
 */
func (handler *handler) PostEmployeePlant(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	emplPlant := domain.EmployeePlants{}
	emplPlant.PlantId = ctx.Param("plantId")

	err := ctx.ShouldBindJSON(&emplPlant)
	if err != nil {
		ctx.Error(err)
		return
	}

	id, err := handler.service.AddEmployeePlant(c, &emplPlant)
	if err != nil {
		helper.HTTPResponseError(ctx, 400, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 201, gin.H{
		"employeePlantId": id,
	})
}

/**
* TODO: Menampilkan daftar semua daftar plan dan beserta employee
 */
func (handler *handler) GeteEmployeePlant(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	plantId := ctx.Param("plantId")

	emplPlant, err := handler.service.GetByPlantId(c, plantId)
	if err != nil {
		helper.HTTPResponseError(ctx, 400, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"employeePlant": emplPlant,
	})
}

/**
* TODO: Menampilkan plant employee berdasarkan employee id
 */
func (handler *handler) GetEmployeePlantByEmployeeId(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	plantId := ctx.Param("plantId")

	emplPlant, err := handler.service.GetByPlantId(c, plantId)
	if err != nil {
		helper.HTTPResponseError(ctx, 400, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"employeePlant": emplPlant,
	})

	helper.HTTPResponseSuccess(ctx, 200, "Get Employee plant By Employee Id")
}

/**
* TODO: Menampilkan semua employee berdasarkan plant id
 */
func (handler *handler) GetEmployeePlantByPlantId(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	plantId := ctx.Param("plantId")

	emplPlant, err := handler.service.GetByPlantId(c, plantId)
	if err != nil {
		helper.HTTPResponseError(ctx, 400, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"employeePlant": emplPlant,
	})
}

/**
* TODO: Memperbaharui employee
* Params: [plantId, emplId]
 */
func (handler *handler) PutEmployeePlant(ctx *gin.Context) {
	helper.HTTPResponseSuccess(ctx, 200, "Update Employee plant")
}

/**
* TODO: Mengahpus employee dari daftar plant
* Params: [plantId, emplId]
 */
func (handler *handler) DeleteEmployeePlant(ctx *gin.Context) {
	helper.HTTPResponseSuccess(ctx, 200, "Delete Employee plant")
}

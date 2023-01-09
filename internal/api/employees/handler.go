package employees

import (
	"context"
	"net/http"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service ports.EmployeeService
}

func NewEmployeeHandler(employeeService ports.EmployeeService) *handler {
	return &handler{
		service: employeeService,
	}
}

func (handler *handler) AddEmployee(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var employee *domain.Employees
	err := ctx.ShouldBindJSON(&employee)
	if err != nil {
		ctx.Error(err)
		return
	}

	result, err := handler.service.AddEmployee(c, employee)
	if err != nil {
		helper.HTTPResponseError(ctx, http.StatusBadRequest, "fail", err.Error(), nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 201, gin.H{
		"employeeId": result.Id,
	})
}

func (handler *handler) GetEmployeeById(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := ctx.Param("id")

	result, err := handler.service.GetById(c, id)
	if err != nil {
		helper.HTTPResponseError(ctx, 404, "fail", "employee id tidak ditemukan", nil)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"employee": result,
	})
}

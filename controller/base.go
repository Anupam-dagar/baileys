package controller

import (
	"github.com/Anupam-dagar/baileys/interfaces"
	"github.com/Anupam-dagar/baileys/service"
	"github.com/Anupam-dagar/baileys/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseControllerInterface interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Search(ctx *gin.Context)
}

type baseController[T interfaces.Entity] struct {
	baseService service.BaseServiceInterface[T]
}

func NewBaseController[T interfaces.Entity](rg *gin.RouterGroup) BaseControllerInterface {
	bc := new(baseController[T])
	bc.baseService = service.NewBaseService[T]()
	rg.GET("", bc.Search)
	rg.GET("/:id", bc.GetById)
	rg.POST("", bc.Create)
	rg.PUT("/:id", bc.Update)
	rg.DELETE("/:id", bc.Delete)

	return bc
}

func (bc *baseController[T]) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := bc.baseService.GetById(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully fetched by Id", data)
}

func (bc *baseController[T]) Create(ctx *gin.Context) {
	var payload T

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)

		return
	}

	data, err := bc.baseService.Create(ctx, payload)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully created", data)
}

func (bc *baseController[T]) Update(ctx *gin.Context) {
	var payload T

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)

		return
	}

	id := ctx.Param("id")

	data, err := bc.baseService.Update(ctx, id, payload)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully updated", data)
}

func (bc *baseController[T]) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := bc.baseService.Delete(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully deleted", nil)
}

func (bc *baseController[T]) Search(ctx *gin.Context) {
	var payload T
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)

		return
	}

	data, err := bc.baseService.Search(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully searched", data)
}

package response

import (
	"github.com/Anupam-dagar/baileys/dto"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func ErrorResponse(ctx *gin.Context, statusCode int, err error) {
	log.Error(err)
	statusResponse := &dto.Status{Code: statusCode, Message: err.Error(), Type: "error"}
	errorResponse := &dto.BaseResponse{Status: statusResponse}
	ctx.JSON(statusCode, errorResponse)
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	statusResponse := &dto.Status{Code: statusCode, Message: message, Type: "success"}
	successResponse := &dto.BaseResponse{Status: statusResponse, Data: data}
	ctx.JSON(statusCode, successResponse)
}

func SuccessResponseWithCount(ctx *gin.Context, statusCode int, message string, count int, data interface{}) {
	statusResponse := &dto.Status{Code: statusCode, Message: message, TotalCount: &count, Type: "success"}
	successResponse := &dto.BaseResponse{Status: statusResponse, Data: data}
	ctx.JSON(statusCode, successResponse)
}

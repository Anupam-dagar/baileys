package util

import (
	"baileys/model"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func ErrorResponse(ctx *gin.Context, statusCode int, err error) {
	log.Error(err)
	statusResponse := &model.Status{Code: statusCode, Message: err.Error(), Type: "error"}
	errorResponse := &model.BaseResponse{Status: statusResponse}
	ctx.JSON(statusCode, errorResponse)
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	statusResponse := &model.Status{Code: statusCode, Message: message, Type: "success"}
	successResponse := &model.BaseResponse{Status: statusResponse, Data: data}
	ctx.JSON(statusCode, successResponse)
}

func SuccessResponseWithCount(ctx *gin.Context, statusCode int, message string, count int, data interface{}) {
	statusResponse := &model.Status{Code: statusCode, Message: message, TotalCount: &count, Type: "success"}
	successResponse := &model.BaseResponse{Status: statusResponse, Data: data}
	ctx.JSON(statusCode, successResponse)
}

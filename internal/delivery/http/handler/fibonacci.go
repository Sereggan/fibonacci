package handler

import (
	"fibonachi/internal/delivery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type request struct {
	X uint64 `json:"x"`
	Y uint64 `json:"y"`
}

type response struct {
	Numbers []uint64 `json:"numbers"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) Calculate(context *gin.Context) {
	var request request
	err := context.BindJSON(&request)

	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	isValid := delivery.ValidateInputParameters(request.X, request.Y)
	if !isValid {
		logrus.Errorf("Validation failed for x: %d, y: %d", request.X, request.Y)
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	numbers, err := h.services.Fibonacci.CalculateResult(context, request.X, request.Y)

	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, response{
		Numbers: numbers,
	})
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

package response

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
}

type PaginatedResponse struct {
	Code       int            `json:"code"`
	Status     string         `json:"status"`
	Message    string         `json:"message"`
	Data       interface{}    `json:"data,omitempty"`
	Pagination PaginationMeta `json:"pagination"`
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SuccessWithPagination(c *gin.Context, statusCode int, message string, data interface{}, page, limit int, total int64) {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(statusCode, PaginatedResponse{
		Code:    statusCode,
		Status:  "success",
		Message: message,
		Data:    data,
		Pagination: PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			TotalItems:  total,
			TotalPages:  totalPages,
		},
	})
}

func Error(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Status:  "error",
		Message: message,
		Errors:  errors,
	})
}

func OK(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusOK, message, data)
}

func Created(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusCreated, message, data)
}

func BadRequest(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusBadRequest, message, errors)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, nil)
}

func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message, nil)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

func InternalServerError(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusInternalServerError, message, errors)
}

func ValidationError(c *gin.Context, errors interface{}) {
	Error(c, http.StatusUnprocessableEntity, "Validation failed", errors)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

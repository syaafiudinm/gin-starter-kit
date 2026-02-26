package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/syaafiudinm/go-starter-kit/internal/dto"
	"github.com/syaafiudinm/go-starter-kit/internal/service"
	"github.com/syaafiudinm/go-starter-kit/pkg/response"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result, err := h.userService.Create(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created successfully", result)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	result, err := h.userService.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User retrieved successfully", result)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, total, err := h.userService.GetAll(page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve users", err.Error())
		return
	}

	response.SuccessWithPagination(c, http.StatusOK, "Users retrieved successfully", users, page, limit, total)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result, err := h.userService.Update(uint(id), &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User updated successfully", result)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	if err := h.userService.Delete(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User deleted successfully", nil)
}

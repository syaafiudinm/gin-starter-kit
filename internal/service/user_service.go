package service

import (
	"errors"

	"github.com/syaafiudinm/go-starter-kit/internal/dto"
	"github.com/syaafiudinm/go-starter-kit/internal/model"
	"github.com/syaafiudinm/go-starter-kit/internal/repository"
)

type UserService interface {
	Create(req *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetByID(id uint) (*dto.UserResponse, error)
	GetAll(page, limit int) ([]dto.UserResponse, int64, error)
	Update(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	existingUser, _ := s.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *userService) GetByID(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *userService) GetAll(page, limit int) ([]dto.UserResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, total, err := s.repo.FindAll(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, toUserResponse(&user))
	}

	return responses, total, nil
}

func (s *userService) Update(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		if req.Email != user.Email {
			existingUser, _ := s.repo.FindByEmail(req.Email)
			if existingUser != nil {
				return nil, errors.New("email already exists")
			}
		}
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *userService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repo.Delete(id)
}

func toUserResponse(user *model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

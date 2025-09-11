package service

import (
	"context"
	"typenowsql/models"
	resource "typenowsql/resource"
)

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
	GetCities(ctx context.Context) ([]*models.City, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, id int, req *models.CreateUserRequest) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
}

// CreateUser(ctx context.Context, user *models.User) error
// GetUserByID(ctx context.Context, id int) (*models.User, error)
// GetCities(ctx context.Context) ([]*models.City, error)
// UpdateUser(ctx context.Context, user *models.User) error
// DeleteUser(ctx context.Context, id int) error

type userService struct {
	userRepo resource.UserResource
}

func NewUserService(userRepo resource.UserResource) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id int) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *userService) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.GetUsers(ctx)
}

func (s *userService) GetCities(ctx context.Context) ([]*models.City, error) {
	return s.userRepo.GetCities(ctx)
}

func (s *userService) UpdateUser(ctx context.Context, id int, req *models.CreateUserRequest) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	user.Email = req.Email

	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	return s.userRepo.DeleteUser(ctx, id)
}

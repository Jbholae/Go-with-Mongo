package services

import (
	"context"
	"saurav/mongo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) GetUser(userName *string) (*models.User, error) {
	return nil, nil
}

func (*UserServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(userName *string) error {
	return nil
}

package services

import (
	"context"
	"errors"
	"saurav/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
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
	_,err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(userName *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key:"user_name",Value:userName}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (*UserServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key:"user_name",Value:user.Name}}
	update := bson.D{bson.E{Key:"$set",Value: bson.D{bson.E{Key:"user_name",Value: user.Name},bson.E{Key:"user_age",Value: user.Age},bson.E{Key:"user_adddress",Value: user.Address}}}}
	result,_ := u.userCollection.UpdateOne(u.ctx,filter,update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(userName *string) error {
	return nil
}

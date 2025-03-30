package resolvers

import (
	"context"
	"strconv"

	"github.com/Jeanga7/graphql-go-api.git/graph/model"
	"github.com/Jeanga7/graphql-go-api.git/models"
	"gorm.io/gorm"
)

type UserResolver struct {
    DB *gorm.DB
}

func (r *UserResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	var result []*model.User
	for _, user := range users {
		result = append(result, &model.User{
			ID:       strconv.Itoa(int(user.ID)),
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return result, nil
}

func (r *UserResolver) CreateUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &model.User{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

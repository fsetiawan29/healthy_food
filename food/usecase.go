package food

import (
	"context"

	"github.com/febrycode/healthy_food/models"
)

type Usecase interface {
	CreateFood(ctx context.Context, foodData *models.Food) error
}
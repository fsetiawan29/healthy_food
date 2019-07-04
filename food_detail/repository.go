package food_detail

import (
	"context"

	"github.com/febrycode/healthy_food/models"
)

// Repository represent the article's repository contract
type Repository interface {
	CreateFoodDetail(ctx context.Context, foodDetailData *models.FoodDetail) error
}
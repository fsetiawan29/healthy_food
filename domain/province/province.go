package province

import (
	"context"

	"github.com/fsetiawan29/healthy_food/models"
)

type Repository interface {
	GetAllProvince(ctx context.Context) (result []models.Province, err error)
	GetProvinceByID(ctx context.Context, provinceID int64) (provinceData models.Province, err error)
}

type Usecase interface {
	GetAllProvince(ctx context.Context) (result []models.Province, err error)
}

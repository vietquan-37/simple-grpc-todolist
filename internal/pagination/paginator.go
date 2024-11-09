package pagination

import (
	"math"

	"gorm.io/gorm"
)

type Result[T any] struct {
	TotalPage  int64
	PageNumber int64
	PageSize   int64
	Results    []T
}

func Paginate[T any](db *gorm.DB, PageNumber, PageSize int64) (*Result[T], error) {
	var result Result[T]
	var data []T
	var count int64
	if PageNumber < 0 {
		PageNumber = 0
	}
	if PageSize <= 0 {
		PageSize = 5
	}

	db.Count(&count)
	err := db.
		Limit(int(PageSize)).
		Offset(int(PageNumber - 1)).
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	result.Results = data
	result.PageNumber = PageNumber
	result.PageSize = PageSize
	result.TotalPage = int64(math.Ceil(float64(count) / float64(PageSize)))

	return &result, nil
}

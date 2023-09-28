package models

import (
	"math"
	"time"

	"hospital-app/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string    `gorm:"primarykey" json:"id"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	UUID, _ := uuid.NewRandom()
	u.ID = UUID.String()
	return
}

func Paginate(value interface{}, pagination *pkg.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Where(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

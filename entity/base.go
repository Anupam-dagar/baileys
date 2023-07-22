package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        string         `gorm:"primary_key;column:id;" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	CreatedBy string         `gorm:"column:created_by" json:"createdBy"`
	UpdatedBy string         `gorm:"column:updated_by" json:"updatedBy"`
	DeletedBy string         `gorm:"column:deleted_by" json:"deletedBy"`
}

func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	bm.Id = uuid.NewString()
	bm.CreatedAt = time.Now()
	bm.UpdatedAt = time.Now()

	return
}

func (bm *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	bm.UpdatedAt = time.Now()

	return
}

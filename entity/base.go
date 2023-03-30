package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        string         `gorm:"primary_key;column:id;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	CreatedBy string         `gorm:"column:created_by" json:"createdBy"`
	UpdatedBy string         `gorm:"column:updated_by" json:"updatedBy"`
	DeletedBy string         `gorm:"column:deleted_by" json:"deletedBy"`
}

package entity

import (
	"time"
)

type UserTypes struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Type      string    `json:"employeeType" gorm:"not null;column:type"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at"`
	IsActive  bool      `json:"isActive" gorm:"default:true;column:is_active"`
}

func (UserTypes) TableName() string {
	return "user_types"
}

type User struct {
	ID         uint       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name       string     `gorm:"type:varchar(100);not null;column:name" json:"name"`
	Age        int        `gorm:"not null;column:age" json:"age"`
	Email      string     `gorm:"uniqueIndex;type:varchar(100);not null;column:email" json:"email"`
	BirthDate  *time.Time `gorm:"index;column:birth_date" json:"birthDate,omitempty"`
	UserTypeID int        `gorm:"column:user_type_id" json:"userTypeId"`
	UserType   UserTypes  `gorm:"foreignKey:UserTypeID;references:Id"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `gorm:"index;column:deleted_at" json:"deletedAt,omitempty"`
}

func (User) TableName() string {
	return "users"
}

package corey

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID    int    `gorm:"primarykey"`
	Name  string `gorm:"column:name;index;not null" json:"name" binding:"required"`
	Email string `gorm:"column:email;unique;not null" json:"email" binding:"required,email"`
	Tasks []Task `gorm:"foreignKey:ContactID"`
}

type Task struct {
	ID          int        `gorm:"primarykey"`
	Title       string     `gorm:"column:title;not null" json:"title" binding:"required"`
	Description string     `gorm:"column:description;not null" json:"description" binding:"required"`
	Reminder    *time.Time `gorm:"column:reminder;not null" json:"reminder" binding:"omitempty" time_format:"2006-01-02T15:04:05Z07:00"`
	ContactID   int        `gorm:"column:contact_id;not null" json:"contact_id" binding:"required"`
}

func MigrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Contact{}, &Task{})
	if err != nil {
		return err
	}
	db.Model(&Task{}).Association("Contact")
	return nil
}

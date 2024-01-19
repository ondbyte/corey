package corey

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID    int    `gorm:"primarykey"`
	Name  string `gorm:"column:name;index;not null" json:"name"`
	Email string `gorm:"column:email;unique;not null" json:"email" binding:"required,email"`
	Tasks []Task `gorm:"foreignKey:ContactID"`
}

type Priority string

const (
	HIGH   Priority = "HIGH"
	MEDIUM Priority = "MEDIUM"
	LOW    Priority = "LOW"
)

func IsValidPriority(p Priority) bool {
	return p == HIGH || p == MEDIUM || p == LOW
}

type Task struct {
	ID          int        `gorm:"primarykey"`
	Title       string     `gorm:"column:title;not null" json:"title"`
	Description string     `gorm:"column:description;not null" json:"description"`
	Priority    Priority   `gorm:"column:priority;not null" json:"priority"`
	Reminder    *time.Time `gorm:"column:reminder;not null" json:"reminder" binding:"omitempty" time_format:"2006-01-02T15:04:05Z07:00"`
	ContactID   int        `gorm:"column:contact_id;not null" json:"contact_id"`
}

func MigrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Contact{}, &Task{})
	if err != nil {
		return err
	}
	db.Model(&Task{}).Association("Contact")
	return nil
}

func DropTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(&Contact{}, &Task{})
	if err != nil {
		return err
	}
	return nil
}

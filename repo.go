package corey

import (
	"fmt"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}
func (r *Repo) GetContact(id uint) (t *Contact, err error) {
	t = &Contact{}
	tx := r.db.Model(t).Find(t, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("cannot find for id")
	}
	return t, nil
}

func (r *Repo) AddContact(c *Contact) error {
	tx := r.db.Model(c).Create(c)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repo) GetAllContacts() (t []*Contact, err error) {
	t = []*Contact{}
	tx := r.db.Model(t).Scan(&t)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return t, nil
}

func (r *Repo) GetTask(id uint) (t *Task, err error) {
	t = &Task{}
	tx := r.db.Model(t).Find(t, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("cannot find for id")
	}
	return t, nil
}

func (r *Repo) GetAllTasks() (t []*Task, err error) {
	t = []*Task{}
	tx := r.db.Model(t).Scan(&t)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return t, nil
}

func (r *Repo) AddTask(t *Task) error {
	tx := r.db.Model(t).Create(t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repo) DeleteTask(t *Task) error {
	tx := r.db.Model(t).Delete(t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

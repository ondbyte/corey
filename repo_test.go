package corey_test

import (
	"testing"
	"time"

	"github.com/onbyte/corey"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepo(t *testing.T) {
	assert := assert.New(t)
	db, err := gorm.Open(postgres.Open(corey.Dsn), &gorm.Config{})
	if !assert.NoError(err) {
		return
	}
	err = corey.MigrateModels(db)
	if !assert.NoError(err) {
		return
	}
	defer db.Migrator().DropTable(&corey.Task{}, &corey.Contact{})

	repo := corey.NewRepo(db)

	contact := &corey.Contact{
		Name:  "yadu",
		Email: "i@m.com",
	}

	err = repo.AddContact(contact)
	if !assert.NoError(err) {
		return
	}
	err = repo.AddTask(&corey.Task{
		Title:       "yadus task",
		Description: "yadus task d",
	})
	if !assert.Error(err) {
		return
	}
	now := time.Now()
	err = repo.AddTask(&corey.Task{
		Title:       "yadus task",
		Description: "yadus task d",
		ContactID:   contact.ID,
		Reminder:    &now,
	})
	if !assert.NoError(err) {
		return
	}
}

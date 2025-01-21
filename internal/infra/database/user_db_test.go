package database

import (
	"testing"

	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserDb_ShouldInsertAUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, err := entity.NewUser("Peter Parker", "peter.parker@email.com", "1234567")
	if err != nil {
		t.Error(err)
	}

	userDb := NewUserDb(db)
	err = userDb.Create(user)
	assert.Nil(t, err)

	var userFromDb entity.User
	err = db.First(&userFromDb, "id = ?", user.GetID()).Error
	assert.Nil(t, err)
	assert.Equal(t, user.GetID(), userFromDb.GetID())
	assert.Equal(t, user.GetName(), userFromDb.GetName())
	assert.Equal(t, user.GetEmail(), userFromDb.GetEmail())
	assert.True(t, userFromDb.IsValidPassword("1234567"))
}

func TestUserDb_ShouldFindAUserByEmail(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, err := entity.NewUser("Mary Jane", "mary.jane@email.com", "7654321")
	if err != nil {
		t.Error(err)
	}

	userDb := NewUserDb(db)
	err = userDb.Create(user)
	assert.Nil(t, err)

	userFromDb, err := userDb.FindByEmail(user.GetEmail())
	assert.Nil(t, err)
	assert.Equal(t, user.GetID(), userFromDb.GetID())
	assert.Equal(t, user.GetName(), userFromDb.GetName())
	assert.Equal(t, user.GetEmail(), userFromDb.GetEmail())
	assert.True(t, userFromDb.IsValidPassword("7654321"))
}

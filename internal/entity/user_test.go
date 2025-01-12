package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_NewUser(t *testing.T) {
	user, err := NewUser("John", "john.doe@mail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.GetID())
	assert.NotEmpty(t, user.GetPassword())
	assert.Equal(t, "John", user.GetName())
	assert.Equal(t, "john.doe@mail.com", user.GetEmail())
}

func TestUser_IsValidPassword(t *testing.T) {
	user, err := NewUser("John", "john.doe@mail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.IsValidPassword("123456"))
	assert.False(t, user.IsValidPassword("1234567"))
	assert.NotEqual(t, "123456", user.GetPassword())
}

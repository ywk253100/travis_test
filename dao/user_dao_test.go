package dao

import (
	"testing"
)

func TestUserDao(t *testing.T) {
	users, err := GetAllUser()
	if err != nil {
		t.Error(err)
	}

	t.Log(users)
}

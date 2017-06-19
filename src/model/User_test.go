package model

import (
	"testing"
)

func TestUserAttributes(t *testing.T){
	var testUser = User{BaseModel{1},"frank"}

	if testUser.Id != 1 {
		t.Error("ID is not 1")
	}
	if testUser.Name != "frank" {
		t.Error("Name is not frank")
	}
}
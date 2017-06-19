package model

import (
	"testing"
	"github.com/frankrafael/predictive-maintenance-go/src/model"
)

func TestUserAttributes(t *testing.T){
	var testUser = model.User{model.BaseModel{1},"frank"}

	if testUser.Id != 1 {
		t.Error("ID is not 1")
	}
	if testUser.Name != "frank" {
		t.Error("Name is not frank")
	}
}
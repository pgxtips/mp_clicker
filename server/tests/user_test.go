package tests

import (
	"testing"
	"server/internal"
	"server/internal/models"
)

func TestGetUser(t *testing.T){
	internal.InitApplication()
	appdata := internal.GetAppData()

	u, err := models.GetUser(appdata.DB, "dev1")

	if err != nil {
		t.Errorf("%v", err)
	}

	if u.Email != "dev1@dev.com" || u.Username != "dev1" {
		t.Errorf("failed to retrieved correct user. \nusername: %v  email: %v", u.Username, u.Email)
	}
}

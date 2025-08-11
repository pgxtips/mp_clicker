package tests

import (
	"testing"
	"server/internal"
	"server/internal/models"
)

func TestGetWorld(t *testing.T){
	internal.InitApplication()
	appdata := internal.GetAppData()

	w, err := models.GetWorldData(appdata.DB, 2)

	if err != nil {
		t.Errorf("%v", err)
	}

	if w.Name != "dev2 world 1" {
		t.Errorf("failed to retrieved correct world. \nname: %v", w.Name)
	}
}

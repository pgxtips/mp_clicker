package models

import (
	"time"
)

type WorldData struct{
	id int
	name string
	owner int
	created_at time.Time
	world_data string
	modifier_data string
}

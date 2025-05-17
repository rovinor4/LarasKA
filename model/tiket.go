package model

import "time"

type Tiket struct {
	Penumpang Penumpang
	Rute Rute
	CreatedAt time.Time
}

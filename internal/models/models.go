package models

import (
	"time"
)

// Plant represents the plant entity for sqlx
type Plant struct {
	ID           int       `db:"id" json:"id"`
	Species      string    `db:"species" json:"species"`
	Cultivar     string    `db:"cultivar" json:"cultivar"`
	PlantingDate time.Time `db:"planting_date" json:"planting_date"`
	IsCross      bool      `db:"is_cross" json:"is_cross"`
}

// Cross represents the cross-breeding entity for sqlx
type Cross struct {
	ID         int `db:"id" json:"id"`
	PlantID    int `db:"plant_id" json:"plant_id"`   // Foreign key to Plant
	MotherID   int `db:"mother_id" json:"mother_id"` // Foreign key to Plant
	FatherID   int `db:"father_id" json:"father_id"` // Foreign key to Plant
	Generation int `db:"generation" json:"generation"`
}

// Notes represents the log entity for sqlx
type Notes struct {
	ID      int    `db:"id" json:"id"`
	PlantID int    `db:"plant_id" json:"plant_id"` // Foreign key to Plant
	Height  int    `db:"height" json:"height"`
	Leaves  int    `db:"leaves" json:"leaves"`
	Buds    int    `db:"buds" json:"buds"`
	Notes   string `db:"notes" json:"notes"`
}

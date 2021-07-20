// models/hero.go

package models

// Hero structure
type Hero struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

//HeroInput structure
type HeroInput struct {
	Name string `json:"name" binding:"required"`
}

package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Activity struct {
	gorm.Model
	DocumentId string  `gorm:"uniqueIndex:index_activity;not null"`
	CustomId   string  `gorm:"uniqueIndex:index_activity;not null"`
	Distance   float64 `gorm:"not null"`
	Speed      Speed   `gorm:"foreignKey:DocumentId,CustomId;references:DocumentId,CustomId;"`
}

type Type struct {
	Id          uint `gorm:"primaryKey"`
	Description string
}

type Speed struct {
	gorm.Model
	Mile       float64
	DocumentId string `gorm:"uniqueIndex:index_speed"`
	CustomId   string `gorm:"uniqueIndex:index_speed"`
}

func main() {
	dsn := "[INSERT YOUR OWN CREDENTIAL]"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Activity{}, &Type{}, &Speed{})

	var t = Type{Id: 01, Description: "Type 01"}
	// Create Type
	db.Clauses(clause.OnConflict{
		DoNothing: true,
		Columns:   []clause.Column{{Name: "id"}},
	}).Create(&t)

	// Create Activities
	var activities = []Activity{
		{CustomId: "123", DocumentId: "234", Distance: 53, Speed: Speed{Mile: 32, DocumentId: "234", CustomId: "123"}},
		{CustomId: "345", DocumentId: "456", Distance: 53, Speed: Speed{Mile: 12, DocumentId: "345", CustomId: "123"}},
	}
	db.Debug().Clauses(clause.OnConflict{
		DoNothing: true,
		Columns:   []clause.Column{{Name: "custom_id"}, {Name: "document_id"}},
	}).Create(&activities)
}

func (u *Speed) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Creating speed")
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "custom_id"}, {Name: "document_id"}},
		DoNothing: true,
	})
	return
}
